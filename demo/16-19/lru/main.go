package main

import (
	"errors"
	"lru/queue"
	"sync"
)

type LRUCache struct {
	cacheSize int
	queue     *queue.ArrayQueue
	cache     *sync.Map
}

func (c *LRUCache) Set(key int, value int) {
	if c.IsFull() {
		//缓存队列已满，删除队尾元素
		c.cache.Delete(c.queue.Peek())
	}
	//新key添加到队列头部
	c.queue.AddTop(key)
	//将缓存节点添加到hash表
	c.cache.Store(key, value)
}

func (c *LRUCache) IsFull() bool {
	return c.queue.Size() == c.cacheSize
}

//获取缓存节点
//如果key在缓存中，则将其移动到队头，并返回其value

func (c *LRUCache) Get(key int) (int, error) {

	value, ok := c.cache.Load(key)

	if ok {
		//key在缓存中，且不在队头，将其移动到队头
		if c.queue.GetTop() != key {
			c.queue.Remove(key)
			c.queue.AddTop(key)
		}
		return value.(int), nil
	}
	//key不在缓存中,通知调用者找不到
	return -1, errors.New("key not found")
}

func main() {
	lru := LRUCache{
		cacheSize: 3,
		queue:     &queue.ArrayQueue{},
		cache:     &sync.Map{},
	}
	lru.Set(1, 1)
	lru.Set(2, 2)
	lru.Set(3, 3)
	lru.Set(4, 4)
	lru.queue.Print()
	lru.Get(2)
	lru.queue.Print()

}
