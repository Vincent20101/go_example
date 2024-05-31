package main

import (
	"log"
	"runtime"
)

func init() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)
}

//go1.2版本前后表现不同
const N = 1000000

func main() {
	printMemStats("初始情况")

	m := make(map[int][128]byte, N)
	for i := 0; i < N; i++ {
		m[i] = [128]byte{}
	}

	runtime.GC()
	log.Println("删除前Map长度：", len(m))
	printMemStats("添加100万个键值对后")
	for i := 0; i < N/2; i++ {
		delete(m, i)
	}
	log.Println("删除后Map长度：", len(m))

	// 再次进行手动GC回收
	runtime.GC()
	printMemStats("删除50万个键值对后")

	for i := N / 2; i < N; i++ {
		delete(m, i)
	}
	runtime.GC()
	printMemStats("删除全部键值对后")

	// 设置为nil进行回收
	m = nil
	runtime.GC()
	printMemStats("设置为nil后")
}

func printMemStats(mag string) {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	log.Printf("%v：分配的内存 = %vKB, GC的次数 = %v\n", mag, m.Alloc/1024, m.NumGC)
}
