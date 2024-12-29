package controller

import "net/http"

func NewHelloWorldController() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World"))
	}
}