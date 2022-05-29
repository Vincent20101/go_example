package main

import (
	"github.com/davecgh/go-spew/spew"
	"github.com/pkg/errors"
)

func main() {
	i := 0
	s := "哈哈"
	m := map[int]string{1: "1", 2: "2"}
	e := errors.New("嘿嘿，错误")
	p := persons{Name: "张三"}
	spew.Dump(i, s, m, e, p)
}

type persons struct {
	Name string
}
