package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"log"
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

type Flags struct {
	f string
	d string
	s bool
}

func readFromStdin() ([]string, error) {
	var stringsSlice []string

	fmt.Println("Enter text:")

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}
		stringsSlice = append(stringsSlice, line)
	}
	if scanner.Err() != nil {
		return nil, scanner.Err()
	}
	return stringsSlice, nil
}

func parseFlags() Flags {
	flags := Flags{}
	flag.StringVar(&flags.f, "f", "", "fields")
	flag.StringVar(&flags.d, "d", "\t", "delimiter")
	flag.BoolVar(&flags.s, "s", false, "separated")
	flag.Parse()
	return flags
}

func parseFields(fields string) ([]int, error) {
	switch {
	case strings.Contains(fields, "-"):
		nums := strings.Split(fields, "-")
		if len(nums) > 2 {
			return nil, errors.New("wrong fields")
		}
		num1, err := strconv.Atoi(nums[0])
		if err != nil {
			return nil, err
		}

		num2, err := strconv.Atoi(nums[1])
		if err != nil {
			return nil, err
		}

		if num1 > num2 {
			return nil, errors.New("wrong fields")
		}

		var res []int
		for i := num1; i <= num2; i++ {
			res = append(res, i)
		}
		return res, nil
	case strings.Contains(fields, ","):
		nums := strings.Split(fields, ",")
		res := make([]int, len(nums))
		for i, v := range nums {
			num, err := strconv.Atoi(v)
			if err != nil {
				return nil, err
			}
			res[i] = num
		}
		return res, nil
	default:
		num, err := strconv.Atoi(fields)
		if err != nil {
			return nil, err
		}
		return []int{num}, nil
	}
}

/*
Поддержать флаги:
-f - "fields" - выбрать поля (колонки)
-d - "delimiter" - использовать другой разделитель
-s - "separated" - только строки с разделителем
*/
func cut(lines []string, flags Flags) ([]string, error) {
	nums, err := parseFields(flags.f)
	if err != nil {
		return nil, err
	}
	var res []string
	for _, line := range lines {
		if flags.s && !strings.Contains(line, flags.d) {
			continue
		}
		if !flags.s && !strings.Contains(line, flags.d) {
			res = append(res, line)
			continue
		}
		if len(nums) == 0 {
			res = append(res, line)
			continue
		}

		parts := strings.Split(line, flags.d)

		var stringsSlice []string
		for _, num := range nums {
			if num > len(parts) {
				continue
			}
			stringsSlice = append(stringsSlice, parts[num-1])
		}
		res = append(res, strings.Join(stringsSlice, flags.d))
	}
	return res, nil
}

func main() {
	flags := parseFlags()
	lines, err := readFromStdin()
	if err != nil {
		log.Fatalln(err)
	}
	res, err := cut(lines, flags)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("Result:")
	fmt.Println(strings.Join(res, "\n"))
}
