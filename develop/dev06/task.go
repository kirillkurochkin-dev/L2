package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

/*
=== Утилита cut ===

Принимает STDIN, разбивает по разделителю (TAB) на колонки, выводит запрошенные

Поддержать флаги:
-f - "fields" - выбрать поля (колонки)
-d - "delimiter" - использовать другой разделитель
-s - "separated" - только строки с разделителем

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

func filtr(spliter string) map[int][]string {
	scanner := bufio.NewScanner(os.Stdin)
	var sl []string
	dictionary := make(map[int][]string)

	for i := 0; i < 5; i++ {
		scanner.Scan()
		sl = append(sl, scanner.Text())
	}

	for _, s := range sl {
		splitWords := strings.Split(s, spliter)
		for i2, word := range splitWords {
			dictionary[i2] = append(dictionary[i2], word)
		}
	}

	return dictionary
}

func main() {

	keyFFlag := flag.Bool("f", false, "Выбор поля")
	keyDFlag := flag.Bool("d", false, "Использовать другой разделитель")
	keySFlag := flag.Bool("s", false, "Только строки с разделителем")
	keyParam := flag.String("param", "0", "Параметр для f или d")

	flag.Parse()

	switch true {
	case *keyFFlag:
		if *keyParam == "" {
			fmt.Println("Индекс колонки не указан")
			return
		}

		index, err := strconv.Atoi(*keyParam)

		if err != nil {
			fmt.Println("Индекс колонки задан не числом")
			return
		}

		dictionary := filtr("")

		for _, s := range dictionary[index] {
			fmt.Println(s)
		}
		return

	case *keyDFlag:
		if *keyParam == "" {
			fmt.Println("Разделитель не указан")
			return
		}

		dictionary := filtr(*keyParam)

		for _, s := range dictionary {
			fmt.Println(s)
		}
		return

	case *keySFlag:
		return

	default:
		dictionary := filtr("")
		fmt.Println(dictionary)
		return
	}
}
