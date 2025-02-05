package main

import "fmt"

// 链表定义
type LNode struct {
	Data interface{}
	Next *LNode
}

func NewLNode() *LNode {
	return &LNode{}
}

// 打印链表的方法
func (l *LNode) PrintNode() {
	for cur := l.Next; cur != nil; cur = cur.Next {
		fmt.Print(cur.Data, " ")
	}
	fmt.Println()
}

// 在链表尾部添加新节点
func (l *LNode) AppendNode(data interface{}) {
	cur := l
	for cur.Next != nil {
		cur = cur.Next
	}
	cur.Next = &LNode{Data: data}
}

func main() {
	list := NewLNode()
	list.AppendNode(1)
	list.AppendNode(2)
	list.AppendNode(3)
	list.PrintNode()
}
