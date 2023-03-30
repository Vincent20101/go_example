package main

import (
	"context"
	"fmt"
	"time"
)

type paramKey struct{}
type student struct {
	name string
	age  int
}

func main() {
	c := context.WithValue(context.Background(),
		paramKey{}, student{
			name: "lhb",
			age:  11,
		})
	c, cancel := context.WithTimeout(c, 3*time.Second)
	defer cancel()
	go mainTask(c)

	var cmd string
	for {
		fmt.Scan(&cmd)
		if cmd == "c" {
			cancel()
		}
	}
}

func mainTask(c context.Context) {
	fmt.Printf("main task started with param %v\n", c.Value(paramKey{}))
	go func() {
		c1, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		smallTask(c1, "background_task", 9*time.Second)
	}()
	go func() {
		c1, cancel := context.WithTimeout(c, 10*time.Second)
		defer cancel()
		smallTask(c1, "sub_task", 1*time.Second)
	}()
	smallTask(c, "same_task", 1*time.Second)
}

func smallTask(c context.Context, name string, d time.Duration) {
	fmt.Printf("%s started with param %v\n", name, c.Value(paramKey{}))
	select {
	case <-time.After(d):
		time.Sleep(10 * time.Second)
		fmt.Printf("%s done\n", name)
	case <-c.Done():
		fmt.Printf("%s cancelled\n", name)
	}
}
