package myapp

import "net/http"

func NewHandler() http.Handler {
	mux := http.NewServeMux()
	return mux
}
