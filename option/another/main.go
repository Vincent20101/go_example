package main

import (
	"fmt"
)

//如何向func传递默认值

type dialOption struct {
	Username string
	Password string
	Service  string
}

type DialOption interface {
	apply(*dialOption)
}

type funcOption struct {
	f func(*dialOption)
}

func (fdo *funcOption) apply(do *dialOption) {
	fdo.f(do)
}

func newFuncOption(f func(*dialOption)) *funcOption {
	return &funcOption{
		f: f,
	}
}

func withUserName(s string) DialOption {
	return newFuncOption(func(o *dialOption) {
		o.Username = s
	})
}

func withPasswordd(s string) DialOption {
	return newFuncOption(func(o *dialOption) {
		o.Password = s
	})
}

func withService(s string) DialOption {
	return newFuncOption(func(o *dialOption) {
		o.Service = s
	})
}

//默认参数
func defaultOptions() dialOption {
	return dialOption{
		Service: "test",
	}
}

type clientConn struct {
	timeout int
	dopts   dialOption
}

func NewClient(address string, opts ...DialOption) {
	cc := &clientConn{
		timeout: 30,
		dopts:   defaultOptions(),
	}
	//循环调用opts
	for _, opt := range opts {
		opt.apply(&cc.dopts)
	}

	fmt.Printf("%+v\n", cc.dopts)
	//var i = math.MaxInt32
	//fmt.Println(i)
	//fmt.Println(16 << 10)
	//bb := []byte{1, 2, 3, 4, 5, 6, 76, 4, 6, 7, 8, 9, 0}
	//fmt.Println(bb[:0])
	//_ = append(bb[:0], []byte{6, 7, 8, 9}...)
	//fmt.Println(bb)
	//fmt.Println(8 ^ 5)
	//fmt.Println(1 << 31)
	//
	//fmt.Println(struct {
	//	b bytes.Buffer
	//}{})
}

func main() {
	NewClient("127.0.0.1", withPasswordd("654321"), withService("habox"))
	NewClient("127.0.0.1", withService("habox"))
}
