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

// 桌子接口
type Table interface {
	Furniture
	HasDrawer() bool
}

// 抽象工厂接口
type FurnitureFactory interface {
	CreateChair() Chair
	CreateSofa() Sofa
	CreateTable() Table
}

// 现代风格家具
type ModernChair struct{}

func (c ModernChair) GetStyle() string {
	return "Modern"
}

func (c ModernChair) GetBrand() string {
	return "BrandA"
}

func (c ModernChair) HasLegs() bool {
	return true
}

type ModernSofa struct{}

func (s ModernSofa) GetStyle() string {
	return "Modern"
}

func (s ModernSofa) GetBrand() string {
	return "BrandB"
}

func (s ModernSofa) HasCushions() bool {
	return true
}

type ModernTable struct{}

func (t ModernTable) GetStyle() string {
	return "Modern"
}

func (t ModernTable) GetBrand() string {
	return "BrandC"
}

func (t ModernTable) HasDrawer() bool {
	return true
}

// 现代风格家具工厂
type ModernFurnitureFactory struct{}

func (f ModernFurnitureFactory) CreateChair() Chair {
	return ModernChair{}
}

func (f ModernFurnitureFactory) CreateSofa() Sofa {
	return ModernSofa{}
}

func (f ModernFurnitureFactory) CreateTable() Table {
	return ModernTable{}
}

// 古典风格家具
type ClassicChair struct{}

func (c ClassicChair) GetStyle() string {
	return "Classic"
}

func (c ClassicChair) GetBrand() string {
	return "BrandD"
}

func (c ClassicChair) HasLegs() bool {
	return true
}

type ClassicSofa struct{}

func (s ClassicSofa) GetStyle() string {
	return "Classic"
}

func (s ClassicSofa) GetBrand() string {
	return "BrandE"
}

func (s ClassicSofa) HasCushions() bool {
	return true
}

type ClassicTable struct{}

func (t ClassicTable) GetStyle() string {
	return "Classic"
}

func (t ClassicTable) GetBrand() string {
	return "BrandF"
}

func (t ClassicTable) HasDrawer() bool {
	return false
}

// 古典风格家具工厂
type ClassicFurnitureFactory struct{}

func (f ClassicFurnitureFactory) CreateChair() Chair {
	return ClassicChair{}
}

func (f ClassicFurnitureFactory) CreateSofa() Sofa {
	return ClassicSofa{}
}

func (f ClassicFurnitureFactory) CreateTable() Table {
	return ClassicTable{}
}

// 客户端代码
func main() {
	// 创建现代风格家具工厂
	modernFactory := ModernFurnitureFactory{}
	classicFactory := ClassicFurnitureFactory{}

	// 创建椅子
	modernChair := modernFactory.CreateChair()
	fmt.Printf("Modern Chair: Style - %s, Brand - %s, HasLegs - %t\n", modernChair.GetStyle(), modernChair.GetBrand(), modernChair.HasLegs())

	// 创建沙发
	classicSofa := classicFactory.CreateSofa()
	fmt.Printf("Classic Sofa: Style - %s, Brand - %s, HasCushions - %t\n", classicSofa.GetStyle(), classicSofa.GetBrand(), classicSofa.HasCushions())

	// 创建桌子
	modernTable := modernFactory.CreateTable()
	fmt.Printf("Modern Table: Style - %s, Brand - %s, HasDrawer - %t\n", modernTable.GetStyle(), modernTable.GetBrand(), modernTable.HasDrawer())
}
