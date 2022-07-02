package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancelFunc := context.WithCancel(context.TODO())
	time.AfterFunc(1*time.Second, func() {
		cancelFunc()
	})
	go timeout(ctx)

	time.Sleep(2 * time.Second)

}

func timeout(ctx context.Context) {
	fmt.Println("开始sleep")
	time.Sleep(50 * time.Second)
	fmt.Println("结束了")
}
