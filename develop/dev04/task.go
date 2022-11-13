package main

import (
	"fmt"
	"sort"
	"strings"
)

/*
=== Поиск анаграмм по словарю ===

Напишите функцию поиска всех множеств анаграмм по словарю.
Например:
'пятак', 'пятка' и 'тяпка' - принадлежат одному множеству,
'листок', 'слиток' и 'столик' - другому.

Входные данные для функции: ссылка на массив - каждый элемент которого - слово на русском языке в кодировке utf8.
Выходные данные: Ссылка на мапу множеств анаграмм.
Ключ - первое встретившееся в словаре слово из множества
Значение - ссылка на массив, каждый элемент которого, слово из множества. Массив должен быть отсортирован по возрастанию.
Множества из одного элемента не должны попасть в результат.
Все слова должны быть приведены к нижнему регистру.
В результате каждое слово должно встречаться только один раз.

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

func formSetOfAnagrams(words []string) map[string][]string {
	mapOfAnagrams := make(map[string][]string)
L1:
	for _, word := range words {
		word = strings.ToLower(word)
		// если слово в ключе мапы, значит это повтор слова, поэтому скипаем
		if _, ok := mapOfAnagrams[word]; ok {
			continue
		}
		for i, anagrams := range mapOfAnagrams {
			// если слово уже есть в множестве, то скипаем
			for _, anagram := range anagrams {
				if anagram == word {
					continue L1
				}
			}
			// если ключ мапы и текущее слово анаграма, то аппендим слово к множеству
			if checkIfAnagram(word, i) {
				mapOfAnagrams[i] = append(mapOfAnagrams[i], word)
				continue L1
			}
		}
		// тут мы оказались, слово впервые появилось
		mapOfAnagrams[word] = append(mapOfAnagrams[word], word)
	}
	// удаляем множества из одного элемента
	for i, anagrams := range mapOfAnagrams {
		if len(anagrams) == 1 {
			delete(mapOfAnagrams, i)
		}
	}
	return mapOfAnagrams
}

func checkIfAnagram(str1, str2 string) bool {
	str1Slice := strings.Split(str1, "")
	str2Slice := strings.Split(str2, "")
	sort.Strings(str1Slice)
	sort.Strings(str2Slice)
	return strings.Join(str1Slice, "") == strings.Join(str2Slice, "")
}

func main() {
	words := []string{"пятак", "лИсток", "слиток", "слиток", "листок", "пятка", "столик", "тяпка"}
	fmt.Println(formSetOfAnagrams(words))
}
