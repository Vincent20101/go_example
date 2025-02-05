package main

import (
	"fmt"
	"net/http"
	"reflect"
	"sync"

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

	var tt []byte
	st := []byte{}
	fmt.Println(st)
	equal := reflect.DeepEqual(tt, st)
	//equal := bytes.Equal(tt, st)
	fmt.Println(equal)

	s1Val := S1{}
	s1Ptr := &S1{}
	s2Val := S2{}
	s2Ptr := &S2{}

	var i F
	i = s1Val
	i = s1Ptr
	i = s2Ptr

	//  下面代码无法通过编译。因为 s2Val 是一个值，而 S2 的 f 方法中没有使用值接收器
	//i = s2Val
	fmt.Println(i, s2Val)

	trips := []Trip{"a", "b", "c", "d", "e", "f"}
	var d1 Driver
	d1.SetTrips(trips)
	//d1.trips = append(d1.trips, "g")
	fmt.Println(d1.trips)
	// 你是要修改 d1.trips 吗？
	trips[0] = "aa"
	fmt.Println(trips)
	fmt.Println(d1.trips)

	stats := Stats{
		counters: make(map[string]int),
	}
	// snapshot 不再受互斥锁保护
	// 因此对 snapshot 的任何访问都将受到数据竞争的影响
	// 影响 stats.counters
	snapshot := stats.Snapshot()
	fmt.Println(snapshot)
	snapshot["all"] = 0
	fmt.Println(snapshot)
	fmt.Println(stats.Snapshot())
}

type Trip string
type Driver struct {
	trips []Trip
}

func (d *Driver) SetTrips(trips []Trip) {
	d.trips = trips

	// right way
	//d.trips = make([]Trip, len(trips))
	//copy(d.trips, trips)
}

type Stats struct {
	mu sync.Mutex

	counters map[string]int
}

// Snapshot 返回当前状态。
func (s *Stats) Snapshot() map[string]int {
	s.mu.Lock()
	defer s.mu.Unlock()
	//fmt.Println("Snapshot")
	//return s.counters

	// right way
	result := make(map[string]int, len(s.counters))
	for k, v := range s.counters {
		result[k] = v
	}
	return result
}
