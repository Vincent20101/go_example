package main

import (
	"fmt"
	"math"
	"math/cmplx"
	"reflect"
	"runtime"
	"strconv"
)

func main() {
	euler()
	triangle()
	triangle2()
	consts()
	fmt.Println(convertToBin(5))
	fmt.Println(convertToBin(13))
	fmt.Println(convertToBin(51))
	apply(pow, 3, 4)
	fmt.Println(apply(func(a int, b int) int {
		return int(math.Pow(float64(a), float64(b)))
	}, 3, 4))
}

// 欧拉公式
func euler() {
	fmt.Println(
		cmplx.Exp(1i*math.Pi)+1,
		cmplx.Pow(math.E, 1i*math.Pi)+1)
	fmt.Printf("%.3f\n%.3f\n",
		cmplx.Exp(1i*math.Pi)+1,
		cmplx.Pow(math.E, 1i*math.Pi)+1)
}

func triangle() {
	var a, b int = 3, 4
	var c int
	c = int(math.Sqrt(float64(a*a + b*b)))
	fmt.Println(c)
}

func triangle2() {
	// const 数值可作为各种类型使用
	const a, b = 3, 4
	var c int
	c = int(math.Sqrt(a*a + b*b))
	fmt.Println(c)
}

func consts() {
	const (
		b = 1 << (10 * iota)
		kb
		mb
		gb
		tp
		pb
	)
	fmt.Println(b, kb, mb, gb, tp, pb)
}

func convertToBin(n int) string {
	result := ""
	for ; n > 0; n /= 2 {
		lsb := n % 2
		result = strconv.Itoa(lsb) + result
	}
	return result
}

func pow(a, b int) int {
	return int(math.Pow(float64(a), float64(b)))
}

func apply(op func(int, int) int, a, b int) int {
	fmt.Printf("Calling %s with %d, %d\n",
		runtime.FuncForPC(reflect.ValueOf(op).Pointer()).Name(), a, b)
	return op(a, b)
}

func apply2(op func(int, int) int, a, b int) int {
	p := reflect.ValueOf(op).Pointer()
	opName := runtime.FuncForPC(p).Name()
	fmt.Printf("Calling %s with %d, %d\n",
		opName, a, b)
	return op(a, b)
}
