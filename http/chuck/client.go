package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {
	Client()
}

func Client() {
	resp, err := http.Get("http://127.0.0.1:8080/name")
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	fmt.Println("transferEncoding:", resp.TransferEncoding)

	reader := bufio.NewReader(resp.Body)
	for {
		line, err := reader.ReadString('\n')
		if len(line) > 0 {
			fmt.Println("line: ", line)
		}
		fmt.Println("err: ", err)
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
	}
}
