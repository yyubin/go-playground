package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func PrintHangul() {
	hanguls := []rune{'가', '나', '다', '라', '마', '바', '사'}
	for _, v := range hanguls {
		time.Sleep(300 * time.Millisecond) // 1sec = 1000millisecond
		fmt.Printf("%c ", v)
	}
}

func PrintNumber() {
	for i := 1; i < 6; i++ {
		time.Sleep(400 * time.Millisecond)
		fmt.Printf("%d ", i)
	}
}

var wg sync.WaitGroup

func SumAtoB(a, b int) {
	sum := 0
	for i := a; i <= b; i++ {
		sum += i
	}
	fmt.Printf("%d부터 %d까지 합계는 %d입니다.\n", a, b, sum)
	wg.Done()
}

type Account struct {
	Balance int
}

func DepositAndWithdraw(account *Account) {
	mutex.Lock()         // 추가
	defer mutex.Unlock() // 추가

	if account.Balance < 0 {
		panic(fmt.Sprintf("잔액은 음수가 될 수 없습니다 : %d", account)) // 이론상 발생 x, 발생되는 경우가 있음
	}
	account.Balance += 1000
	time.Sleep(time.Millisecond)
	account.Balance -= 1000
}

var mutex sync.Mutex

func diningProblem(name string, first, second *sync.Mutex, firstName, secondName string) {
	for i := 0; i < 100; i++ {
		fmt.Printf("%s 밥을 먹으려 합니다\n", name)
		first.Lock()
		fmt.Printf("%s %s 획득\n", name, firstName)
		second.Lock()
		fmt.Printf("%s %s 획득\n", name, secondName)

		fmt.Printf("%s 밥을 먹습니다", name)
		time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)

		second.Unlock()
		first.Unlock()
	}
	wg.Done()
}

type Job interface {
	Do()
}

type SquareJob struct {
	index int
}

func (j *SquareJob) Do() {
	fmt.Printf("%d 작업 시작\n", j.index)
	time.Sleep(1 * time.Second)
	fmt.Printf("%d 작업 완료 - 결과 %d\n", j.index, j.index*j.index)
}

func main() {
	go PrintHangul() // go 루틴 실행
	go PrintNumber()

	time.Sleep(3 * time.Second) // 메인함수는 3초동안 잠듦

	// 총 3개의 고루틴 생성
	// cpu 코어 3개 이상이라면 컨텍스트 스위칭 발생 X
	// 컨텍스트 스위칭 : IP, 레지스터 카운터, stack 메모리 등 쓰레드 변환 과정
	// *힙메모리는 모든 프로세스, 쓰레드가 공유

	// 서브 고루틴이 종료될때까지 대기
	wg.Add(10) // wg.Done()이 되면 10에서 1씩 줄고 10개가 다 끝나면 종료
	for i := 0; i < 10; i++ {
		go SumAtoB(1, 1000000000)
	}

	wg.Wait()

	// 고루틴 동작원리
	// 고루틴은 OS쓰레드를 이용하는 경량 쓰레드 (고루틴 != 쓰레드)
	// 고루틴이 쓰레드를 이용함

	// 코어가 2개, 고루틴이 3개일 경우
	// 쓰레드를 추가로 생성하지 않고 고루틴을 대기시킴
	// 코어1개 - os 쓰레드1개 - 고루틴1개
	// 고루틴 하나의 작업이 끝나야 대기 중인 고루틴 실행

	// 시스템콜 호출시
	// 시스템콜이 끝날 때까지 고루틴은 대기 상태로 이동함(스위칭)

	// os 쓰레드를 코어 갯수만큼만 만들고 고루틴을 교체해가며 사용
	// 장점 : 많은 고루틴을 생성해도 컨텍스트 스위칭 비용 발생 X (OS단위에서의)
	// 고루틴이 교체될때 컨텍스트 스위칭이 발생하긴 함 하지만 이 비용이 적게 개발됨 (스택사이즈가 훨씬 작다)

	// 동시성 프로그래밍 (concurrent programming)
	// 동일한 메모리 자원을 여러 고루틴에서 집근할 시 동시성 문제 발생

	var wg2 sync.WaitGroup
	account := &Account{10}
	wg2.Add(10)

	for i := 0; i < 10; i++ {
		go func() {
			for {
				DepositAndWithdraw(account)
			}
			wg2.Done()
		}()
		wg2.Wait()
	}

	// 해결법 : Lock 사용, 뮤텍스(Mutual Exclusion)
	// 항상 하나의 고루틴만 접근하도록 하여 자원 보호
	// 뮤텍스 추가한 후 panic이 발생하지 않는다

	// 뮤텍스 단점
	// 1. 동시성 프로그래밍으로 인한 성능 향상 얻을 수 없음, 과도한 락킹으로 하락할 수 있음
	// 2. 데드락 문제 발생 (고루틴을 완전히 멈추게 만듦)

	rand.Seed(time.Now().UnixNano())

	wg.Add(2)
	fork := &sync.Mutex{}
	spoon := &sync.Mutex{}

	go diningProblem("A", fork, spoon, "포크", "수저")
	go diningProblem("B", spoon, fork, "수저", "포크")
	wg.Wait()
	// 결과 : fatal error: all goroutines are asleep - deadlock!

	var jobList [10]Job
	for i := 0; i < 10; i++ {
		jobList[i] = &SquareJob{i}
	}

	var wg3 sync.WaitGroup
	wg3.Add(10)

	for i := 10; i < 10; i++ {
		job := jobList[i]
		go func() {
			job.Do()
			wg3.Done()
		}()
	}

	wg3.Wait()
}
