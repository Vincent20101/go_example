package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httputil"
	"os"
	"time"
)

type Student struct {
	Name string
}

func main() {
	// 创建一个监听8000端口的服务器
	http.ListenAndServe(":8000", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//fmt.Println(spew.Sdump(r))
		fmt.Println(r)
		ctx := r.Context()
		fmt.Println("=====ctx:", ctx)
		deadline, ok := r.Context().Deadline()
		fmt.Println(deadline, "===", ok)
		// 输出到STDOUT展示处理已经开始
		fmt.Fprint(os.Stdout, "processing request\n")

		dump, errs := httputil.DumpRequest(r, true)
		fmt.Printf("\n%s, DumpRequest error:%v\n", string(dump), errs)

		// 当body没有被读取时，客户端 ctx cancel，服务端 ctx 也不会 cancel
		// 获取请求报文的内容长度
		length := r.ContentLength
		fmt.Println("length: ", length)
		// 新建一个字节切片，长度与请求报文的内容长度相同
		body := make([]byte, length)
		// 读取 r 的请求主体，并将具体内容读入 body 中
		r.Body.Read(body)
		r.Body = io.NopCloser(bytes.NewBuffer(body))
		// 将字节切片内容写入相应报文
		fmt.Println("body is: ", string(body), "\n")
		fmt.Println("body is: ", r.Body, "\n")
		defer r.Body.Close()
		//fmt.Println("request: ", r.URL.Path)
		// 通过select监听多个channel
		s := Student{
			Name: "lhb",
		}
		sjosn, _ := json.Marshal(s)
		select {
		case <-time.After(1 * time.Second):
			// 如果两秒后接受到了一个消息后，意味请求已经处理完成
			// 我们写入"request processed"作为响应
			fmt.Println("time.After=====")
			time.Sleep(15 * time.Second)
			fmt.Println("time.after.15")
			//w.Write([]byte("request processed"))
			w.Write(sjosn)
		case <-ctx.Done():

			// 如果处理完成前取消了，在STDERR中记录请求被取消的消息
			fmt.Fprint(os.Stderr, "request cancelled\n")
			//default:
			//time.Sleep(10 * time.Second)
			//fmt.Println("test11")
			//w.Write([]byte("request processed"))
			//w.Write(sjosn)
		}

	}))
}

// https://studygolang.com/articles/6301
//golang里长短连接的一些处理
//1)golang强制短连接
//出了在请求的头里加上connection:close, 也可以设置request结构体Close成员变量为true，比如：
//
//req, _ := http.NewRequest("Get", "http://example.com", nil)
//req.Close = true
//2）golang对长连接的要求
//golang client不设置主动断连，还得注意下，想要保持长连接，得保证以下两个点
//
//1. defer resp.Body.Close()  //别忘了close body，不然长连接保持不了
//2. body, err := ioutil.ReadAll(resp.Body)//记得读完resp.body或者放置一个结束符号 io.Copy(ioutil.Discard,resp.Body)
//3) 解决产生大量close_wait
//解决方案网上挺多，待总结个比较合适的方案，再补上
