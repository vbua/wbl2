package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

/*
=== Утилита grep ===

Реализовать утилиту фильтрации (man grep)

Поддержать флаги:
-A - "after" печатать +N строк после совпадения
-B - "before" печатать +N строк до совпадения
-C - "context" (A+B) печатать ±N строк вокруг совпадения
-c - "count" (количество строк)
-i - "ignore-case" (игнорировать регистр)
-v - "invert" (вместо совпадения, исключать)
-F - "fixed", точное совпадение со строкой, не паттерн
-n - "line num", печатать номер строки

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

type Flags struct {
	i, c, v, n, F     bool
	A, B, C           int
	fileName, pattern string
}

func parseFlags() Flags {
	flags := Flags{}
	flag.BoolVar(&flags.i, "i", false, "Ignore case")
	flag.IntVar(&flags.A, "A", 0, "Print N lines after match")
	flag.IntVar(&flags.B, "B", 0, "Print N lines before match")
	flag.IntVar(&flags.C, "C", 0, "Print N lines around match")
	flag.BoolVar(&flags.c, "c", false, "Print number of matches")
	flag.BoolVar(&flags.v, "v", false, "Invert")
	flag.BoolVar(&flags.n, "n", false, "Print line number")
	flag.BoolVar(&flags.F, "F", false, "Fixed")
	flag.Parse()

	flags.pattern = os.Args[len(os.Args)-2:][0]
	flags.fileName = os.Args[len(os.Args)-1]

	return flags
}

func grep(flags Flags, lines []string) {
	mapMatchedLines := make(map[int]struct{}, 0)
	matchedLinesSlice := make([]int, 0)
	allLinesNumbers := make([]int, 0)
	for lineNumber, line := range lines {
		contains := false
		if flags.i {
			line = strings.ToLower(line)
			flags.pattern = strings.ToLower(flags.pattern)
		}
		if flags.F {
			contains = line == flags.pattern
		} else {
			contains = strings.Contains(line, flags.pattern)
		}

		if (!flags.v && contains) || (flags.v && !contains) {
			mapMatchedLines[lineNumber] = struct{}{}
			matchedLinesSlice = append(matchedLinesSlice, lineNumber)
		}

		allLinesNumbers = append(allLinesNumbers, lineNumber)
	}

	// просто выводим количество совпадений
	if flags.c {
		fmt.Println(len(matchedLinesSlice))
		return
	}

	// A и B имеют преимущество перед C, но если A или B не установлены, то работает C
	if flags.C != 0 && flags.A == 0 {
		flags.A = flags.C
	}
	if flags.C != 0 && flags.B == 0 {
		flags.B = flags.C
	}

	// проверяем флаги на каждом совпадении и дополняем вывод, если нужно
	for i, line := range matchedLinesSlice {
		// надо вытащить A строк ПОСЛЕ текущей строки
		if flags.A > 0 {
			var nextMatchedLineNumber int
			if i+1 <= len(matchedLinesSlice)-1 {
				nextMatchedLineNumber = matchedLinesSlice[i+1]
			} else {
				nextMatchedLineNumber = len(lines) - 1
			}
			lineAfter := getLinesAfter(flags.A, line, nextMatchedLineNumber, allLinesNumbers)
			mapMatchedLines[line] = struct{}{}
			for _, v := range lineAfter {
				mapMatchedLines[v] = struct{}{}
			}
		}

		// надо вытащить B строк ДО текущей строки
		if flags.B > 0 {
			var previousMatchedLineNumber int
			if i != 0 {
				previousMatchedLineNumber = matchedLinesSlice[i-1]
			} else {
				previousMatchedLineNumber = i
			}
			lineBefore := getLinesBefore(flags.B, line, previousMatchedLineNumber, allLinesNumbers)
			mapMatchedLines[line] = struct{}{}
			for _, v := range lineBefore {
				mapMatchedLines[v] = struct{}{}
			}
		}
	}

	// вывод результата построчно
	for i, v := range lines {
		if _, ok := mapMatchedLines[i]; ok {
			if flags.n {
				fmt.Printf("%v: %v\n", i+1, v)
			} else {
				fmt.Println(v)
			}
		}
	}
}

func getLinesBefore(numLinesBefore, currenLine, beforeLine int, lines []int) []int {
	diff := currenLine - beforeLine
	if diff > numLinesBefore {
		return lines[currenLine-numLinesBefore : currenLine]
	} else if diff > 0 {
		return lines[currenLine-diff+1 : currenLine]
	}
	return []int{}
}

func getLinesAfter(numLinesAfter, currenLine, nextLine int, lines []int) []int {
	diff := nextLine - currenLine
	if diff > numLinesAfter {
		return lines[currenLine+1 : currenLine+numLinesAfter+1]
	} else if diff > 0 {
		return lines[currenLine+1 : currenLine+diff+1]
	}
	return []int{}
}

func readFileToStrings(fileName string) []string {
	data, err := os.ReadFile(fileName)
	if err != nil {
		log.Fatal(err)
	}
	lines := strings.Split(string(data), "\n")
	return lines
}

func main() {
	flags := parseFlags()
	lines := readFileToStrings(flags.fileName)
	grep(flags, lines)
}
