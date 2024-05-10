package main

import (
	"fmt"
	"log"

	"github.com/fsnotify/fsnotify"
)

func main() {
	// 创建一个新的监视器对象
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}
	defer watcher.Close()

	done := make(chan bool)
	// 在goroutine中处理事件
	go func() {
		for {
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					return
				}
				// 打印事件信息
				fmt.Println("event:", event)
				if event.Op&fsnotify.Write == fsnotify.Write {
					fmt.Println("modified file:", event.Name)
				}
			case err, ok := <-watcher.Errors:
				if !ok {
					return
				}
				log.Println("error:", err)
			}
		}
	}()

	// 添加一个目录到监视器中，以便监视其事件
	err = watcher.Add("/root/project/goex/go_example/fsnotify/config.json")
	err = watcher.Add("/root/project/goex/go_example/fsnotify/configs.json")
	if err != nil {
		log.Fatal(err)
	}
	<-done
}
