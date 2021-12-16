package main

import (
	"fmt"
	"go-example/grpc/stream_grpc/proto"
	"google.golang.org/grpc"
	"net"
	"sync"
	"time"
)

const PORT = ":50052"

type server struct {
}

func (s *server) GetStream(req *gproto.StreamReqData, res gproto.Greeter_GetStreamServer) error {
	i := 0
	for {
		i++
		_ = res.Send(&gproto.StreamResData{
			Data: fmt.Sprintf("%v", time.Now().Unix()),
		})
		time.Sleep(time.Second)
		if i > 10 {
			break
		}
	}

	return nil
}

func (s *server) PutStream(cliStr gproto.Greeter_PutStreamServer) error {
	for {
		if a, err := cliStr.Recv(); err != nil {
			fmt.Println(err)
			break
		} else {
			fmt.Println(a.Data)
		}
	}

	return nil
}

func (s *server) AllStream(allStr gproto.Greeter_AllStreamServer) error {
	wg := sync.WaitGroup{}
	wg.Add(2)
	go func() {
		defer wg.Done()
		for {
			data, _ := allStr.Recv()
			fmt.Println("收到客户端消息：" + data.Data)
		}
	}()

	go func() {
		defer wg.Done()
		for {
			_ = allStr.Send(&gproto.StreamResData{Data: "我是服务器"})
			time.Sleep(time.Second)
		}
	}()

	wg.Wait()
	return nil
}

func main() {
	lis, err := net.Listen("tcp", PORT)
	if err != nil {
		panic(err)
	}
	s := grpc.NewServer()
	gproto.RegisterGreeterServer(s, &server{})
	err = s.Serve(lis)
	if err != nil {
		panic(err)
	}
}
