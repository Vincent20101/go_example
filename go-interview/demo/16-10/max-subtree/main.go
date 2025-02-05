package main

import (
	"log"
	"math"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var maxSum = math.MinInt64

func sumMaxSubTree(root *TreeNode, maxRoot *TreeNode) int {
	if root == nil {
		return 0
	}
	//节点左子树的和
	lmax := sumMaxSubTree(root.Left, maxRoot)
	//节点右子树的和
	rmax := sumMaxSubTree(root.Right, maxRoot)
	sum := lmax + rmax + root.Val
	if sum > maxSum {
		maxSum = sum
		maxRoot.Val = root.Val
	}
	return sum

}

/*
*	      5
*	 /       \
*	3        -8
*	/  \     / \
-4  1    6   9
*/
func CreateTree() *TreeNode {
	tree := &TreeNode{}
	tree.Val = 5
	tree.Left = &TreeNode{Val: 3}
	tree.Right = &TreeNode{Val: -8}
	tree.Left.Left = &TreeNode{Val: -4}
	tree.Left.Right = &TreeNode{Val: 1}
	tree.Right.Left = &TreeNode{Val: 6}
	tree.Right.Right = &TreeNode{Val: 9}
	return tree
}

func main() {
	tree := CreateTree()
	maxRoot := &TreeNode{}
	log.Println(sumMaxSubTree(tree, maxRoot))

}
