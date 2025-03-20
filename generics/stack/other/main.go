package main

import (
	"fmt"
)

// 定义一个泛型栈
type Stack[T any] struct {
	elements []T
}

// 压栈
func (s *Stack[T]) Push(element T) {
	s.elements = append(s.elements, element)
}

// 弹栈
func (s *Stack[T]) Pop() (T, bool) {
	if len(s.elements) == 0 {
		var zero T
		return zero, false
	}
	index := len(s.elements) - 1
	element := s.elements[index]
	s.elements = s.elements[:index]
	return element, true
}

// 查看栈顶元素
func (s *Stack[T]) Peek() (T, bool) {
	if len(s.elements) == 0 {
		var zero T
		return zero, false
	}
	return s.elements[len(s.elements)-1], true
}

// 解释：
//
// 1）定义了一个泛型栈 Stack[T any]，其中 T 可以是任意类型。
//
// 2）提供了 Push、Pop 和 Peek 方法，分别用于压栈、弹栈和查看栈顶元素。
//
// 3）在 main 函数中，创建了 int 类型和 string 类型的栈实例，展示了泛型数据结构的灵活性。
func main() {
	intStack := Stack[int]{}
	intStack.Push(10)
	intStack.Push(20)
	fmt.Println("Pop from intStack:", intStack.Pop())

	stringStack := Stack[string]{}
	stringStack.Push("hello")
	stringStack.Push("world")
	fmt.Println("Peek from stringStack:", stringStack.Peek())
}
