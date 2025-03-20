package main

import (
	"encoding/json"
	"fmt"
)

// 定义泛型接口
type Serializable[T any] interface {
	Serialize() ([]byte, error)
	Deserialize(data []byte) (T, error)
}

// 实现 Serializable 接口的结构体
type Person struct {
	Name string
	Age  int
}

func (p *Person) Serialize() ([]byte, error) {
	return json.Marshal(p)
}

func (p *Person) Deserialize(data []byte) (Person, error) {
	var person Person
	err := json.Unmarshal(data, &person)
	return person, err
}

// 解释：
//
// 1）定义了一个泛型接口 Serializable[T any]，包含 Serialize 和 Deserialize 方法。
//
// 2）Person 结构体实现了 Serializable 接口，实现了序列化和反序列化的功能。
//
// 3）在 main 函数中，展示了如何使用 Person 结构体进行序列化和反序列化。
func main() {
	person := &Person{Name: "Alice", Age: 30}
	data, err := person.Serialize()
	if err != nil {
		fmt.Println("Serialization error:", err)
		return
	}
	fmt.Println("Serialized data:", string(data))

	newPerson, err := person.Deserialize(data)
	if err != nil {
		fmt.Println("Deserialization error:", err)
		return
	}
	fmt.Println("Deserialized Person:", newPerson)
}
