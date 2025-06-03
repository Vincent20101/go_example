package main

import (
	"context"
	"fmt"
	"net"
	"reflect"
	"time"

	"github.com/vincent20101/go-example/grpc/gtest/test/codec"

	gproto "github.com/vincent20101/go-example/grpc/gtest/proto"

	"google.golang.org/grpc"
)

func main() {
	// var dos []grpc.DialOption
	do, l, r := NewDialOpt("pfcp")
	// dos = append(dos, do, grpc.WithInsecure(), grpc.WithBlock())
	conn, err := grpc.Dial("127.0.0.1:8080", do, grpc.WithInsecure() /*, grpc.WithBlock(), grpc.MaxConcurrentStreams()*/)
	// time.Sleep(5 * time.Second)
	fmt.Println("lhb:", l(), r())
	if !reflect.DeepEqual(l(), reflect.Zero(reflect.TypeOf(l())).Interface()) {
		fmt.Println(l())
	}
	if !reflect.DeepEqual(r(), reflect.Zero(reflect.TypeOf(r())).Interface()) {
		fmt.Println(r())
	}
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	client := gproto.NewGreeterClient(conn)
	//fmt.Println("waiting for...")
	//time.Sleep(time.Second * 10)
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	go func() {
		time.Sleep(time.Second)
		cancel()
	}()
	rsp, err := client.SayHello(ctx, &gproto.HelloRequest{Name: "bobby"}, grpc.ForceCodec(&codec.JSONCoder{}))
	if err != nil {
		panic(err)
		//fmt.Println(err)
	}
	time.Sleep(time.Second * 2)
	fmt.Println("rsp:", rsp.Message)
	//time.Sleep(100 * time.Second)
	//fmt.Println(r.Message)
	//http.StatusTemporaryRedirect

}

type RemoteAddr struct {
	Ipv4 net.IP
	Ipv6 net.IP
	Port uint16
}

func NewDialOpt(serviceName string) (do grpc.DialOption, getRemote, getLocal func() RemoteAddr) {
	// r, l = &RemoteAddr{}, &RemoteAddr{}
	var r, l RemoteAddr
	do = grpc.WithContextDialer(func(ctx context.Context, s string) (net.Conn, error) {
		d := net.Dialer{}
		c, e := d.DialContext(ctx, "tcp", s)
		if e != nil {
			return nil, e
		}

		// 修改闭包内的 r 和 l
		if c.LocalAddr() != nil {
			if localAddr, ok := c.LocalAddr().(*net.TCPAddr); ok {
				if localAddr.IP.To4() != nil && len(l.Ipv4) == 0 {
					l.Ipv4 = localAddr.IP
				} else if localAddr.IP.To16() != nil && localAddr.IP.To4() == nil && len(l.Ipv6) == 0 {
					l.Ipv6 = localAddr.IP
				}
			}
		}

		if c.RemoteAddr() != nil {
			if remoteAddr, ok := c.RemoteAddr().(*net.TCPAddr); ok {
				if remoteAddr.IP.To4() != nil && len(r.Ipv4) == 0 {
					r.Ipv4 = remoteAddr.IP
				} else if remoteAddr.IP.To16() != nil && remoteAddr.IP.To4() == nil && len(r.Ipv6) == 0 {
					r.Ipv6 = remoteAddr.IP
				}
			}
		}

		return c, nil
	})

	// 返回闭包函数来获取 r 和 l
	getRemote = func() RemoteAddr { return r }
	getLocal = func() RemoteAddr { return l }

	return
}
