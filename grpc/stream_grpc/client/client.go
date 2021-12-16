package main

import (
	"context"
	"fmt"
	"go-example/grpc/stream_grpc/proto"
	"sync"
	"time"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:50052", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	//服务端流模式
	c := gproto.NewGreeterClient(conn)
	res, _ := c.GetStream(context.Background(), &gproto.StreamReqData{Data: "muke"})
	for {
		a, err := res.Recv() //如果大家懂socket编程的话就明白 send recv
		if err != nil {
			fmt.Println(err)
			break
		}
		fmt.Println(a.Data)
	}

	//客户端流模式
	putS, _ := c.PutStream(context.Background())
	i := 0
	for {
		i++
		_ = putS.Send(&gproto.StreamReqData{
			Data: fmt.Sprintf("muke%d", i),
		})
		time.Sleep(time.Second)
		if i > 10 {
			break
		}
	}

	//双向流模式
	allStr, _ := c.AllStream(context.Background())
	wg := sync.WaitGroup{}
	wg.Add(2)
	go func() {
		defer wg.Done()
		for {
			data, _ := allStr.Recv()
			fmt.Println("收到服务端消息：" + data.Data)
		}
	}()

	//1. 集中学习protobuf， grpc

	go func() {
		defer wg.Done()
		for {
			_ = allStr.Send(&gproto.StreamReqData{Data: "muke"})
			time.Sleep(time.Second)
		}
	}()

	wg.Wait()
}
