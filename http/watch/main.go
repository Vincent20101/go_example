package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
)

func main() {
	// 定义一个HTTP长链接监听器
	listener := http.ListenAndServe(":8080", nil)

	// 定义一个HTTP长链接
	longLink := &http.Request{
		Method: "GET",
		URL:    "/long-link",
	}

	// 定义一个HTTP长链接处理函数
	handler := func(w http.ResponseWriter, r *http.Request) {
		// 获取HTTP长链接的URL
		url, err := r.URL.Parse()
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		// 获取HTTP长链接的长度
		length, err := strconv.Atoi(r.Header.Get("Content-Length"))
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		// 如果HTTP长链接的长度超过了指定的最大长度，则返回错误
		if length > maxLength {
			http.Error(w, "HTTP长链接的长度超过了指定的最大长度", http.StatusBadRequest)
			return
		}

		// 获取HTTP长链接的内容
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		// 将HTTP长链接的内容转换为字符串并返回
		fmt.Fprintln(w, string(body))