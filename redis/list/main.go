package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"

	"github.com/go-redis/redis"
)

const (
	queueKey    = "failed_requests" // Redis队列的键名
	maxQueueLen = 1000              // 队列最大长度
)

func main() {
	// 创建一个Redis客户端实例
	redisClient := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // 如果需要密码
		DB:       0,  // 使用默认数据库
	})

	// 创建一个HTTP客户端
	httpClient := &http.Client{
		Timeout: time.Second * 10,
	}

	// 定义要发送的URL列表
	urls := []string{
		"https://example.com/request1",
		"https://example.com/request2",
		"https://example.com/request3",
	}

	// 使用等待组来等待所有Worker完成
	var wg sync.WaitGroup

	// 启动Worker处理Redis队列消息
	go func() {
		for {
			// 从Redis队列中获取请求URL
			url, err := redisClient.LPop(queueKey).Result()
			if err != nil {
				// 处理错误
				if err == redis.Nil {
					// 队列为空，等待一段时间后重试
					time.Sleep(time.Second * 5)
					continue
				}
				fmt.Printf("从Redis队列中获取请求失败：%s\n", err.Error())
				continue
			}

			// 处理请求URL
			fmt.Printf("从Redis队列中获取请求：%s\n", url)
			// ...

			// 将处理任务添加到等待组
			wg.Add(1)

			// 启动goroutine处理请求URL
			go func(url string) {
				defer wg.Done()

				retries := 3

				for retries > 0 {
					resp, err := httpClient.Get(url)
					if err != nil {
						fmt.Printf("请求失败：%s\n", err.Error())

						// 重试计数减少
						retries--
						continue
					}

					// 检查响应状态码
					if resp.StatusCode == http.StatusOK {
						// 请求成功，处理响应数据
						// ...

						resp.Body.Close()
						break
					}

					resp.Body.Close()

					// 重试计数减少
					retries--
				}

				// 如果重试仍然失败，则将请求URL放入Redis中
				if retries == 0 {
					// 将请求URL放入Redis队列尾部
					err := redisClient.LPush(queueKey, url).Err()
					if err != nil {
						fmt.Printf("将请求放入Redis失败：%s\n", err.Error())
					} else {
						fmt.Printf("将请求放入Redis成功：%s\n", url)
					}

					// 获取当前队列长度
					queueLen := redisClient.LLen(queueKey).Val()
					// 检查队列是否超过最大长度
					if queueLen > maxQueueLen {
						// 移除最早的消息
						err := redisClient.RPop(queueKey).Err()
						if err != nil {
							fmt.Printf("移除队列中最早的消息失败：%s\n", err.Error())
						}
					}
				}
			}(url)
		}
	}()

	// 发送初始请求
	for _, url := range urls {
		wg.Add(1)

		go func(url string) {
			defer wg.Done()

			retries := 3

			for retries > 0 {
				resp, err := httpClient.Get(url)
				if err != nil {
					fmt.Printf("请求失败：%s\n", err.Error())

					// 重试计数减少
					retries--
					continue
				}

				// 检查响应状态码
				if resp.StatusCode == http.StatusOK {
					// 请求成功，处理响应数据
					// ...

					resp.Body.Close()
					break
				}

				resp.Body.Close()

				// 重试计数减少
				retries--
			}

			// 如果重试仍然失败，则将请求URL放入Redis中
			if retries == 0 {
				// 将请求URL放入Redis队列尾部
				err := redisClient.LPush(queueKey, url).Err()
				if err != nil {
					fmt.Printf("将请求放入Redis失败：%s\n", err.Error())
				} else {
					fmt.Printf("将请求放入Redis成功：%s\n", url)
				}

				// 获取当前队列长度
				queueLen := redisClient.LLen(queueKey).Val()

				// 检查队列是否超过最大长度
				if queueLen > maxQueueLen {
					// 移除最早的消息
					err := redisClient.RPop(queueKey).Err()
					if err != nil {
						fmt.Printf("移除队列中最早的消息失败：%s\n", err.Error())
					}
				}
			}
		}(url)
	}

	// 等待所有Worker完成
	wg.Wait()

	// 关闭Redis客户端连接
	redisClient.Close()
}
