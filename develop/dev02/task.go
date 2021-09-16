package main

import (
	"fmt"
	"strings"
	"unicode"
)

/*
=== Задача на распаковку ===

Создать Go функцию, осуществляющую примитивную распаковку строки, содержащую повторяющиеся символы / руны, например:
	- "a4bc2d5e" => "aaaabccddddde"
	- "abcd" => "abcd"
	- "45" => "" (некорректная строка)
	- "" => ""
Дополнительное задание: поддержка escape - последовательностей
	- qwe\4\5 => qwe45 (*)
	- qwe\45 => qwe44444 (*)
	- qwe\\5 => qwe\\\\\ (*)

В случае если была передана некорректная строка функция должна возвращать ошибку. Написать unit-тесты.

Функция должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

func RLERevert(line string) (string, error) {
	runeArr := []rune(line)
	sb := strings.Builder{}
	for i := 0; i < len(runeArr)-1; i++ {
		if !unicode.IsDigit(runeArr[i]) {
			if unicode.IsDigit(runeArr[i+1]) {
				sb.WriteString(strings.Repeat(string(runeArr[i]), int(runeArr[i+1]-'0')))
			} else {
				sb.WriteString(strings.Repeat(string(runeArr[i]), 1))
			}
		}
	}

	return sb.String(), nil
}

func main() {
	str := "a4bc2d5e"
	fmt.Println(RLERevert(str))
}
