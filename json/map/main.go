package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	// 定义一个 []interface{} 切片
	interfaces := []interface{}{
		map[string]interface{}{"name": "Alice", "age": 30},
		map[string]interface{}{"name": "Bob", "age": 35},
	}

	// 序列化为 JSON 字节串
	jsonBytes, err := json.Marshal(interfaces)
	if err != nil {
		fmt.Println("Error marshaling JSON:", err)
		return
	}

	// 打印 JSON 字节串
	fmt.Println("Serialized JSON:", string(jsonBytes))

	// 反序列化到结构体切片
	var people []Person
	err = json.Unmarshal(jsonBytes, &people)
	if err != nil {
		fmt.Println("Error unmarshaling JSON:", err)
		return
	}

	// 打印反序列化后的结果
	fmt.Println("Deserialized:")
	for _, p := range people {
		fmt.Printf("Name: %s, Age: %d\n", p.Name, p.Age)
	}
}
