package main

import "fmt"

// 家具接口
type Furniture interface {
	GetStyle() string
	GetBrand() string
}

// 椅子接口
type Chair interface {
	Furniture
	HasLegs() bool
}

// 沙发接口
type Sofa interface {
	Furniture
	HasCushions() bool
}

// 抽象工厂接口
type FurnitureFactory interface {
	CreateChair() Chair
	CreateSofa() Sofa
}

// 现代风格的椅子家具
type ModernChair struct{}

func (m ModernChair) GetStyle() string {
	return "Modern"
}

func (m ModernChair) GetBrand() string {
	return "Levis"
}

func (m ModernChair) HasLegs() bool {
	return true
}

// 现代风格的沙发家具
type ModernSofa struct{}

func (m ModernSofa) GetStyle() string {
	return "Modern"
}

func (m ModernSofa) GetBrand() string {
	return "Adidas"
}

func (m ModernSofa) HasCushions() bool {
	return true
}

// 现代风格家具工厂
type ModernFurnitureFactory struct{}

func (m ModernFurnitureFactory) CreateChair() Chair {
	return ModernChair{}
}

func (m ModernFurnitureFactory) CreateSofa() Sofa {
	return ModernSofa{}
}

// 古典风格的椅子家具
type ClassicChair struct{}

func (c ClassicChair) GetStyle() string {
	return "Classic"
}

func (c ClassicChair) GetBrand() string {

	return "Fender"
}

func (c ClassicChair) HasLegs() bool {
	return true
}

// 古典风格的沙发家具
type ClassicSofa struct{}

func (c ClassicSofa) GetStyle() string {
	return "Classic"
}

func (c ClassicSofa) GetBrand() string {
	return "Gucci"
}

func (c ClassicSofa) HasCushions() bool {
	return true
}

// 古典风格家具工厂
type ClassicFurnitureFactory struct{}

func (c ClassicFurnitureFactory) CreateChair() Chair {
	return ClassicChair{}
}

func (c ClassicFurnitureFactory) CreateSofa() Sofa {
	return ClassicSofa{}
}

func main() {
	// 创建现代风格家具工厂
	modernFurnitureFactory := ModernFurnitureFactory{}

	// 创建现代风格的椅子
	modernChair := modernFurnitureFactory.CreateChair()
	fmt.Printf("Modern Chair: Style - %s, Brand - %s, HasLegs - %t\n", modernChair.GetStyle(), modernChair.GetBrand(), modernChair.HasLegs())

	// 创建古典风格家具工厂
	classicFurnitureFactory := ClassicFurnitureFactory{}

	// 创建古典风格的椅子
	classicChair := classicFurnitureFactory.CreateChair()
	fmt.Printf("Classic Chair: Style - %s, Brand - %s, HasLegs - %t\n", classicChair.GetStyle(), classicChair.GetBrand(), classicChair.HasLegs())

	// 创建古典风格的沙发
	classicSofa := classicFurnitureFactory.CreateSofa()
	fmt.Printf("Classic Sofa: Style - %s, Brand - %s, HasCushions - %t\n", classicSofa.GetStyle(), classicSofa.GetBrand(), classicSofa.HasCushions())
}
