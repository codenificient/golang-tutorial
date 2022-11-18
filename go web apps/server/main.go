package main

import (
	"net/http"
)

func helloWorldPage(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/url":
		url(w, r)

	case "/body":
		body(w, r)

	default:
		w.Write([]byte("Welcome to the HomePage!"))
	}
}

func main() {
	http.HandleFunc("/", helloWorldPage)
	http.ListenAndServe(":8080", nil)
}
