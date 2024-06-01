package main

import (
	"btree-equal/btree"
	"log"
)

func IsEqual(root1, root2 *btree.TreeNode) bool {
	if root1 == nil && root2 == nil {
		return true
	}
	if root1 == nil && root2 != nil {
		return false
	}
	if root1 != nil && root2 == nil {
		return false
	}

	if root1.Val == root2.Val {
		return IsEqual(root1.Left, root2.Left) && IsEqual(root1.Right, root2.Right)
	}
	return false

}

func main() {
	root1 := &btree.TreeNode{Val: 5}
	root1.Insert(3)
	root1.Insert(8)
	root1.Insert(1)
	root1.Insert(4)
	root1.Insert(6)
	root1.Insert(9)

	root2 := &btree.TreeNode{Val: 5}
	root2.Insert(3)
	root2.Insert(8)
	root2.Insert(1)
	root2.Insert(4)
	root2.Insert(6)
	root2.Insert(9)
	log.Println(IsEqual(root1, root2))
}
