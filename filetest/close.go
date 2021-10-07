package main

var c = make(chan int, 10)
var a string

func f() {
	a = "hello world"
	close(c)
}
func main() {
	go f()
	b, ok := <-c
	println(b)
	println(ok)
	println(a)
}
