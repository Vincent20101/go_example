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

type simpleFactory struct{}

func (f *simpleFactory) CreateProduct(productType string) Product {
	switch productType {
	case "A":
		return &ConcreteProductA{}
	case "B":
		return &ConcreteProductB{}
	default:
		return nil
	}
}
func main() {
	factory := &simpleFactory{}
	productA := factory.CreateProduct("A")
	productB := factory.CreateProduct("B")
	productA.Use()
	productB.Use()
}
