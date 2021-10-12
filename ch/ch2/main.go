package main

import (
	"bufio"
	"fmt"
	"os"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {

	// time.Sleep(1e9) // 1 * 10 的 9次方
	// // time.Sleep(time.Second * 2) // 延迟2s
	// fmt.Println("hello world")
	// fmt.Println(time.Now())

	// go 主进序也是一个协程 -> 以协程方式切换
	// 通过关键词 go声明方法也是一个协程
	fmt.Println(time.Now())
	// WriteFile("D:\\phpStudy\\PHPTutorial\\WWW\\go\\src\\go2011\\day10\\t1.txt", "shineyork 666!")
	wg.Add(1) // 登记协程
	go redis()
	wg.Add(1) // 登记协程、
	go mysql()
	// 进程 -》 单线程-》协程 ；

	// 进程 -》协程
	// 通道
	// go file()
	wg.Wait() // 等待登记的协程执行完再执行后续的程序

	// time.Sleep(time.Second * 4)
	fmt.Println(time.Now())
}
func redis() {
	defer wg.Done() // 结束登记
	fmt.Println("start 读取redis中的信息")
	WriteFile("D:\\t.txt", "shineyork 666!")
	// time.Sleep(time.Second * 1)
	fmt.Println("end 读取redis中的信息")
}

func mysql() {
	defer wg.Done() // 结束登记
	fmt.Println("start 读取MySQL中的信息")
	time.Sleep(time.Second * 2)
	fmt.Println("end 读取MySQL中的信息")
}
func file() {
	fmt.Println("start 从文件中获取信息")
	time.Sleep(time.Second * 3)
	fmt.Println("end 从文件中获取信息")
}

func WriteFile(path, data string) (bool, error) {
	// 打开文件
	// 0666 是文件的写入权限
	outputFile, outputError := os.OpenFile(path, os.O_WRONLY|os.O_CREATE, 0666)

	if outputError != nil {
		return false, outputError
	}
	defer outputFile.Close()
	// 创建写的缓冲区
	outputWriter := bufio.NewWriter(outputFile)
	// 写入信息
	outputWriter.WriteString(data)
	// 刷新
	outputWriter.Flush()
	return true, nil
}
