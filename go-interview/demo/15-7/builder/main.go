package main

import "fmt"

// 产品结构体
type Product struct {
	PartA string
	PartB string
	PartC string
}

// 抽象建造者接口
type Builder interface {
	BuildPartA()
	BuildPartB()
	BuildPartC()
	GetProduct() Product
}

// 具体建造者结构体
type ConcreteBuilder struct {
	product Product
}

// 实现建造者接口的方法
func (b *ConcreteBuilder) BuildPartA() {
	b.product.PartA = "PartA built"
}

func (b *ConcreteBuilder) BuildPartB() {
	b.product.PartB = "PartB built"
}

func (b *ConcreteBuilder) BuildPartC() {
	b.product.PartC = "PartC built"
}

func (b *ConcreteBuilder) GetProduct() Product {
	return b.product
}

// 指挥者结构体
type Director struct {
	builder Builder
}

// 指挥者通过建造者构建产品
func (d *Director) Construct() {
	d.builder.BuildPartA()
	d.builder.BuildPartB()
	d.builder.BuildPartC()
}

func main() {
	// 创建具体建造者
	builder := &ConcreteBuilder{}
	// 创建指挥者，并将具体建造者传递给它
	director := &Director{builder: builder}
	// 指挥者使用建造者构建产品
	director.Construct()
	// 获取构建完成的产品
	product := builder.GetProduct()
	// 打印产品信息
	fmt.Printf("Product Parts: %s, %s, %s\n", product.PartA, product.PartB, product.PartC)
}
