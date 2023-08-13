package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	client := &http.Client{}

	url := "http://localhost:8080/name"
	resp, err := client.Head(url)
	if err != nil {
		fmt.Println("Request error:", err)
		return
	}
	defer resp.Body.Close()
	fmt.Println(resp.Header)
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Read response error:", err)
		return
	}

	fmt.Printf("Server response: %s\n", body)
}
