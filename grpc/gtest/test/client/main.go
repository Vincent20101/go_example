package main

import (
	"context"
	"fmt"
	"time"

	gproto "github.com/vincent20101/go-example/grpc/gtest/proto"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("127.0.0.1:8080", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	client := gproto.NewGreeterClient(conn)
	ctx, _ := context.WithTimeout(context.Background(), 1*time.Second)
	_, err = client.SayHello(ctx, &gproto.HelloRequest{Name: "bobby"})
	if err != nil {
		//panic(err)
		fmt.Println(err)
	}
	time.Sleep(100 * time.Second)
	//fmt.Println(r.Message)
	//http.StatusTemporaryRedirect

}
