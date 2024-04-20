// Написать функцию поиска всех множеств анаграмм по словарю.

// Например:
// 'пятак', 'пятка' и 'тяпка' - принадлежат одному множеству,
// 'листок', 'слиток' и 'столик' - другому.

// Требования:
// 1. Входные данные для функции: ссылка на массив, каждый элемент которого - слово на русском языке в кодировке utf8
// 2. Выходные данные: ссылка на мапу множеств анаграмм
// 3. Ключ - первое встретившееся в словаре слово из множества. Значение - ссылка на массив, каждый элемент которого,
// слово из множества.
// 4. Массив должен быть отсортирован по возрастанию.
// 5. Множества из одного элемента не должны попасть в результат.
// 6. Все слова должны быть приведены к нижнему регистру.
// 7. В результате каждое слово должно встречаться только один раз.

package main

import (
	"sort"
	"strings"
)

func isAnagrams(str1, str2 string) bool {
	if len(str1) != len(str2) {
		return false
	}

	tempMap := make(map[rune]int)

	for _, val := range str1 {
		_, ok := tempMap[val] 
		if !ok {
			tempMap[val] = 1
		} else {
			tempMap[val] += 1
		}
	}

	for _, val := range str2 {
		_, ok := tempMap[val]
		if !ok {
			return false
		}
		tempMap[val] -= 1

	}

	for _, val := range tempMap {
		if val != 0 {
			return false
		}
	}

	return true
}

func sortMap(data map[string][]string) {
	for key, val := range data {
		if len(val) == 1 {
			delete(data, key)
		} else {
			sort.Strings(val)
		}
	}
}

func isExist(data []string, item string) bool {
	for _, val := range data {
		if item == val {
			return true
		}
	}

	return false
}


func getAnagrams(str string) map[string][]string {
	str = strings.ToLower(str)
	data := strings.Split(str, " ")
	res := make(map[string][]string)

	for _, val := range data {
		added := false

		for key, _ := range res {
			if isExist(res[key], val) {
				added = true
				break
			}

			if isAnagrams(key, val){
				res[key] = append(res[key], val)
				added = true
			}
		}

		if !added {
			newSet := []string{val}
			res[val] = newSet
		}
	}

	sortMap(res)

	return res
}