package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
)

/*
=== Утилита sort ===

Отсортировать строки (man sort)
Основное

Поддержать ключи

-k — указание колонки для сортировки
-n — сортировать по числовому значению
-r — сортировать в обратном порядке
-u — не выводить повторяющиеся строки

Дополнительное

Поддержать ключи

-M — сортировать по названию месяца
-b — игнорировать хвостовые пробелы
-c — проверять отсортированы ли данные
-h — сортировать по числовому значению с учётом суффиксов

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

// Сортировка для ключа KeyK
type ByWordPos struct {
	data      []string
	wordIndex int
}

func (a ByWordPos) Len() int {
	return len(a.data)
}
func (a ByWordPos) Swap(i, j int) {
	a.data[i], a.data[j] = a.data[j], a.data[i]
}
func (a ByWordPos) Less(i, j int) bool {
	wordsI := strings.Fields(a.data[i])
	wordsJ := strings.Fields(a.data[j])

	// Проверяем, достаточно ли слов в каждой строке
	if len(wordsI) <= a.wordIndex || len(wordsJ) <= a.wordIndex {
		return false // По умолчанию, если нет указанного слова, не меняем порядок
	}

	// Сравниваем слова по индексу
	return wordsI[a.wordIndex] < wordsJ[a.wordIndex]
}

// Ключи

// Отсортировать по заданной колонке
func keyK(sl []string, wordIndex int) ([]string, error) {
	sort.Sort(ByWordPos{data: sl, wordIndex: wordIndex})
	return sl, nil
}

// Отсортировать как числа
func keyN(sl []string) ([]string, error) {
	sort.SliceStable(sl, func(i, j int) bool {
		numI, errI := strconv.Atoi(sl[i])
		numJ, errJ := strconv.Atoi(sl[j])

		if errI != nil || errJ != nil {
			return sl[i] < sl[j]
		}
		return numI < numJ
	})

	return sl, nil
}

// Отсортировать в обратном порядке
func keyR(sl []string) ([]string, error) {
	for i, j := 0, len(sl)-1; i < j; i, j = i+1, j-1 {
		sl[i], sl[j] = sl[j], sl[i]
	}
	return sl, nil
}

// Отсортировать и удалить повторяющиеся
func keyU(sl []string) ([]string, error) {
	dict := make(map[string]bool)

	var unique []string
	for _, s := range sl {
		if !dict[s] {
			unique = append(unique, s)
			dict[s] = true
		}
	}

	return unique, nil
}

// Простая сортировка в лексикографическом порядке
func deflt(sl []string) ([]string, error) {
	sort.Strings(sl)
	return sl, nil
}

func main() {
	// Флаги командной строки
	inputFilePath := flag.String("path", "", "Путь к файлу для сортировки")
	keyKFlag := flag.Bool("k", false, "Сортировка по ключу")
	keyNFlag := flag.Bool("n", false, "Сортировка по числовому значению")
	keyRFlag := flag.Bool("r", false, "Обратная сортировка")
	keyUFlag := flag.Bool("u", false, "Уникальные значения")
	keyIndex := flag.Int("index", -1, "Индекс колонки для ключа -k")

	flag.Parse()

	// Чтение файла
	file, err := os.Open(*inputFilePath)
	if err != nil {
		fmt.Println("Ошибка открытия файла:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Ошибка чтения файла:", err)
		return
	}

	// Применение ключа сортировки
	switch true {
	case *keyKFlag:
		if *keyIndex < 0 {
			fmt.Println("Индекс колонки не указан для ключа -k")
			return
		}
		lines, err = keyK(lines, *keyIndex)
	case *keyNFlag:
		lines, err = keyN(lines)
	case *keyRFlag:
		lines, err = keyR(lines)
	case *keyUFlag:
		lines, err = keyU(lines)
	default:
		lines, err = deflt(lines)
	}

	if err != nil {
		fmt.Println("Ошибка сортировки:", err)
		return
	}

	// Запись отсортированных данных в файл
	directory := filepath.Dir(*inputFilePath)
	outputFilePath := filepath.Join(directory, "output.txt")
	outputFile, err := os.Create(outputFilePath)
	if err != nil {
		fmt.Println("Ошибка создания файла:", err)
		return
	}
	defer outputFile.Close()

	writer := bufio.NewWriter(outputFile)
	for _, line := range lines {
		line += "\n"
		writer.WriteString(line)
	}
	writer.Flush()

	fmt.Println("Сортировка завершена. Результат записан в output.txt")
}
