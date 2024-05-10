package main

import (
	"fmt"
	"runtime"

	"github.com/shirou/gopsutil/v3/mem"
)

var (
	mbSize = 1024 * 1024
)

func main() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)

	fmt.Println(m.Alloc)
	fmt.Println(m.TotalAlloc / 1024)
	fmt.Println(m.Frees)
	fmt.Println(m.DebugGC)
	fmt.Println(m.NumGC)
	fmt.Println(m.HeapInuse)
	fmt.Println(m.HeapIdle)
	fmt.Println(m.Sys)
	// 计算已使用内存的比例
	usedPercent := float64(m.Alloc) / float64(m.TotalAlloc) * 100

	// 判断是否超过80%
	if usedPercent >= 80 {
		fmt.Printf("Memory usage is %.2f%%, exceeds 80%%\n", usedPercent)
	} else {
		fmt.Printf("Memory usage is %.2f%%\n", usedPercent)
	}

	getMemInfo()
}

func getMemInfo() {
	memInfo, err := mem.VirtualMemory()
	if err != nil {
		fmt.Println("get memory info fail. err： ", err)
	}
	// 获取总内存大小，单位GB
	memTotal := memInfo.Total / 1024 / 1024 / 1024
	// 获取已用内存大小，单位MB
	memUsed := memInfo.Used / 1024 / 1024
	// 可用内存大小
	memAva := memInfo.Available / 1024 / 1024
	// 内存可用率
	memUsedPercent := memInfo.UsedPercent
	fmt.Printf("总内存: %v GB, 已用内存: %v MB, 可用内存: %v MB, 内存使用率: %.3f %% \n", memTotal, memUsed, memAva, memUsedPercent)
}
