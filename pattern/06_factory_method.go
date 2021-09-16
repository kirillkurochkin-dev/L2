package pattern

import "fmt"

/*
	Реализовать паттерн «фабричный метод».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Factory_method_pattern
*/

type IProduct interface {
	GetName() string
}

type ProductA struct {
	name string
}

func (p *ProductA) GetName() string {
	return p.name
}

type ProductB struct {
	name string
}

func (p *ProductB) GetName() string {
	return p.name
}

type ICreator interface {
	CreateProduct() IProduct
}

type CreatorA struct{}

func (c *CreatorA) CreateProduct() IProduct {
	return &ProductA{name: "ProductA"}
}

type CreatorB struct{}

func (c *CreatorB) CreateProduct() IProduct {
	return &ProductB{name: "ProductB"}
}

func main() {
	creatorA := &CreatorA{}
	productA := creatorA.CreateProduct()
	fmt.Println(productA.GetName()) // "ProductA"

	creatorB := &CreatorB{}
	productB := creatorB.CreateProduct()
	fmt.Println(productB.GetName()) // "ProductB"
}
