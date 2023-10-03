package main

import (
	"bytes"
	"fmt"
	"mime/multipart"
	"net/http"
	"net/textproto"
)

func main() {
	// 创建一个缓冲区，用于保存multipart/form-data的内容
	var requestBody bytes.Buffer
	multipartWriter := multipart.NewWriter(&requestBody)

	// 设置一个自定义的内容类型
	mimeHeader := make(textproto.MIMEHeader)
	mimeHeader.Set("Content-Type", "application/vnd.3gpp.5gnas")

	// 创建一个部分，并设置部分的头部
	partW, err := multipartWriter.CreatePart(mimeHeader)
	if err != nil {
		fmt.Println("Error creating part:", err)
		return
	}

	// 数据，你可以替换为实际的数据
	n1Data := []byte("This is some sample data.")

	// 使用 partW.Writer 写入数据到部分
	//_, err = io.Copy(partW, bytes.NewReader(n1Data))
	_, err = partW.Write(n1Data)

	if err != nil {
		fmt.Println("Error writing data to part:", err)
		return
	}

	// 结束multipart写入
	multipartWriter.Close()

	fmt.Println(requestBody.String())
	//return
	// 创建一个POST请求，并将multipart/form-data请求主体设置为requestBody
	req, err := http.NewRequest("POST", "http://example.com/upload", &requestBody)
	if err != nil {
		fmt.Println("Error creating HTTP request:", err)
		return
	}

	// 设置请求头，指定Content-Type为multipart/form-data
	req.Header.Set("Content-Type", multipartWriter.FormDataContentType())

	// 发送HTTP请求
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending HTTP request:", err)
		return
	}
	defer resp.Body.Close()

	// 处理服务器的响应
	fmt.Println("HTTP Status:", resp.Status)
}
