package main

import (
	"encoding/json"
	"fmt"
	"net/http"
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

		//dump, errs := httputil.DumpRequest(r, true)
		//fmt.Printf("\n%s, DumpRequest error:%v\n", string(dump), errs)

		// 当body没有被读取时，客户端 ctx cancel，服务端 ctx 也不会 cancel
		// 获取请求报文的内容长度
		length := r.ContentLength
		fmt.Println("length: ", length)
		// 新建一个字节切片，长度与请求报文的内容长度相同
		body := make([]byte, length)
		// 读取 r 的请求主体，并将具体内容读入 body 中
		r.Body.Read(body)
		// 将字节切片内容写入相应报文
		fmt.Println("body is: ", string(body), "\n")
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
