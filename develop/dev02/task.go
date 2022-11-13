package main

import (
	"errors"
	"fmt"
	"strings"
	"unicode"
	"unicode/utf8"
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

func UnpackString(str string) (string, error) {
	if str == "" {
		return "", nil
	}

	var newStr string
	var lastRune rune
	var isEscaped bool
	for _, v := range str {
		switch {
		case unicode.IsLetter(v) || isEscaped: // либо буква, либо экранированный символ
			newStr += string(v)
			lastRune = v
			isEscaped = false
		case unicode.IsDigit(v) && lastRune != 0: // нулевая руна, если это начало или это мы ее сами занулили после повтора
			if v == '0' { // если ноль, то значит предыдущий символ не выводим
				newStr = trimLastChar(newStr)
			} else {
				newStr += strings.Repeat(string(lastRune), int(v-'0')-1)
			}
			lastRune = 0
		case v == '\\': // отмечаем, что следующий символ экранированный
			isEscaped = true
		}
	}

	if newStr == "" {
		return "", errors.New("incorrect string")
	}
	return newStr, nil
}

func trimLastChar(s string) string {
	r, size := utf8.DecodeLastRuneInString(s)
	if r == utf8.RuneError && (size == 0 || size == 1) {
		size = 0
	}
	return s[:len(s)-size]
}

func main() {
	fmt.Println(UnpackString("a4bc2d5e"))
	fmt.Println(UnpackString("abcd"))
	fmt.Println(UnpackString("45"))
	fmt.Println(UnpackString(""))
	fmt.Println(UnpackString(`qwe\4\5`))
	fmt.Println(UnpackString(`qwe\45`))
	fmt.Println(UnpackString(`qwe\\5`))
}
