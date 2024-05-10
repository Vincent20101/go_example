package main

import (
	"fmt"

	"github.com/vincent20101/go-example/buf/pkg"
)

type stu struct {
	age  int
	name string
}

func main() {
	var b []byte
	fmt.Println(b)
	b = nil
	fmt.Println(b[:])

	fmt.Println(pkg.PACKAGE_PATH)

	var s stu
	getStu(&s)
	fmt.Println(s)
}

func getStu(s *stu) {
	s.age = 18
	s.name = "lhb"
}
