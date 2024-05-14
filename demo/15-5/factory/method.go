package main

import "fmt"

// 产品接口
type Product interface {
	Use()
}

// 具体产品A
type ConcreteProductA struct{}

func (p ConcreteProductA) Use() {
	fmt.Println("Using ConcreteProductA")
}

// 具体产品B
type ConcreteProductB struct{}

func (p ConcreteProductB) Use() {
	fmt.Println("Using ConcreteProductB")
}

// 工厂接口
type Factory interface {
	CreateProduct() Product
}

// 具体工厂A
type ConcreteFactoryA struct{}

func (f ConcreteFactoryA) CreateProduct() Product {
	return ConcreteProductA{}
}

// 具体工厂B
type ConcreteFactoryB struct{}

func (f ConcreteFactoryB) CreateProduct() Product {
	return ConcreteProductB{}
}

func main() {
	// 使用工厂A创建产品A
	var factory Factory = ConcreteFactoryA{}
	productA := factory.CreateProduct()
	productA.Use()

	// 使用工厂B创建产品B
	factory = ConcreteFactoryB{}
	productB := factory.CreateProduct()
	productB.Use()
}
