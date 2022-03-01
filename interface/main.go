package main

import (
	"net/http"

	"go.uber.org/zap"
)

// =====================================
type F interface {
	f()
}

type S1 struct{}

func (s S1) f() {}

type S2 struct{}

func (s *S2) f() {}

// f1.f()无法修改底层数据
// f2.f() 可以修改底层数据,给接口变量f2赋值时使用的是对象指针
var f1 F = S1{}
var f2 F = &S2{}

type Handler struct {
	// ...
}

// =====================================
// 用于触发编译期的接口的合理性检查机制
// 如果Handler没有实现http.Handler,会在编译期报错
var _ http.Handler = (*Handler)(nil)

func (h *Handler) ServeHTTP(
	w http.ResponseWriter,
	r *http.Request,
) {
	// ...
}

// =====================================
type LogHandler struct {
	h   http.Handler
	log *zap.Logger
}

var _ http.Handler = LogHandler{}

func (h LogHandler) ServeHTTP(
	w http.ResponseWriter,
	r *http.Request,
) {
	// ...
}

// =====================================
type S struct {
	data string
}

func (s S) Read() string {
	return s.data
}

func (s *S) Write(str string) {
	s.data = str
}

func main() {
	sVals := map[int]S{1: {"A"}}
	// 你只能通过值调用 Read
	sVals[1].Read()

	// 这不能编译通过：
	//  sVals[1].Write("test")

	sPtrs := map[int]*S{1: {"A"}}

	// 通过指针既可以调用 Read，也可以调用 Write 方法
	sPtrs[1].Read()
	sPtrs[1].Write("test")
}
