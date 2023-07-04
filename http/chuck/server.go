package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	Server()
}

func Server() {
	http.HandleFunc("/name", func(w http.ResponseWriter, r *http.Request) {
		flusher := w.(http.Flusher)
		for i := 0; i < 2; i++ {
			fmt.Fprintf(w, "daniel Hu\n")
			flusher.Flush()
			<-time.Tick(time.Second * 1)
		}
	})
	log.Fatal(http.ListenAndServe(":8080", nil))
}
