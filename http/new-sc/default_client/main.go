package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.TODO(), 10*time.Second)
	defer cancel()
	//go httpDirectGet()
	go httpGetWithContext(ctx)
	<-ctx.Done()
	//time.Sleep(1000e9)
}

func httpDirectGet() {
	resp, err := http.Get("http://localhost:8083/moreJSON")
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	fmt.Println(string(data))
}

type Student struct {
	Id   string
	Name string
}

func httpGetWithContext(ctx context.Context) {
	ctx, _ = context.WithTimeout(ctx, 3*time.Second)
	//ctx = context.Background()
	//defer cancel()
	stu := Student{
		Id:   "2ed4tg5fe35fgty3yy6uh",
		Name: "amber",
	}
	stus, err := json.Marshal(stu)
	fmt.Println(err)
	reader := bytes.NewReader(stus)
	fmt.Println(string(stus))
	fmt.Println(reader)
	//reader := strings.NewReader(stus)
	req, err := http.NewRequest(http.MethodPost, "http://localhost:8000", reader)
	//dump, errs := httputil.DumpRequestOut(req, true)
	//fmt.Printf("\n%s, DumpRequestOut error:%v\n", string(dump), errs)

	if err != nil {
		log.Fatal("无法生成请求：", err)
	}
	//req.Header.Add("User-Agent", "cli")
	req.Header.Set("Content-Type", "application/json")
	ctx = context.WithValue(ctx, "p", "q")

	// 只是配置了客户端的超时,并不能传递到服务端
	req = req.WithContext(ctx)

	fmt.Println("第一：", req.Context().Value("p"))
	fmt.Println("ctx 第二：", ctx)
	fmt.Println("req的值：", req)
	fmt.Println("req的值：", req.Context())

	client := http.Client{}
	resp, err := client.Do(req)
	//resp, err := http.Post("http://localhost:8000/", "application/json", reader)
	//resp, err := http.Get("http://localhost:8000/")
	//io.Copy(ioutil.Discard, resp.Body)
	fmt.Println("resp====:", resp)
	fmt.Println("err=======:", err)
	defer func() {
		if resp != nil {
			e := resp.Body.Close()
			fmt.Println("error: ", e)
		}
	}()

	if err != nil {
		log.Println("无法发送请求：", err)
		return
	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("无法读取返回内容：", err)
	}
	fmt.Println("data 结果：", data)
	data1, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("data1 结果：", data1)
	type Student struct {
		Name string
	}
	var s Student
	json.Unmarshal(data, &s)
	fmt.Println("结果：", string(data))
	fmt.Printf("结果：[%+v]\n", s)
}
