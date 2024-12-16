package main

import (
	"fmt"
	"html"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func startWebserver() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(r.URL.Host)
		fmt.Println(r.URL.Path)
		fmt.Println(r.Host)
		fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Host))
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
	time.Sleep(time.Second * 10)
	startLoadTest()

}
