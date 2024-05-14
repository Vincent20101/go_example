package main

import "fmt"

// 组件接口
type Component interface {
	Operation() string
}

// 具体组件
type ConcreteComponent struct{}

func (c *ConcreteComponent) Operation() string {
	return "ConcreteComponent: Doing something basic."
}

// 抽象装饰器
type Decorator struct {
	component Component
}

func (d *Decorator) Operation() string {
	if d.component != nil {
		return d.component.Operation()
	}
	return ""
}

// 具体装饰器A
type ConcreteDecoratorA struct {
	Decorator
}

func (a *ConcreteDecoratorA) Operation() string {
	return fmt.Sprintf("ConcreteDecoratorA: %s", a.Decorator.Operation())
}

// 具体装饰器B
type ConcreteDecoratorB struct {
	Decorator
}

func (b *ConcreteDecoratorB) Operation() string {
	return fmt.Sprintf("ConcreteDecoratorB: %s", b.Decorator.Operation())
}

func main() {
	// 创建具体组件
	component := &ConcreteComponent{}
	fmt.Println(component.Operation())

	// 使用具体装饰器A装饰组件
	decoratorA := &ConcreteDecoratorA{Decorator{component: component}}
	fmt.Println(decoratorA.Operation())

	// 使用具体装饰器B装饰组件
	decoratorB := &ConcreteDecoratorB{Decorator{component: component}}
	fmt.Println(decoratorB.Operation())

	// 使用具体装饰器A和B依次装饰组件
	decoratorAB := &ConcreteDecoratorA{Decorator{component: decoratorB}}
	fmt.Println(decoratorAB.Operation())
}
