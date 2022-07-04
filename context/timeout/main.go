package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	//ctx, cancelFunc := context.WithCancel(context.TODO())
	//time.AfterFunc(1*time.Second, func() {
	//	cancelFunc()
	//})
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	go func(ctx context.Context) {
		fmt.Println("开始sleep")
		time.Sleep(5 * time.Second)
		select {
		case <-ctx.Done():
			return
		default:
		}
		fmt.Println("结束了")
	}(ctx)

	time.Sleep(15 * time.Second)

}

func timeout(ctx context.Context) {
	fmt.Println("开始sleep")
	time.Sleep(5 * time.Second)
	fmt.Println("结束了")
}
