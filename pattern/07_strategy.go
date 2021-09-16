package pattern

import "fmt"

/*
	Реализовать паттерн «стратегия».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Strategy_pattern
*/

type Strategy interface {
	DoOperation(int, int) int
}

type AddStrategy struct{}

func (s *AddStrategy) DoOperation(a, b int) int {
	return a + b
}

type MultiplyStrategy struct{}

func (s *MultiplyStrategy) DoOperation(a, b int) int {
	return a * b
}

type Context struct {
	strategy Strategy
}

func (c *Context) SetStrategy(s Strategy) {
	c.strategy = s
}

func (c *Context) ExecuteStrategy(a, b int) int {
	return c.strategy.DoOperation(a, b)
}

func main() {
	context := &Context{}

	addStrategy := &AddStrategy{}
	context.SetStrategy(addStrategy)
	fmt.Println("Сложение:", context.ExecuteStrategy(3, 4))

	multiplyStrategy := &MultiplyStrategy{}
	context.SetStrategy(multiplyStrategy)
	fmt.Println("Умножение:", context.ExecuteStrategy(3, 4))
}
