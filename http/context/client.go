package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

func main() {

	client := http.Client{}
	request, err := http.NewRequest(http.MethodPost, "http://127.0.0.1:8888/hello", nil)
	ctx, cancelFunc := context.WithCancel(request.Context())
	request = request.WithContext(ctx)
	if err != nil {
		return
	}
	response, err := client.Do(request)
	if err != nil {
		log.Fatal(err)
	}
	cache := make([]byte, 128)
	timer := time.NewTimer(time.Millisecond)
	go func() {
		select {
		case <-timer.C:
			cancelFunc()
		}
	}()
	for {
		read, err := response.Body.Read(cache)
		if err == nil {
			fmt.Println(string(cache[:read]))
			continue
		}
		if err == io.EOF {
			fmt.Println(string(cache[:read]))
			break
		}
		log.Fatal(err)
	}

}
