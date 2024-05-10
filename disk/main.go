package main

import (
	"fmt"
	"syscall"
	"time"
)

func main() {
	fmt.Println(time.Now().UnixMilli())
	fs := syscall.Statfs_t{}
	err := syscall.Statfs("/", &fs)
	if err != nil {
		panic(err)
	}

	// 计算总空间大小 (单位为MB)
	totalSpace := (fs.Blocks * uint64(fs.Bsize)) / (1024 * 1024)

	// 计算可用空间 (单位为MB)
	freeSpace := (fs.Bfree * uint64(fs.Bsize)) / (1024 * 1024)

	// 计算已使用空间 (单位为MB)
	usedSpace := totalSpace - freeSpace

	fmt.Printf("Total space: %d MB\n", totalSpace)
	fmt.Printf("Free space: %d MB\n", freeSpace)
	fmt.Printf("Used space: %d MB\n", usedSpace)
}
