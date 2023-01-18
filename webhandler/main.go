package main

import (
	"goweb2/webhandler/myapp"
	"net/http"
)

func NewHandler() http.Handler {
	mux := myapp.NewHandler()
	return mux
}

func main() {
	mux := NewHandler()

	http.ListenAndServe(":3000", mux)
}
