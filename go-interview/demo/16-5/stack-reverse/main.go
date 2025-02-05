package main

import (
	"log"
	"stack-reverse/stack"
)

func init() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)
}

var i = 0

func recursive(seq int) {
	i += 1
	log.Println(seq)
	if i < 5 {
		recursive(i)
	}

	log.Println(seq)
}

/**
2,3,4,5  top=1  top := s.Pop()
3,4,5  top=2  top := s.Pop()
4,5  top=3  top := s.Pop()
5  top=4  top := s.Pop()
   top=5  top := s.Pop()

5            top=5       s.Push(top)
4,5          top=4       s.Push(top)
3,4,5        top=3       s.Push(top)
2,3,4,5      top=2       s.Push(top)
1,2,3,4,5    top=1       s.Push(top)

*/
func reverseStack(s *stack.ArrayStack) {
	if s.IsEmpty() {
		return
	}
	//把栈底元素移动到栈顶
	moveBottomToTop(s)
	//弹出栈顶元素
	top := s.Pop()

	//递归处理不包含栈顶元素的子栈
	reverseStack(s)
	s.Push(top)
}

/**
1,2，3，4，5
1,2，3，4  top=5     s.Pop()
1,2，3     top=4     s.Pop()
1,2        top=3     s.Pop()
1          top=2     s.Pop()
1          top=1     s.Push(top)



2,1	           top=2,top1=1             top1 := s.Pop() s.Push(top) s.Push(top1)
2,3,1          top=3,top1=1             top1 := s.Pop() s.Push(top) s.Push(top1)
2,3,4,1	       top=4,top1=1             top1 := s.Pop() s.Push(top) s.Push(top1)
2,3,4,5,1	   top=5,top1=1             top1 := s.Pop() s.Push(top) s.Push(top1)



*/
func moveBottomToTop(s *stack.ArrayStack) {
	if s.IsEmpty() {
		return
	}
	//弹出栈顶元素
	top := s.Pop()
	//递归处理不包含栈顶元素的子栈
	if !s.IsEmpty() {
		moveBottomToTop(s)
		top1 := s.Pop()
		s.Push(top)
		s.Push(top1)
	} else {
		s.Push(top)
	}
}

func main() {
	recursive(0)

	//stack := stack.NewArrayStack()
	//stack.Push(1).Push(2).Push(3).Push(4).Push(5)
	//stack.Print()
	//reverseStack(stack)
	//stack.Print()

}
