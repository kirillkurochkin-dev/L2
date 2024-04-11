package pattern

<<<<<<< HEAD
import "fmt"

=======
>>>>>>> origin/main
/*
	Реализовать паттерн «комманда».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Command_pattern
*/
<<<<<<< HEAD

// Интерфейс команды
type Command interface {
	Execute()
}

// Команда "Вкл"
type OnCommand struct {
	device *Device
}

func (c *OnCommand) Execute() {
	fmt.Println("Включение устройства")
	c.device.On()
}

// Команда "Выкл"
type OffCommand struct {
	device *Device
}

func (c *OffCommand) Execute() {
	fmt.Println("Выключение устройства")
	c.device.Off()
}

// Устройство
type Device struct {
	on bool
}

func (d *Device) On() {
	d.on = true
}

func (d *Device) Off() {
	d.on = false
}

// Пульт управления
type Invoker struct {
	commands []Command
}

func (i *Invoker) AddCommand(command Command) {
	i.commands = append(i.commands, command)
}

func (i *Invoker) ExecuteCommands() {
	for _, command := range i.commands {
		command.Execute()
	}
}

func main() {
	// Создание устройства
	device := &Device{}

	// Создание команд
	onCommand := &OnCommand{device: device}
	offCommand := &OffCommand{device: device}

	// Создание пульта управления
	invoker := &Invoker{}

	// Добавление команд на пульт
	invoker.AddCommand(onCommand)
	invoker.AddCommand(offCommand)

	// Выполнение команд
	invoker.ExecuteCommands()
}
=======
>>>>>>> origin/main
