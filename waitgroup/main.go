package main

import (
	"fmt"
	"sync"
)

//总结：如果一个协程A发生了panic  协程B会因为协程A的panic而挂掉
//如果协程A发生了panic 协程B不能用recover捕获到协程A当中的panic
//所以 多协程并发的时候绝对不能因为一个子协程的panic而搞死所有的子协程！！！
//所以 协程当中只有协程自己内部的recover才能捕获到自己抛出的panic
var wg sync.WaitGroup

func main() {
	//1.声明channel
	var ch chan int
	//1.1make实例化一下
	ch = make(chan int)
	//2.起来两个协程
	for i := 1; i <= 10; i++ {
		//2.1使用协程等待 所以wg.add(1)

		wg.Add(1)
		go func(index int) {
			//2.2判断协程内部是否出现错误 如果出现错误需要defer当中用recover()捕获错误进行处理
			defer func() {
				//2.4 协程内部的panic 必须在本协程内部自己通过recover来进行捕获 不然一旦出现panic会搞死所有的子协程
				err := recover()
				if err != nil {
					fmt.Println(err)
				}
				//2.3不要忘记wg.Done()告诉协程完事了
				wg.Done()
			}()
			//3.这里我们故意的搞一个错误panic  只要遇到panic程序立马中断 死掉之前会执行defer 所以我们在defer当中通过revocer()来捕获异常进行处理
			//3.1 协程内部的panic 必须在本协程内部自己通过recover来进行捕获 不然一旦出现panic会搞死所有的子协程
			//if index == 2 {
			//	panic("故意的错误")
			//}
			//4.写入值到无缓冲的channel当中去
			ch <- index
			//4.1 这里是在写入无缓冲channel之后的操作 如果下边遍历ch 那么必须要使用协程等待 不然这个地方可能就会执行不到
			fmt.Println(fmt.Sprintf("我是第：%d个", index))
		}(i)
	}

	//5.我们起来一个单独的协程 用来监控所有的子协程的完成情况 一旦完成 我们关闭close(ch)通道
	go func() {
		//5.1 为啥要关闭呢？因为下边for range了通道 会一直监控着ch通道  所以之类监控子协程一旦全部完成 立马close掉通道 那么底下的for range ch 就不会发生死锁
		defer close(ch)
		wg.Wait()
	}()

	//6.for range ch 一直遍历ch 通道  知道ch通道被close掉
	for v := range ch {
		fmt.Println(v)
	}

}
