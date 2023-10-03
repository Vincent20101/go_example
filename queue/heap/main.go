package main

import (
	"container/heap"
	"fmt"
	"sync"
	"time"
)

// Message 表示队列中的消息
type Message struct {
	ID        int
	Content   string
	DelayTime time.Time
}

// MessageQueue 是消息队列的类型
type MessageQueue struct {
	mu       sync.Mutex
	messages map[int]*Message
	pq       MessagePriorityQueue
}

func NewMessageQueue() *MessageQueue {
	mq := &MessageQueue{
		messages: make(map[int]*Message),
	}
	heap.Init(&mq.pq)
	return mq
}

func (mq *MessageQueue) AddMessage(content string, delay time.Duration) {
	mq.mu.Lock()
	defer mq.mu.Unlock()

	message := &Message{
		ID:        len(mq.messages) + 1,
		Content:   content,
		DelayTime: time.Now().Add(delay),
	}
	mq.messages[message.ID] = message
	heap.Push(&mq.pq, message)
}

func (mq *MessageQueue) RemoveMessageByID(id int) {
	mq.mu.Lock()
	defer mq.mu.Unlock()

	if message, ok := mq.messages[id]; ok {
		delete(mq.messages, id)
		heap.Remove(&mq.pq, message.ID)
	}
}

func (mq *MessageQueue) ConsumeMessages() {
	now := time.Now()

	mq.mu.Lock()
	defer mq.mu.Unlock()

	for len(mq.pq) > 0 && mq.pq[0].DelayTime.Before(now) {
		message := heap.Pop(&mq.pq).(*Message)
		delete(mq.messages, message.ID)
		fmt.Printf("消费消息 %d: %s\n", message.ID, message.Content)
	}
}

// MessagePriorityQueue 优先队列类型
type MessagePriorityQueue []*Message

func (pq MessagePriorityQueue) Len() int {
	return len(pq)
}

func (pq MessagePriorityQueue) Less(i, j int) bool {
	return pq[i].DelayTime.Before(pq[j].DelayTime)
}

func (pq MessagePriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].ID = i
	pq[j].ID = j
}

func (pq *MessagePriorityQueue) Push(x interface{}) {
	n := len(*pq)
	message := x.(*Message)
	message.ID = n
	*pq = append(*pq, message)
}

func (pq *MessagePriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	message := old[n-1]
	message.ID = -1 // 从堆中移除的元素的索引设置为-1
	*pq = old[0 : n-1]
	return message
}

func main() {
	messageQueue := NewMessageQueue()

	// 添加一些消息到队列中
	messageQueue.AddMessage("消息1", 2*time.Second)
	messageQueue.AddMessage("消息2", 5*time.Second)
	messageQueue.AddMessage("消息3", 1*time.Second)

	// 删除队列中的消息
	messageQueue.RemoveMessageByID(2)

	// 模拟消息的消费
	time.Sleep(3 * time.Second)
	messageQueue.ConsumeMessages()
}
