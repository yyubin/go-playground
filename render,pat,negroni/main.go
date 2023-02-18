package main

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/gorilla/pat"
	"github.com/unrolled/render"
	"github.com/urfave/negroni"
)

var rd *render.Render

type User struct {
	Name      string    `json: "name"`
	Email     string    `json: "email"`
	CreatedAt time.Time `json: "created_at"`
}

func getUserInfoHandler(w http.ResponseWriter, r *http.Request) {
	user := User{Name: "yyubin", Email: "hazing120@gmail.com"}

	rd.JSON(w, http.StatusOK, user)

	// w.Header().Add("Content-type", "application/json")
	// w.WriteHeader(http.StatusOK)

	// data, _ := json.Marshal(user)
	// fmt.Fprint(w, string(data))
}

func addUserHandler(w http.ResponseWriter, r *http.Request) {
	user := new(User)
	err := json.NewDecoder(r.Body).Decode(user)
	if err != nil {
		// w.WriteHeader(http.StatusBadRequest)
		// fmt.Fprint(w, err)
		rd.Text(w, http.StatusBadRequest, err.Error())
		return
	}
	user.CreatedAt = time.Now()

	rd.JSON(w, http.StatusOK, user)

	// w.Header().Add("Content-type", "application/json")
	// w.WriteHeader(http.StatusOK)

	// data, _ := json.Marshal(user)
	// fmt.Fprint(w, string(data))

}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	// tmpl, err := template.New("Hello").ParseFiles("templates/hello.tmpl")
	// if err != nil {
	// 	// w.WriteHeader(http.StatusInternalServerError)

	// 	// fmt.Fprint(w, err)
	// 	rd.Text(w, http.StatusInternalServerError, err.Error())
	// 	return
	// }
	user := User{Name: "yyubin", Email: "hazing120@gmail.com"}
	rd.HTML(w, http.StatusOK, "body", user)
	// tmpl.ExecuteTemplate(w, "hello.tmpl", "Yubin!!!")
}

func main() {
	rd = render.New(render.Options{
		// directory 추가시 사용 default: templates
		Directory:  "template",
		Extensions: []string{".html", ".tmpl"},
		Layout:     "hello",
	})
	mux := pat.New()

	mux.Get("/users", getUserInfoHandler)
	mux.Post("/users", addUserHandler)
	mux.Get("/hello", helloHandler)

	// mux.Handle("/", http.FileServer(http.Dir("public")))
	n := negroni.Classic()
	n.UseHandler(mux) // 기본 로그기능, 파일서브 기능 제공해줌 (데코레이터 패턴사용)

	http.ListenAndServe(":3000", n)
}
