package pattern

import "fmt"

/*
	Реализовать паттерн «фасад».
Объяснить применимость паттерна, его плюсы и минусы,а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Facade_pattern
*/

// Подсистемы
type Subsystem1 struct{}

func (s *Subsystem1) Operation1() {
	fmt.Println("Subsystem1: Operation1")
}

type Subsystem2 struct{}

func (s *Subsystem2) Operation2() {
	fmt.Println("Subsystem2: Operation2")
}

type Subsystem3 struct{}

func (s *Subsystem3) Operation3() {
	fmt.Println("Subsystem3: Operation3")
}

// Фасад
type Facade struct {
	s1 *Subsystem1
	s2 *Subsystem2
	s3 *Subsystem3
}

func NewFacade(s1 *Subsystem1, s2 *Subsystem2, s3 *Subsystem3) *Facade {
	return &Facade{
		s1: s1,
		s2: s2,
		s3: s3,
	}
}

func (f *Facade) Operation() {
	fmt.Println("Facade: ")
	f.s1.Operation1()
	f.s2.Operation2()
	f.s3.Operation3()
}

func main() {
	s1 := &Subsystem1{}
	s2 := &Subsystem2{}
	s3 := &Subsystem3{}

	facade := NewFacade(s1, s2, s3)
	facade.Operation()
}
