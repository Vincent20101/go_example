package main

import (
	"context"
	"fmt"
	"time"

	"github.com/vincent20101/go-example/grpc/metadata_test/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

func main() {
	//stream
	conn, err := grpc.Dial("127.0.0.1:50051", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	c := proto.NewGreeterClient(conn)

	//md := metadata.Pairs("timestamp", time.Now().Format("2006-01-02"))
	md := metadata.New(map[string]string{
		"name":    "bobby",
		"pasword": "imooc",
	})
	ctx1 := context.WithValue(context.Background(), "a", "b")
	ctx := metadata.NewOutgoingContext(ctx1, md)
	ctx, _ = context.WithTimeout(ctx, time.Second)
	r, err := c.SayHello(ctx, &proto.HelloRequest{Name: "bobby"})
	fmt.Println("a的值为：", ctx.Value("a"))
	if err != nil {
		panic(err)
	}
	fmt.Println(r.Message)
}
