package main

import (
	"fmt"
	"sync"
	"time"
)

// 채널 : 고루틴간의 메세지 큐
// Thread safe queue
// 멀티쓰레드 환경에서 lock 없이 사용가능한 큐

type Car struct {
	Body  string
	Tire  string
	Color string
}

var startTime = time.Now()
var wg sync.WaitGroup

func main() { // 메인 고루틴
	// var wg sync.WaitGroup
	// ch := make(chan int)

	// wg.Add(1)
	// go square(&wg, ch)

	// ch <- 9 // 다른 고루틴끼리 메세지 전달할 때 사용
	// wg.Wait()

	// 채널 크기
	// 기본 크기는 0 make(chan int)
	// 수신자가 올 때까지 (데이터를 가져갈 때까지) 기다림

	// 크기 설정 : make(chan int, 2) // 크기 2로 설정
	// ch1 := make(chan int, 2)
	// go square()
	// ch1 <- 9
	// fmt.Println("Never Print")
	// 크기 지정을 안하면 sleep 만 무한히 출력
	// 크기 지정하면 수신자가 올때까지 기다리지 않기 때문에 바로 Never Printe 후 프로그램 종료됨

	// 채널에서 데이터 대기
	// var wg sync.WaitGroup
	// ch := make(chan int)

	// wg.Add(1)
	// go square(&wg, ch)

	// for i := 0; i < 10; i++ {
	// 	ch <- i * 2 // 데이터 10번 집어넣기
	// }
	// //close(ch)
	// wg.Wait() // 메인 고루틴은 여기에서 멈춰있음
	// 모든 고루틴이 멈춰 있기 때문에 데드락
	// 좀비 고루틴 : 채널을 닫아주지 않아서 무한 대기를 하는 고루틴을 좀비 고루틴, 고루틴 릭이라고 함
	// 해결법 : close()로 채널 닫아주기

	// select구문 : 여러 채널에서 동시에 데이터를 기다릴 때 사용
	// close를 사용하지 않는다면 select 구문활용하여 정지할 수 있다

	// 채널로 생산자/소비자 패턴 구현
	tireCh := make(chan *Car)
	paintCh := make(chan *Car)

	fmt.Printf("Start factory\n")

	wg.Add(3)
	go MakeBody(tireCh)
	go InstallTire(tireCh, paintCh)
	go PaintCar(paintCh)

	wg.Wait()
	fmt.Println("Close the factory")

}

func MakeBody(tireCh chan *Car) {
	tick := time.Tick(time.Second)
	after := time.After(10 * time.Second)
	for {
		select {
		case <-tick:
			car := &Car{}
			car.Body = "Sports car"
			tireCh <- car
		case <-after:
			close(tireCh)
			wg.Done()
			return
		}
	}
}

func InstallTire(tireCh, paintCh chan *Car) {
	for car := range tireCh {
		time.Sleep(time.Second)
		car.Tire = "Winter tire"
		paintCh <- car
	}
	wg.Done()
	close(paintCh)
}

func PaintCar(paintCh chan *Car) {
	for car := range paintCh {
		time.Sleep(time.Second)
		car.Color = "Red"
		duration := time.Now().Sub(startTime)
		fmt.Printf("%.2f Complete Car : %s %s %s\n", duration.Seconds(), car.Body, car.Tire, car.Color)
	}
	wg.Done()
}

// func square(wg *sync.WaitGroup, ch chan int) { // 다른 고루틴
// 	n := <-ch
// 	time.Sleep(time.Second)
// 	fmt.Print("Square: ", n*n)
// 	wg.Done()
// }

// func square() {
// 	for {
// 		time.Sleep(2 * time.Second)
// 		fmt.Println("sleep")
// 	}
// }

// func square(wg *sync.WaitGroup, ch chan int) {
// 	for n := range ch { // 무한히 데이터를 뽑으려고 함, 무한 대기, done으로 넘어가지 않음
// 		fmt.Println("Square : ", n*n)
// 		time.Sleep(time.Second)
// 	}
// 	wg.Done()
// }

func square(wg *sync.WaitGroup, ch chan int) {
	tick := time.Tick(time.Second) //tick도 채널
	terminate := time.After(10 * time.Second)

	for {
		select {
		case <-tick:
			fmt.Println("Tick")
		case <-terminate:
			fmt.Println("Terminated")
			wg.Done()
			return
		case n := <-ch:
			fmt.Println("Squqre: ", n*n)
			time.Sleep(time.Second)
		}
	}
}
