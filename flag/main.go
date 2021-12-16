package main

import (
	"flag"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"testing"
)

var dir string

func init() {
	flag.StringVar(&dir, "dir", "", "请输入一个目录路径")
	//原因：flag.Parse() 把go test的test当作一个参数处理，而flag没有对这个参数做处理导致解析失败。
	//解决办法：
	//方法1、把flag.Parse()放到main函数中
	//方法2、在flag.Parse()前加上testing.Init()方法
	testing.Init()
	flag.Parse()

}
func main() {
	if !checkDir(dir) {
		os.Exit(1)
	}
	err := filepath.Walk(dir, func(path string, info fs.FileInfo, err error) error {
		fmt.Println("path=>", path, "isDir=>", info.IsDir(), "name=>", info.Name())
		return nil
	})
	if err != nil {
		return
	}
}

func checkDir(dir string) bool {
	return true
}
