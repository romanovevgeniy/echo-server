package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "hello")
}

func main() {
	http.HandleFunc("/hello", handler)
	fmt.Println("Сервер запущен")
	http.ListenAndServe(":1234", nil)
}