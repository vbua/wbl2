package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"reflect"
	"sort"
	"strconv"
	"strings"
)

/*
=== Утилита sort ===

Отсортировать строки в файле по аналогии с консольной утилитой sort (man sort — смотрим описание и основные параметры):
на входе подается файл из несортированными строками, на выходе — файл с отсортированными.

Реализовать поддержку утилитой следующих ключей:

-k — указание колонки для сортировки (слова в строке могут выступать в качестве колонок, по умолчанию разделитель — пробел)
-n — сортировать по числовому значению
-r — сортировать в обратном порядке
-u — не выводить повторяющиеся строки

Дополнительно

Реализовать поддержку утилитой следующих ключей:

-M — сортировать по названию месяца
-b — игнорировать хвостовые пробелы
-c — проверять отсортированы ли данные
-h — сортировать по числовому значению с учетом суффиксов
*/

var (
	k             int
	n, c, r, u, b bool
)

func Sort() {
	flag.IntVar(&k, "k", 0, "Set column name")
	flag.BoolVar(&n, "n", false, "Sort numerically")
	flag.BoolVar(&c, "c", false, "Check if sorted")
	flag.BoolVar(&r, "r", false, "Reverse sorting")
	flag.BoolVar(&u, "u", false, "Remove duplicates")
	flag.BoolVar(&b, "b", false, "Remove trailing blanks")
	flag.Parse()

	data := readFile()

	// делаем 2 слайса: один будет изменяться при сортировке, а второй оригинальный для сравнения по флагу -c
	lines := strings.Split(string(data), "\n")
	originalLines := make([]string, len(lines))
	copy(originalLines, lines)

	// обрезаем справа пробелы
	if b {
		for i, line := range lines {
			lines[i] = strings.TrimRight(line, "")
		}
	}

	// либо по числам сортируем, либо по алфавиту
	if n {
		lines = SortNum(lines)
	} else {
		lines = sortStrings(lines)
	}

	// переворачиваем слайс
	if r {
		lines = reverseSliceOfStrings(lines)
	}

	// сравнение двух слайсов
	if c {
		isEqual := reflect.DeepEqual(lines, originalLines)
		if isEqual {
			fmt.Println("Ordered")
		} else {
			fmt.Println("Not ordered")
		}
		return
	}

	// убираем дубликаты
	if u {
		lines = removeDuplicateStrFromSlice(lines)
	}

	fmt.Println(strings.Join(lines, "\n"))

	text := strings.Join(lines, "\n")
	writeResultToFile(text)
}

func writeResultToFile(text string) {
	err := os.WriteFile("result.txt", []byte(text), 0644)
	if err != nil {
		log.Fatal(err)
	}
}

func readFile() string {
	fileName := os.Args[len(os.Args)-1]
	data, err := os.ReadFile(fileName)
	if err != nil {
		log.Fatal(err)
	}
	return string(data)
}

func reverseSliceOfStrings(ss []string) []string {
	last := len(ss) - 1
	for i := 0; i < len(ss)/2; i++ {
		ss[i], ss[last-i] = ss[last-i], ss[i]
	}
	return ss
}

func removeDuplicateStrFromSlice(strSlice []string) []string {
	allKeys := make(map[string]bool)
	var list []string
	for _, item := range strSlice {
		if _, value := allKeys[item]; !value {
			allKeys[item] = true
			list = append(list, item)
		}
	}
	return list
}

/*
если сортировка по колонке, то все строки без нужной колонки сортируются в отдельном слайсе и кладутся выше тех, которые
можно отсортировать по колонке
*/
func sortStrings(lines []string) []string {
	stringStrings := make(map[string][]string, len(lines))
	strs := make([]string, 0, len(lines))
	columnStrings := make([]string, 0)
	noColumnStrings := make([]string, 0)
	if k == 0 {
		sort.Strings(lines)
	} else {
		for _, line := range lines {
			subStrings := strings.Split(line, " ")
			if k > len(subStrings) {
				noColumnStrings = append(noColumnStrings, line)
				continue
			}
			// не плодим дубликаты в слайсе
			if _, ok := stringStrings[subStrings[k-1]]; !ok {
				strs = append(strs, subStrings[k-1])
			}
			stringStrings[subStrings[k-1]] = append(stringStrings[subStrings[k-1]], line)
		}
		sort.Strings(noColumnStrings)
		sort.Strings(strs)
		for _, v := range strs {
			sort.Strings(stringStrings[v])
			columnStrings = append(columnStrings, stringStrings[v]...)
		}
		lines = append(noColumnStrings, columnStrings...)
	}
	return lines
}

/*
SortNum - если сортировка по колонке, то все строки без нужной колонки сортируются в отдельном слайсе и кладутся выше тех,
которые можно отсортировать по колонке. тоже самое и со строками, где нет числа. безчисловы и безколоночные сортируются
по алфавиту в отдельном слайсе и кладутся вверх.
*/
func SortNum(lines []string) []string {
	numStrings := make(map[float64][]string, len(lines))
	nums := make([]float64, 0, len(lines))
	numericalStrings := make([]string, 0)
	notNumericalStrings := make([]string, 0)
	for _, line := range lines {
		subStrings := strings.Split(line, " ")
		if k > len(subStrings) {
			notNumericalStrings = append(notNumericalStrings, line)
			continue
		}
		num, err := strconv.ParseFloat(subStrings[k-1], 64)
		if err != nil {
			notNumericalStrings = append(notNumericalStrings, line)
			continue
		}
		if _, ok := numStrings[num]; !ok {
			nums = append(nums, num)
		}
		numStrings[num] = append(numStrings[num], line)
	}
	sort.Strings(notNumericalStrings)
	sort.Float64s(nums)
	for _, v := range nums {
		numericalStrings = append(numericalStrings, sortStrings(numStrings[v])...)
	}
	result := append(notNumericalStrings, numericalStrings...)
	return result
}

func main() {
	Sort()
}
