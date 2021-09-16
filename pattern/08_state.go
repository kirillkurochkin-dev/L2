package pattern

import "fmt"

/*
	Реализовать паттерн «состояние».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/State_pattern
*/

// Интерфейс состояния
type State interface {
	DoAction()
}

// Состояние "Включено"
type OnState struct{}

func (s *OnState) DoAction() {
	fmt.Println("Включено")
}

// Состояние "Выключено"
type OffState struct{}

func (s *OffState) DoAction() {
	fmt.Println("Выключено")
}

// Контекст
type Tex struct {
	state State
}

// Setter для изменения состояния
func (c *Tex) SetState(s State) {
	c.state = s
}

// Выполнение действия
func (c *Tex) DoAction() {
	c.state.DoAction()
}

func main() {
	context := &Tex{
		state: &OffState{},
	}

	context.DoAction() // Выводит "Выключено"

	context.SetState(&OnState{})
	context.DoAction() // Выводит "Включено"
}
