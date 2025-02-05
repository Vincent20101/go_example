package main

import (
	"log"
	"sync"
)

func init() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)
}

type LinkStack struct {
	top    *node
	length int
	lock   *sync.RWMutex
}
type node struct {
	value interface{}
	prev  *node
}

// Create a new stack
func NewLinkStack() *LinkStack {
	return &LinkStack{nil, 0, &sync.RWMutex{}}
}

// Return the number of items in the stack
func (l *LinkStack) Len() int {
	return l.length
}

// View the top item on the stack
func (l *LinkStack) Peek() interface{} {
	if l.length == 0 {
		return nil
	}
	return l.top.value
}

// Pop the top item of the stack and return it
func (l *LinkStack) Pop() interface{} {
	l.lock.Lock()
	defer l.lock.Unlock()
	if l.length == 0 {
		return nil
	}
	n := l.top
	l.top = n.prev
	l.length--
	return n.value
}

// Push a value onto the top of the stack
func (l *LinkStack) Push(value interface{}) {
	l.lock.Lock()
	defer l.lock.Unlock()
	n := &node{value, l.top}
	l.top = n
	l.length++
}

// Print the contents of the stack
func (l *LinkStack) Print() {
	l.lock.RLock()
	defer l.lock.RUnlock()
	n := l.top
	for n != nil {
		log.Printf("%v", n.value)
		n = n.prev
	}
}
func main() {
	ls := NewLinkStack()
	ls.Push(1)
	ls.Push(2)
	ls.Push(3)
	ls.Print()
	log.Printf("%v", ls.Peek())
	log.Printf("%v", ls.Pop())
	ls.Print()
}
