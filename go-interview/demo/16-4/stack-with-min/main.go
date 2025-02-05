package main

import (
	"log"
	"math"
	"stack-with-min/stack"
)

func init() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)
}

type StackWithMin struct {
	//存储栈中元素
	elemStack *stack.ArrayStack

	//存储栈中最小元素
	minStack *stack.ArrayStack
}

// 入栈操作
func (s *StackWithMin) Push(v int) {
	s.elemStack.Push(v)
	//这里以int类型为例
	if s.minStack.IsEmpty() || v < s.minStack.Peek().(int) {
		s.minStack.Push(v)
	}
}

func (s *StackWithMin) Min() int {
	if s.minStack.IsEmpty() {
		return math.MaxInt32
	} else {
		return s.minStack.Peek().(int)
	}
}

// 出栈操作
func (s *StackWithMin) Pop() int {
	top := s.elemStack.Pop().(int)
	if top == s.Min() {
		s.minStack.Pop()
	}
	return top
}

func main() {
	stack := &StackWithMin{
		elemStack: &stack.ArrayStack{},
		minStack:  &stack.ArrayStack{},
	}
	stack.Push(2)
	log.Println("最小值为", stack.Min())
	stack.Push(3)
	log.Println("最小值为", stack.Min())
	stack.Push(1)
	log.Println("最小值为", stack.Min())
	stack.Pop()
	log.Println("最小值为", stack.Min())
}
