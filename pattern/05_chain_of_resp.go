package pattern

import "fmt"

/*
	Реализовать паттерн «цепочка вызовов».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Chain-of-responsibility_pattern
*/

// Интерфейс обработчика
type Handler interface {
	// Обрабатывает запрос
	Handle(request string) bool
}

// Конкретный обработчик 1
type Handler1 struct{}

func (h *Handler1) Handle(request string) bool {
	if request == "1" {
		fmt.Println("Обработчик 1 обработал запрос")
		return true
	}
	return false
}

// Конкретный обработчик 2
type Handler2 struct{}

func (h *Handler2) Handle(request string) bool {
	if request == "2" {
		fmt.Println("Обработчик 2 обработал запрос")
		return true
	}
	return false
}

// Цепочка обязанностей
type Chain struct {
	handlers []Handler
}

func (c *Chain) AddHandler(handler Handler) {
	c.handlers = append(c.handlers, handler)
}

// Обрабатывает запрос, передавая его по цепочке
func (c *Chain) Handle(request string) {
	for _, handler := range c.handlers {
		if handler.Handle(request) {
			return
		}
	}
	fmt.Println("Запрос не обработан")
}

func main() {
	// Создаем цепочку
	chain := Chain{}

	// Добавляем обработчики
	chain.AddHandler(&Handler1{})
	chain.AddHandler(&Handler2{})

	// Обрабатываем запросы
	chain.Handle("1")
	chain.Handle("2")
	chain.Handle("3")
}
