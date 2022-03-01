package main

import (
	"log"
	"net/http"
)

func main() {

	fs := http.FileServer(http.Dir("F:\\"))
	http.Handle("/static/", http.StripPrefix("/static", fs))
	log.Fatal(http.ListenAndServe(":8080", nil))

}
