package main

import (
	"fmt"
	"runtime"
	"strings"
)

func getGoroutineID() int64 {
	buf := make([]byte, 64)
	n := runtime.Stack(buf, false)
	gidField := strings.Fields(strings.TrimSpace(string(buf[:n])))[1]
	gidStr := strings.TrimPrefix(gidField, "goroutine")
	var gid int64
	fmt.Sscanf(gidStr, "%d", &gid)
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
