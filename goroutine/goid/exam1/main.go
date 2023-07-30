package main

import (
	"fmt"
	"runtime"
	"strings"
)

func getGoroutineID() int64 {
	buf := make([]byte, 64)
	n := runtime.Stack(buf, false)
	fmt.Println(strings.TrimSpace(string(buf[:n])))
	gidField := strings.Fields(strings.TrimSpace(string(buf[:n])))[1]
	fmt.Println(gidField)
	var gid int64
	sscanf, err := fmt.Sscanf(gidField, "%d", &gid)
	fmt.Println(sscanf, err)
	if err != nil {
		fmt.Println("sdfsdf")
	}
	return gid
}

func main() {
	goroutineID := getGoroutineID()
	fmt.Printf("Current goroutine ID: %d\n", goroutineID)

	go func() {
		newGoroutineID := getGoroutineID()
		fmt.Printf("New goroutine ID: %d\n", newGoroutineID)
	}()

	fmt.Println("Main goroutine sleeping...")
	runtime.Gosched()
	fmt.Println("Main goroutine finished.")
}
