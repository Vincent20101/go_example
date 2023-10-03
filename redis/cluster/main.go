package main

import (
	"context"
	"fmt"
	"log"

	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()

func main() {
	// 创建一个连接池
	pool := redis.NewPool(func() (redis.Conn, error) {
		return redis.Dial("tcp", "localhost:6379")
	}, 10) // 10个最大连接数

	defer pool.Close()

	// 创建一个连接
	conn := pool.Get()
	defer conn.Close()

	key := "example_key"

	// 使用连接执行命令
	val, err := conn.Do("GET", key)
	if err != nil {
		// 处理MOVED错误
		if movedErr, ok := err.(*redis.CmdError); ok && movedErr.Err == "MOVED" {
			parts := movedErr.Args
			slot := parts[1]
			targetAddr := parts[2]
			fmt.Printf("Key moved to slot %s, new target: %s\n", slot, targetAddr)

			// 更新连接信息
			conn.Close()
			conn = pool.Get()
			conn.Do("SET", key, "new_value") // 在新连接上重试操作
		} else {
			log.Fatal("Error:", err)
		}
	} else {
		// 处理正常情况下的响应
		fmt.Printf("%s: %s\n", key, val)
	}
}
