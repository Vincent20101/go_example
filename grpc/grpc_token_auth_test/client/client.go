package main

import (
	"context"
	"fmt"
	"time"

	"github.com/vincent20101/go-example/grpc/grpc_token_auth_test/proto"
	"google.golang.org/grpc"
)

type customCredential struct{}

func (c customCredential) GetRequestMetadata(ctx context.Context, uri ...string) (map[string]string, error) {
	return map[string]string{
		"appid":  "101010",
		"appkey": "i am key",
	}, nil
}

// RequireTransportSecurity indicates whether the credentials requires
// transport security.
func (c customCredential) RequireTransportSecurity() bool {
	return false
}

func main() {
	//stream
	//interceptor := func(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error{
	//	start := time.Now()
	//	//md := metadata.New(map[string]string{
	//	//	"appid":"10101",
	//	//	"appkey":"i am key",
	//	//})
	//	ctx = metadata.NewOutgoingContext(context.Background(), md)
	//	err := invoker(ctx, method, req, reply, cc, opts...)
	//	fmt.Printf("耗时：%s\n", time.Since(start))
	//	return err
	//}
	for i := 0; i < 5; i++ {
		grpc.WithPerRPCCredentials(customCredential{})
		var opts []grpc.DialOption
		opts = append(opts, grpc.WithInsecure())
		opts = append(opts, grpc.WithPerRPCCredentials(customCredential{}))
		conn, err := grpc.Dial("127.0.0.1:12345", opts...)
		fmt.Println("lianjie: ", conn)
		if err != nil {
			panic(err)
		}
		defer conn.Close()

		c := proto.NewGreeterClient(conn)
		r, err := c.SayHello(context.Background(), &proto.HelloRequest{Name: "linhuanbo1"})
		if err != nil {
			panic(err)
		}
		fmt.Println(r.Message)
		time.Sleep(time.Second * 3)
	}

}
