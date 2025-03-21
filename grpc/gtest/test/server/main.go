package main

import (
	"context"
	"fmt"
	"net"
	"time"

	"google.golang.org/grpc/encoding"

	"github.com/vincent20101/go-example/grpc/gtest/test/codec"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	gproto "github.com/vincent20101/go-example/grpc/gtest/proto"

	"google.golang.org/grpc"
)

type Server struct{}

func (s *Server) SayHello(ctx context.Context, request *gproto.HelloRequest) (*gproto.HelloReply, error) {
	time.Sleep(5 * time.Second)
	if ctx.Err() == context.Canceled {
		return nil, status.Errorf(codes.Canceled, "SearchService.Search canceled")
	}
	if ctx.Err() == context.DeadlineExceeded {
		fmt.Println("DeadlineExceeded")
		return nil, status.Errorf(codes.DeadlineExceeded, "SearchService.Search DeadlineExceeded")
	}
	//time.Sleep(20 * time.Second)
	fmt.Println("test")
	conn, err := grpc.Dial("127.0.0.1:8081", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	//time.Sleep(2 * time.Second)
	client := gproto.NewGreeterClient(conn)
	ctx1, _ := context.WithTimeout(ctx, 30*time.Second)
	//ctx1, _ := context.WithTimeout(context.Background(), 3*time.Second)
	fmt.Println("进来了")
	//time.Sleep(2 * time.Second)
	r, err := client.SayHello(ctx1, &gproto.HelloRequest{Name: "bobby"}, grpc.ForceCodec(&codec.JSONCoder{}))
	if err != nil {
		//panic(err)
		fmt.Println(err)
		fmt.Println("mistaken-------")
		return &gproto.HelloReply{}, nil
	}
	fmt.Println(r.Message)
	return &gproto.HelloReply{
		Message: "Hello " + request.Name,
	}, nil
}

func main() {
	encoding.RegisterCodec(&codec.JSONCoder{})
	g := grpc.NewServer()
	s := Server{}
	gproto.RegisterGreeterServer(g, &s)
	listen, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}
	_ = g.Serve(listen)
}
