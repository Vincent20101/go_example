package main

import (
	"errors"
	"log"
	"os"
	"os/signal"
	"time"
)

// timeout 规定了必须在多少秒内处理完成
const timeout = 3 * time.Second

func main() {
	log.Println("Starting work.")

	// 为本次执行分配超时时间
	r := New(timeout)

	// 加入要执行的任务
	r.Add(createTask(), createTask(), createTask())
	// 执行任务并处理结果
	if err := r.Start(); err != nil {
		switch err {
		case ErrTimeOut:
			log.Println("Terminating due to timeout.")
			os.Exit(1)
		case ErrInterrupt:
			log.Println("Terminating due to interrupt.")
			os.Exit(2)
		}
	}
	//return
	log.Println("Process ended.")
}

//createTask 返回一个根据 id 休眠指定秒数的示例任务
func createTask() func(int) {
	return func(id int) {
		log.Printf("Processor - Task #%d.", id)
		time.Sleep(time.Duration(5) * time.Second)
	}
}

//统一错误处理
//超时错误信息，会在任务执行超时时返回
var ErrTimeOut = errors.New("received timeout")

//中断错误信号，会在接收到操作系统的事件时返回
var ErrInterrupt = errors.New("received interrupt")

// Runner 在给定的超时时间内执行一组任务
// 并且在操作系统发送中断信号时结束这些任务
type Runner struct {
	//从操作系统发送信号
	interrupt chan os.Signal
	//报告处理任务已完成
	complete chan error
	//报告处理任务已经超时
	timeout <-chan time.Time
	//持有一组以索引为顺序的依次执行的以int类型id为参数的函数
	tasks []func(id int)
}

// New 函数返回一个新的准备使用的Runner,d:自定义分配的时间
// 初始化结构体参数
func New(d time.Duration) *Runner {
	return &Runner{
		interrupt: make(chan os.Signal, 1),
		complete:  make(chan error),
		//会在另一线程经过时间段d后向返回值发送当时的时间。
		// 因为 task 字段的零值是 nil，已经满足初始化的要求，所以没有被明确初始化。
		timeout: time.After(d),
	}
}

//将一个任务加入到Runner中
func (r *Runner) Add(tasks ...func(id int)) {
	r.tasks = append(r.tasks, tasks...)
}

//开始执行所有任务，并监控通道事件
func (r *Runner) Start() error {
	//监控所有的中断信号
	signal.Notify(r.interrupt, os.Interrupt)
	//使用不同的goroutine执行不同的任务
	go func() {
		r.complete <- r.run()
	}()
	//使用 select 语句来监控goroutine的通信
	select {
	//等待任务完成
	case err := <-r.complete:
		return err
	//任务超时
	case <-r.timeout:
		return ErrTimeOut
	}
}

//执行每一个已注册的任务
func (r *Runner) run() error {
	for id, task := range r.tasks {
		//检测操作系统的中断信号
		if r.gotInterrupt() {
			return ErrInterrupt
		}
		//执行已注册的任务
		task(id)
	}
	return nil
}

//检测是否收到了中断信号
func (r *Runner) gotInterrupt() bool {
	select {
	//当中断事件被触发时
	case <-r.interrupt:
		//停止接收后续的任何信号
		signal.Stop(r.interrupt)
		return true
		//继续执行
	default:
		return false
	}
}
