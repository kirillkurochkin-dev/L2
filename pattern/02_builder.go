package pattern

import "fmt"

/*
	Реализовать паттерн «строитель».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Builder_pattern
*/

type Pizza struct {
	dough    string
	sauce    string
	toppings []string
}

type PizzaBuilder struct {
	pizza *Pizza
}

func NewPizzaBuilder() *PizzaBuilder {
	return &PizzaBuilder{
		pizza: &Pizza{},
	}
}

func (b *PizzaBuilder) SetDough(dough string) *PizzaBuilder {
	b.pizza.dough = dough
	return b
}

func (b *PizzaBuilder) SetSauce(sauce string) *PizzaBuilder {
	b.pizza.sauce = sauce
	return b
}

func (b *PizzaBuilder) AddTopping(topping string) *PizzaBuilder {
	b.pizza.toppings = append(b.pizza.toppings, topping)
	return b
}

func (b *PizzaBuilder) Build() *Pizza {
	return b.pizza
}

func main() {
	pizzaBuilder := NewPizzaBuilder()
	pizza := pizzaBuilder.
		SetDough("тонкое").
		SetSauce("томатный").
		AddTopping("сыр").
		AddTopping("грибы").
		Build()

	fmt.Println("Пицца:", pizza)
}
