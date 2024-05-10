package main

import (
	"log"
	"net/http"
)

type sayHandler struct {
}

func (h *sayHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello"))

}

func main01() {
	log.Fatal(http.ListenAndServe(":8080", new(sayHandler)))
}

func say(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	// 再次写入响应头，导致冗余调用
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("hello world"))
}

func main() {
	http.HandleFunc("/hello", say)
	http.Handle("test", new(sayHandler))
	log.Fatal(http.ListenAndServe(":8080", nil))
}
