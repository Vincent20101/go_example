package stack

import (
	"log"
	"sync"
)

type Stack interface {
	//往栈中插入一个元素
	Push(interface{}) *ArrayStack

	//取出栈顶元素，如果栈为空，err!=nil
	Pop() interface{}

	//获取栈顶元素，如果栈为空，err!=nil
	Peek() interface{}

	//获取当前栈的大小
	GetSize() int
}
type ArrayStack struct {
	data []interface{}
	size int        // 栈的元素数量
	lock sync.Mutex // 为了并发安全使用的锁
}

func NewArrayStack() *ArrayStack {
	return &ArrayStack{
		data: make([]interface{}, 0, 16),
		size: 0,
		lock: sync.Mutex{},
	}
}

// 入栈操作
func (a *ArrayStack) Push(v interface{}) *ArrayStack {
	a.lock.Lock()
	defer a.lock.Unlock()

	// 放入切片中，后进的元素放在数组最后面
	a.data = append(a.data, v)

	// 栈中元素数量+1
	a.size = a.size + 1
	return a
}

// 出栈操作
func (a *ArrayStack) Pop() interface{} {
	//加锁实现并发安全
	a.lock.Lock()
	defer a.lock.Unlock()

	// 栈中元素已空
	if a.size == 0 {
		return nil
	}

	// 栈顶元素
	v := a.data[a.size-1]

	// 切片收缩，但可能占用空间越来越大
	//如果切片偏移量向前移动 stack.array[0 : stack.size-1]，表明最后的元素已经不属于该数组了，数组变相的缩容了。
	//此时，切片被缩容的部分并不会被回收，仍然占用着空间，所以空间复杂度较高，但操作的时间复杂度为：O(1)
	//a.data = a.data[0 : a.size-1]

	// 创建新的数组，空间占用不会越来越大，但可能移动元素次数过多
	//如果创建新的数组 newArray，然后把老数组的元素复制到新数组，就不会占用多余的空间，
	//但移动次数过多，时间复杂度为：O(n)
	newArray := make([]interface{}, a.size-1, a.size-1)
	for i := 0; i < a.size-1; i++ {
		newArray[i] = a.data[i]
	}
	a.data = newArray

	// 栈中元素数量-1
	a.size = a.size - 1
	return v
}

// 取栈顶元素（数组的最后一个元素）
func (a *ArrayStack) Peek() interface{} {
	// 栈中元素已空
	if a.size == 0 {
		return nil
	}

	// 栈顶元素值
	v := a.data[a.size-1]
	return v
}

func (a *ArrayStack) GetSize() int {
	return a.size
}

// 栈是否为空
func (a *ArrayStack) IsEmpty() bool {
	return a.size == 0
}

func (a *ArrayStack) Print() {
	if a.IsEmpty() {
		log.Println("Empty stack!")
		return
	}
	log.Printf("Stack size: %d, data: %v\n", a.size, a.data)
}
