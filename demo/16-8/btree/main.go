package main

// 二叉树定义
type BNode struct {
	Data       interface{}
	LeftChild  *BNode
	RightChild *BNode
}

func NewBTree() *BNode {
	return &BNode{}
}

// 通过数组创建中序的二叉树
func ArrayToTree(arr []int, start int, end int) *BNode {
	var root *BNode
	if end >= start {
		root = &BNode{}
		mid := (start + end + 1) / 2
		//树的根结点为数组中间的元素
		root.Data = arr[mid]
		//递归的用左半部分数组构造root的左子树
		root.LeftChild = ArrayToTree(arr, start, mid-1)
		//递归的用右半部分数组构造root的右子树
		root.RightChild = ArrayToTree(arr, mid+1, end)
	}
	return root
}

func main() {

}
