// Отсортировать строки в файле по аналогии с консольной утилитой sort
// (man sort — смотрим описание и основные параметры): на входе подается файл из
// несортированными строками, на выходе — файл с отсортированными.

// Реализовать поддержку утилитой следующих ключей:

// -k — указание колонки для сортировки (слова в строке могут выступать в качестве колонок,
//по умолчанию разделитель — пробел)
// -n — сортировать по числовому значению
// -r — сортировать в обратном порядке
// -u — не выводить повторяющиеся строки

// Дополнительно

// Реализовать поддержку утилитой следующих ключей:

// -M — сортировать по названию месяца
// -b — игнорировать хвостовые пробелы
// -c — проверять отсортированы ли данные
// -h — сортировать по числовому значению с учетом суффиксов

package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"sort"
)

//заносит каждую строчку файла в слайс
func parseFile(filename string) ([]string, error) {
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

	return fileData, nil
}

//возвращает слайс с уникальными элементами
func unique(data []string) ([]string) {
	temp := make(map[string]string)
	var res []string

	for _, val := range data {
		_, ok := temp[val] 
		if !ok {
			temp[val] = val
		}
	}

	for _, val := range temp {
		res = append(res, val)
	}

	return res
}

func printData(fileData []string, flag string) {
	if flag == "-r" {
		for i := len(fileData)-1; i >= 0; i-- {
			fmt.Print(fileData[i])
		} 
	} else {
		for _, val := range fileData{
			fmt.Print(val)
		}
	}
}

// поддержка обработки только одного файла и одного флага/без флага одновременно
func sortFile(args []string) error {
	if len(args) < 2 || len(args) > 3 {
		err := errors.New("Invalid input")
		return err
	}

	fileName := args[1]
	flag := ""
	if len(args) == 3  {
		flag = args[2]
	}
	
	
	fileData, err := parseFile(fileName)
	if err != nil {
		return err
	}

	if flag == "-u" {
		fileData = unique(fileData)
	}

	sort.Strings(fileData)

	printData(fileData, flag)

	return nil
}



func main() {

	err := sortFile(os.Args)
	if err != nil {
		fmt.Println(err)
	}
	
}