package main

import (
	"html/template"
	"os"
)

type User struct {
	Name  string
	Email string
	Age   int
}

func (u User) IsOld() bool {
	return u.Age > 30
}

func main() {
	user := User{Name: "yubin", Email: "hazing120@gmail.com", Age: 23}
	user2 := User{Name: "aaa", Email: "aaa@gmail.com", Age: 40}

	users := []User{user, user2}

	tmpl, err := template.New("Tmpl1").ParseFiles("templates/tmpl1.tmpl", "templates/tmpl2.tmpl")
	if err != nil {
		panic(err)
	}

	// Parse 사용시
	// tmpl.Execute(os.Stdout, user)
	// tmpl.Execute(os.Stdout, user2)

	// ParseFilses 사용시
	// tmpl.ExecuteTemplate(os.Stdout, "tmpl2.tmpl", user)
	// tmpl.ExecuteTemplate(os.Stdout, "tmpl2.tmpl", user2)

	// list로 사용시
	tmpl.ExecuteTemplate(os.Stdout, "tmpl2.tmpl", users)

}
