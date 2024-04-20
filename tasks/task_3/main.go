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
	"strconv"
	"strings"
)

type fileData struct {
	fileName string
	flag string
	kVal int
}

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


func sortAsNums(data []string) {
	sort.Slice(data, func(i, j int) bool {
		//убираем все пробелы и перенос строки
		//иначе atoi выдает ошибку
		iTrim := strings.Replace(data[i], " ", "", -1)
		iTrim = strings.Replace(iTrim, "\n", "", -1)
		iConv, iErr := strconv.Atoi(iTrim)

		jTrim := strings.Replace(data[j], " ", "", -1)
		jTrim = strings.Replace(jTrim, "\n", "", -1)
		jConv, jErr := strconv.Atoi(jTrim)

		// если встречаются строки с буквами, они будут сортироваться как строки
		// в ином случае сортировка по числам
		if iErr != nil || jErr != nil {
			return data[i] < data[j]
		}

		return iConv < jConv
	})
}

// флаг к сортируется по определенному столбцу который указывается индексом строки
// столбец для флага к должен идти через пробел!
func sortByKey(data []string, index int) {
	sort.Slice(data, func(i, j int) bool {
		str1 := strings.Split(data[i], " ")
		str2 := strings.Split(data[j], " ")
		
		//если перестановку делать не нужно функции возвращается false
		if index > len(str1) || index > len(str2) {
			return false
		}

		keyWord1 := str1[index-1] 
		keyWord2 := str2[index-1]

		return keyWord1 < keyWord2
	})
}

func newFile(args []string) (*fileData, error) {
	var data fileData
	if len(args) < 2 || len(args) > 4 || (len(args) == 4 && args[1] != "-k") {
		err := errors.New("Invalid input")
		return nil, err
	}

	if len(args) == 4  {
		data.flag = args[1]

		idx, err := strconv.Atoi(args[2])
		if err != nil {
			err := errors.New("Invalid input")
			return nil, err
		}

		data.kVal = idx
		data.fileName = args[3]

	} else if len(args) == 3 {
		data.flag = args[1]
		data.fileName = args[2]

	} else {
		data.fileName = args[1]
	}

	return &data, nil
}

// поддержка обработки только одного файла и одного флага/без флага одновременно
func sortFile() error {
	f, err := newFile(os.Args)
	if err != nil {
		return err
	}
	
	fileData, err := parseFile(f.fileName)
	if err != nil {
		return err
	}

	if f.flag == "-u" {
		fileData = unique(fileData)
	}

	if f.flag == "-n" {
		sortAsNums(fileData)

	} else if f.flag == "-k" {
		sortByKey(fileData, f.kVal)
		
	} else {
		sort.Strings(fileData)
	}

	printData(fileData, f.flag)

	return nil
}




func main() {

	err := sortFile()
	if err != nil {
		fmt.Println(err)
	}
	
}