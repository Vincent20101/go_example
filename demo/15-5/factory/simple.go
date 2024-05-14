package main

import "fmt"

// 产品接口
type Product interface {
	Use()
}

// 具体产品1
type ConcreteProduct1 struct{}

func (p ConcreteProduct1) Use() {
	fmt.Println("Using ConcreteProduct1")
}

// 具体产品2
type ConcreteProduct2 struct{}

func (p ConcreteProduct2) Use() {
	fmt.Println("Using ConcreteProduct2")
}

// 简单工厂
type SimpleFactory struct{}

func (f SimpleFactory) CreateProduct(productType string) Product {
	switch productType {
	case "product1":
		return &ConcreteProduct1{}
	case "product2":
		return &ConcreteProduct2{}
	default:
		return nil
	}
}

func main() {
	factory := SimpleFactory{}

	product1 := factory.CreateProduct("product1")
	product1.Use()

	product2 := factory.CreateProduct("product2")
	product2.Use()
}
