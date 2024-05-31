package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func (node *TreeNode) Insert(val int) {
	// 递归找到插入位置并创建新节点
	if val < node.Val {
		if node.Left == nil {
			node.Left = &TreeNode{Val: val}
		} else {
			node.Left.Insert(val)
		}
	} else {
		if node.Right == nil {
			node.Right = &TreeNode{Val: val}
		} else {
			node.Right.Insert(val)
		}
	}
}

func (node *TreeNode) Print() {
	queue := []*TreeNode{node} // 创建队列，初始包含根节点
	for len(queue) > 0 {
		current := queue[0]         // 获取队列的第一个元素
		queue = queue[1:]           // 队列去掉第一个元素
		fmt.Print(current.Val, " ") // 访问当前节点

		if current.Left != nil {
			queue = append(queue, current.Left) // 左子节点入队
		}
		if current.Right != nil {
			queue = append(queue, current.Right) // 右子节点入队
		}
	}
}

func main() {
	// 创建根节点
	root := &TreeNode{Val: 5}
	// 插入元素
	root.Insert(3)
	root.Insert(8)
	root.Insert(1)
	root.Insert(4)
	root.Insert(6)
	root.Insert(9)
	root.Print()
}
