package main

import (
	"fmt"
	"html"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

func startWebserver() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
	})

	go http.ListenAndServe(":8000", nil)

}

func startLoadTest() {
	count := 0
	for {
		if count == 6 {
			return
		}
		resp, err := http.Get("http://localhost:8000/")
		if err != nil {
			panic(fmt.Sprintf("Got error: %v", err))
		}
		io.Copy(ioutil.Discard, resp.Body) // <-- add this line
		resp.Body.Close()
		log.Printf("Finished GET request #%v", count)
		count += 1
	}

}

func main() {

	startWebserver()

	startLoadTest()

}
