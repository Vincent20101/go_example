package main

import (
	"fmt"
	"sync"
	"time"
)

type Task struct {
	ID         int
	StartTime  time.Time
	RetryCount int
}

type TaskQueue struct {
	mu           sync.Mutex
	tasks        []*Task
	notify       chan struct{}
	closed       bool
	closeCh      chan struct{}
	addCh        chan *Task
	popCh        chan func(*Task) error
	retryBackoff time.Duration
}

func NewTaskQueue(retryBackoff time.Duration) *TaskQueue {
	queue := &TaskQueue{
		notify:       make(chan struct{}, 1),
		closeCh:      make(chan struct{}),
		addCh:        make(chan *Task, 1),
		popCh:        make(chan func(*Task) error),
		retryBackoff: retryBackoff,
	}
	go queue.run()
	return queue
}

func (q *TaskQueue) run() {
	for {
		select {
		case task := <-q.addCh:
			q.mu.Lock()
			q.tasks = append(q.tasks, task)
			q.notify <- struct{}{}
			q.mu.Unlock()

		case popFunc := <-q.popCh:
			go func() {
				q.mu.Lock()
				if len(q.tasks) == 0 {
					q.mu.Unlock()
					<-q.notify
					return
				}

				task := q.tasks[0]
				now := time.Now()
				if task.StartTime.After(now) {
					time.AfterFunc(task.StartTime.Sub(now), func() {
						q.popCh <- popFunc
					})
					q.mu.Unlock()
					return
				}

				q.tasks = q.tasks[1:]
				q.mu.Unlock()

				err := popFunc(task)
				if err != nil {
					q.mu.Lock()
					task.RetryCount++
					retryDuration := time.Duration(task.RetryCount) * q.retryBackoff
					time.Sleep(retryDuration)
					q.mu.Unlock()

					q.Add(task)
				}
			}()

		case <-q.closeCh:
			return
		}
	}
}

func (q *TaskQueue) Add(task *Task) {
	q.addCh <- task
}

func (q *TaskQueue) Pop(popFunc func(*Task) error) error {
	if q.closed {
		return fmt.Errorf("queue is closed")
	}

	promise := make(chan error)
	q.popCh <- func(task *Task) error {
		err := popFunc(task)
		promise <- err
		return err
	}

	return <-promise
}

func (q *TaskQueue) Close() {
	q.mu.Lock()
	defer q.mu.Unlock()

	if q.closed {
		return
	}

	close(q.closeCh)
	q.closed = true
}

func main() {
	retryBackoff := time.Second

	queue := NewTaskQueue(retryBackoff)

	// 模拟添加任务
	queue.Add(&Task{ID: 1, StartTime: time.Now()})
	queue.Add(&Task{ID: 2, StartTime: time.Now().Add(3 * time.Second)})
	queue.Add(&Task{ID: 3, StartTime: time.Now().Add(2 * time.Second)})
	queue.Add(&Task{ID: 4, StartTime: time.Now().Add(5 * time.Second)})

	// 模拟处理任务
	popFunc := func(task *Task) error {
		fmt.Printf("Processing task ID %d\n", task.ID)
		if task.ID == 2 {
			return fmt.Errorf("task processing failed")
		}
		return nil
	}

	wg := sync.WaitGroup{}
	for i := 0; i < 4; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			err := queue.Pop(popFunc)
			if err != nil {
				fmt.Println(err)
			}
		}()
	}

	wg.Wait()

	// 关闭队列
	queue.Close()
}
