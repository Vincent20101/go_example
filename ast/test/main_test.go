package test

import (
	"go/ast"
	"go/parser"
	"go/token"
	"testing"
)

func TestAst(t *testing.T) {
	src := `package main
func main(){
    println("hello world")
}
    `
	//创建用于解析源文件的对象
	fileset := token.NewFileSet()
	//解析源文件，返回ast.File原始文档类型的结构体。
	f, err := parser.ParseFile(fileset, "", src, 0)
	if err != nil {
		panic(err)
	}
	//查看日志打印
	ast.Print(fileset, f)
}
