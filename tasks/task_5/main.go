// Реализовать утилиту фильтрации по аналогии с консольной утилитой (man grep — смотрим описание и основные параметры).


// Реализовать поддержку утилитой следующих ключей:
// -A - "after" печатать +N строк после совпадения
// -B - "before" печатать +N строк до совпадения
// -C - "context" (A+B) печатать ±N строк вокруг совпадения
// -c - "count" (количество строк)
// -i - "ignore-case" (игнорировать регистр)
// -v - "invert" (вместо совпадения, исключать)
// -F - "fixed", точное совпадение со строкой, не паттерн
// -n - "line num", напечатать номер строки

package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
)

func main() {

	f, err := os.Open(filename) 
	if err != nil {
		return nil, err
	}
	defer f.Close()

	reader := bufio.NewReader(f)
	var fileData []string

	for {
		line, err := reader.ReadString('\n')
		if err != nil && err != io.EOF{
			return nil, err
		}
		if err == io.EOF {
			line += "\n"
			fileData = append(fileData, line)
			break
		}
		fileData = append(fileData, line)
	}
	
}