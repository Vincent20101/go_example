package main

import (
	"fmt"
	"math"
	"net"
	"strconv"
	"strings"
	"sync/atomic"
	"time"
)

func main() {

	var ip net.IP = []byte("nil")
	fmt.Println(ip.To4().String() == "<nil>")
	if strings.Contains("nil", "nil") {
		fmt.Println("lhb")
	}
	type st struct {
		name string
		age  int
		b    atomic.Bool
	}
	var stSlice = make([]map[int]*st, 0)
	for i := 0; i < 10; i++ {
		m := make(map[int]*st)
		m[i] = &st{
			name: strconv.Itoa(i),
			age:  i,
		}
		m[i].b.Store(true)
		fmt.Println("lhb:", m[i].b.Load())
		stSlice = append(stSlice, m)
	}
	fmt.Println(stSlice)

	for k, v := range stSlice {
		if v != nil && k == 0 {
			stSlice[k] = stSlice[len(stSlice)-1]
			stSlice = stSlice[:len(stSlice)-1]
		}
	}

	fmt.Println(stSlice)

	fmt.Println(strconv.FormatBool(true))
	fmt.Println(strconv.FormatBool(false))
	fmt.Println(strings.Compare("false", "true"))

	fmt.Println(time.Time{}.Unix())
	smap := make(map[string]int)
	if b, ok := smap["sdffff"]; !ok {
		fmt.Println(b, ok)
	}
	fmt.Println(smap["sdf"])
	smap["a"] = 5
	smap = make(map[string]int)
	fmt.Println(smap)
	var slice1 []string
	for i, s := range slice1 {
		fmt.Println("range slice1", i, s)
	}
	fmt.Println(len(slice1))

	var sss = "0"
	atoi, err2 := strconv.Atoi(sss)
	fmt.Println(atoi, err2)
	fmt.Println("max:", math.MaxInt32)
	fmt.Println("Nanosecond:", time.Now().Nanosecond())
	fmt.Println("Nanosecond:", time.Now().UnixNano())
	fmt.Println(LinearBuckets(0.0001, 2, 6))
	fmt.Println(ExponentialBuckets(0.0001, 10, 6))
	var i uint64
	fmt.Println(atomic.AddUint64(&i, 1))
	fmt.Println(atomic.AddUint64(&i, 1))

	s := []string{"a", "b", "c"}
	fmt.Println("lhb===", s[:3])
	fmt.Println("lhb===", len(s))
	for l := range s {
		fmt.Println(l)
	}
	var a []string
	fmt.Println(a)
	a1 := append(a, "test")
	fmt.Println(a1, a1[0])

	a2, b, c, d, e := 1, 2, 3, 4, 5
	s1 := []*int{&a2, &b, &c, &d, &e}
	s1 = s1[:0]
	fmt.Println(s1)

	fmt.Println(0 & (0x800 | 0x400))

	var t time.Time
	fmt.Println(&t)

	bs := []byte{0, 0, 0, 7}
	fmt.Println(string(bs))

	fmt.Println(32 << (^uint(0) >> 63))

}

// LinearBuckets creates 'count' buckets, each 'width' wide, where the lowest
// bucket has an upper bound of 'start'. The final +Inf bucket is not counted
// and not included in the returned slice. The returned slice is meant to be
// used for the Buckets field of HistogramOpts.
//
// The function panics if 'count' is zero or negative.
func LinearBuckets(start, width float64, count int) []float64 {
	if count < 1 {
		panic("LinearBuckets needs a positive count")
	}
	buckets := make([]float64, count)
	for i := range buckets {
		buckets[i] = start
		start += width
	}
	return buckets
}

// ExponentialBuckets creates 'count' buckets, where the lowest bucket has an
// upper bound of 'start' and each following bucket's upper bound is 'factor'
// times the previous bucket's upper bound. The final +Inf bucket is not counted
// and not included in the returned slice. The returned slice is meant to be
// used for the Buckets field of HistogramOpts.
//
// The function panics if 'count' is 0 or negative, if 'start' is 0 or negative,
// or if 'factor' is less than or equal 1.
func ExponentialBuckets(start, factor float64, count int) []float64 {
	if count < 1 {
		panic("ExponentialBuckets needs a positive count")
	}
	if start <= 0 {
		panic("ExponentialBuckets needs a positive start value")
	}
	if factor <= 1 {
		panic("ExponentialBuckets needs a factor greater than 1")
	}
	buckets := make([]float64, count)
	for i := range buckets {
		buckets[i] = start
		start *= factor
	}
	return buckets
}
