package main

import "fmt"

func main() {
	InitESClient("test", []string{"http://127.0.0.1:9200"}, "admin", "admin",
		WithQueryLogEnable(true), WithSlowQueryLogMillisecond(1000),
		WithScheme("http"))

}

type option struct {
	QueryLogEnable       bool
	SlowQueryMillisecond int64
	Scheme               string
}

type Option func(*option)

func WithQueryLogEnable(enable bool) Option {
	return func(opt *option) {
		opt.QueryLogEnable = enable
	}
}

func WithSlowQueryLogMillisecond(slowQueryMillisecond int64) Option {
	return func(opt *option) {
		opt.SlowQueryMillisecond = slowQueryMillisecond
	}
}

func WithScheme(scheme string) Option {
	return func(opt *option) {
		opt.Scheme = scheme
	}
}

func InitESClient(clientName string, urls []string, username string, password string, options ...Option) {

	//获取传入的配置
	opt := &option{}
	for _, f := range options {
		if f != nil {
			f(opt)
		}
	}
	fmt.Println(opt)
}
