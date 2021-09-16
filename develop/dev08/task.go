package main

import (
	"bufio"
	"fmt"
	os "os"
	"os/exec"
	"strconv"
	"strings"
)

/*
=== Взаимодействие с ОС ===

# Необходимо реализовать собственный шелл

встроенные команды: cd/pwd/echo/kill/ps
поддержать fork/exec команды
конвеер на пайпах

Реализовать утилиту netcat (nc) клиент
принимать данные из stdin и отправлять в соединение (tcp/udp)
Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		args := strings.Split(line, " ")
		switch args[0] {
		case "cd":
			os.Chdir(args[1])
		case "pwd":
			fmt.Println(os.Getwd())
		case "echo":
			fmt.Println(strings.Join(args[1:], " "))
		case "kill":
			_, err := strconv.Atoi(args[1])
			if err != nil {
				fmt.Println("Invalid PID:", args[1])
				continue
			}
		case "ps":
			cmd := exec.Command("ps", "-o", "pid,user,command")
			out, err := cmd.Output()
			if err != nil {
				fmt.Println("Error:", err)
				continue
			}
			fmt.Println(string(out))
		default:
			fmt.Println("Unknown command:", args[0])
		}
	}
}
