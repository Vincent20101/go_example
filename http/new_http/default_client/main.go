package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httputil"
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
	ctx, _ = context.WithTimeout(ctx, 10*time.Second)
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
	fmt.Println("reader len: ==== :", reader.Len())
	req, err := http.NewRequest(http.MethodPost, "http://localhost:8000", reader)

	req.Header.Set("Content-Type", "application/json")
	//req.Header.Add("User-Agent", "cli")
	req.Header.Set("Lhb", "application/json")
	req.Header.Set("3gpp-Sbi-Target-ApiRoot", "http://<HOSTPORT>")
	fmt.Println("lhb header: ", req.Header.Values("3gpp-Sbi-Target-ApiRoot"))

	dump, errs := httputil.DumpRequestOut(req, true)
	fmt.Printf("\n%s, DumpRequestOut error:%v\n", string(dump), errs)

	//getBody, errs := req.GetBody()
	//var bb []byte = make([]byte, req.ContentLength)
	//fmt.Println("getBody: ", getBody, errs)
	//getBody.Read(bb)
	//fmt.Println(string(bb))
	//
	//fmt.Println("ssss==req==", req)
	//fmt.Println("ssss====", req.Context())
	//if err != nil {
	//	log.Fatal("无法生成请求：", err)
	//}

	ctx = context.WithValue(ctx, "p", "q")
	//req = req.WithContext(ctx)
	req = req.WithContext(ctx)
	fmt.Println("第一：", req.Context().Value("p"))
	//deadline, ok := req.Context().Deadline()
	//fmt.Println(deadline, "===", ok)
	fmt.Println("ctx 第二：", ctx)
	fmt.Println("req的值：", req)
	fmt.Println("req的值：", req.Context())

	resp, err := http.DefaultClient.Do(req)
	//resp, err := http.Post("http://localhost:8000/", "application/json", reader)
	//resp, err := http.Get("http://localhost:8000/")
	//io.Copy(ioutil.Discard, resp.Body)
	defer func() {
		e := resp.Body.Close()
		fmt.Println("error: ", e)
	}()
	fmt.Println("resp====:", resp)
	fmt.Println("err=======:", err)
	if err != nil {
		log.Println("无法发送请求：", err)
		return
	}
	//dump, errs = httputil.DumpResponse(resp, true)
	//fmt.Printf("dump DumpResponse error:%v\n", dump)
	//fmt.Printf("\n%s, DumpResponse error:%v\n", string(dump), errs)
	// 获取请求报文的内容长度
	//length := resp.ContentLength
	// 新建一个字节切片，长度与请求报文的内容长度相同
	//body := make([]byte, length)
	// 读取 r 的请求主体，并将具体内容读入 body 中
	//resp.Body.Read(body)
	//fmt.Println("lhb==", body)

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
