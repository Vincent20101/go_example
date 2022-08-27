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
	fmt.Println(context.Background().Done())
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	go func(ctx context.Context) {
		fmt.Println("开始sleep")
		time.Sleep(5 * time.Second)
		select {
		case v, ok := <-ctx.Done():
			fmt.Println(v, ok)
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
