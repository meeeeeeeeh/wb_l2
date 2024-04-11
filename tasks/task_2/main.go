// Создать Go-функцию, осуществляющую примитивную распаковку строки, содержащую повторяющиеся символы/руны, например:
// "a4bc2d5e" => "aaaabccddddde"
// "abcd" => "abcd"
// "45" => "" (некорректная строка)
// "" => ""

// Дополнительно
// Реализовать поддержку escape-последовательностей.
// Например:
// qwe\4\5 => qwe45 (*)
// qwe\45 => qwe44444 (*)
// qwe\\5 => qwe\\\\\ (*)

// В случае если была передана некорректная строка, функция должна возвращать ошибку. Написать unit-тесты.

//Go Vet - это инструмент, встроенный в набор инструментов Go, который выполняет 
//проверку на наличие необычных или некорректных использований 
//конструкций языка, таких как многомерные срезы, переполнение буфера и другие

//golint - это инструмент для статического анализа кода на Go, 
//который проверяет соответствие вашего кода рекомендациям из официального
// руководства по написанию кода на Go (Go Code Review Comments)
// golint main.go - для запуска
package main

import (
	"errors"
	//"fmt"
	"strconv"
)

func isNum(val rune) bool {
	return val >= '0' && val <= '9' 
}

func strIncorrect(str string) bool {
	if len(str) == 0 {
		return false
	}
	
	for _, val := range str {
		if !isNum(rune(val)) {
			return false
		}
	}

	return true
}

func unpackage(str string) (string, error) {
	if strIncorrect(str) {
		err := errors.New("incorrect string")
		return "", err
	}

	res := ""
	if len(str) == 0 {
		return res, nil
	}

	for idx := 0; idx < len(str); idx++ {
		// if str[idx] == '/' {
			
		// }
		if idx+1 != len(str) && !isNum(rune(str[idx])) && isNum(rune(str[idx+1])) {
			// запись количества повторов в отдельную строку
			num := ""
			i := 1
			for ; idx+i != len(str) && isNum(rune(str[idx+i])); i++ {
				num += string(str[idx+i])
			}

			n, _ := strconv.Atoi(num)

			for i := 0; i < n; i++ {
				res += string(str[idx])
			}
			idx += (i-1)

		} else {
			res += string(str[idx])
		}
	}

	return res, nil
}

// func main() {
// 	str := "a10b2"

// 	res, err := Unpackage(str)
// 	if err != nil {
// 		fmt.Println(err)
// 	}


// 	fmt.Println(res)
// }