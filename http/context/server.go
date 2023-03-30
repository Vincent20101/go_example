package main

import (
	"fmt"
	"net/http"
)

func hello(w http.ResponseWriter, r *http.Request) {
	for i := 0; i < 100*10000; i++ {
		fmt.Println(111)
		w.Write([]byte("hello world"))
	}
}

func main() {
	fmt.Println("listening 8888:")
	http.HandleFunc("/hello", hello)
	_ = http.ListenAndServe(":8888", nil)
}
