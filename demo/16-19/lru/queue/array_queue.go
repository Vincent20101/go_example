package queue

import (
	"log"
	"sync"
)

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

// 数组队列，先进先出
type ArrayQueue struct {
	array []int      // 底层切片
	size  int        // 队列的元素数量
	lock  sync.Mutex // 为了并发安全使用的锁
}

func (queue *ArrayQueue) Size() int {
	queue.lock.Lock()
	defer queue.lock.Unlock()
	return queue.size
}

// 入队
// 直接将元素放在数组最后面即可，和栈一样，时间复杂度为：O(n)
func (queue *ArrayQueue) Add(v int) {
	queue.lock.Lock()
	defer queue.lock.Unlock()

	// 放入切片中，后进的元素放在数组最后面
	queue.array = append(queue.array, v)

	// 队中元素数量+1
	queue.size = queue.size + 1
}

// 从头部入队
func (queue *ArrayQueue) AddTop(v int) {
	queue.lock.Lock()
	defer queue.lock.Unlock()
	// 放入切片中，后进的元素放在数组最后面
	queue.array = append([]int{v}, queue.array...)
	// 队中元素数量+1
	queue.size = queue.size + 1
}

// 出队 从队尾弹出元素
func (queue *ArrayQueue) Peek() int {
	queue.lock.Lock()
	defer queue.lock.Unlock()

	// 队中元素已空
	if queue.size == 0 {
		panic("empty")
	}

	// 队列最尾部元素
	v := queue.array[queue.size-1]

	// 创建新的数组
	newArray := make([]int, queue.size-1, queue.size-1)
	newArray = queue.array[0 : queue.size-1]
	queue.array = newArray

	// 队中元素数量-1
	queue.size = queue.size - 1
	return v
}

// 移除指定元素
func (queue *ArrayQueue) Remove(item int) {
	queue.lock.Lock()
	defer queue.lock.Unlock()
	for k, v := range queue.array {
		if v == item {
			queue.array = append(queue.array[:k], queue.array[k+1:]...)
		}
	}
}
func (queue *ArrayQueue) GetTop() int {
	queue.lock.Lock()
	defer queue.lock.Unlock()
	return queue.array[0]
}
func (queue *ArrayQueue) Print() {
	log.Println(queue.array)
}
