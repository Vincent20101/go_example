package main

type Product interface {
	Use()
}

type ConcreteProductA struct{}

func (p *ConcreteProductA) Use() {
	println("Product A is used")
}

type ConcreteProductB struct{}

func (p *ConcreteProductB) Use() {
	println("Product B is used")
}

type Factory interface {
	CreateProduct() Product
}

type ConcreteFactoryA struct{}

func (f *ConcreteFactoryA) CreateProduct() Product {
	return &ConcreteProductA{}
}

type ConcreteFactoryB struct{}

func (f *ConcreteFactoryB) CreateProduct() Product {
	return &ConcreteProductB{}
}

func main() {
	factory := &ConcreteFactoryA{}
	productA := factory.CreateProduct()
	productA.Use()

	// 使用工厂B创建产品B
	factoryB := &ConcreteFactoryB{}
	productB := factoryB.CreateProduct()
	productB.Use()

}
