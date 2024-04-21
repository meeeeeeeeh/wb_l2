// Реализовать утилиту аналог консольной команды cut (man cut).
//Утилита должна принимать строки через STDIN, разбивать по разделителю (TAB)
//на колонки и выводить запрошенные.

// Реализовать поддержку утилитой следующих ключей:
// -f - "fields" - выбрать поля (колонки)
// -d - "delimiter" - использовать другой разделитель
// -s - "separated" - только строки с разделителем

package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

// поддержка только одного шаблона, одного файла, одного флага одновременно
type data struct {
	fileName string
	flag     string
	sep      string
	field    int

	//fileByLines []string
}

func newRequest(args []string) (*data, error) {
	var d data

	if len(args) < 2 || len(args) > 5 {
		err := errors.New("Invalid input")
		return nil, err
	}

	if len(args) == 5 && (args[1] == "-A" || args[1] == "-B" || args[1] == "-C") {
		amount, err := strconv.Atoi(args[2])
		if err != nil {
			err := errors.New("Invalid input")
			return nil, err
		}

		d.lineAmount = amount
		d.template = args[3]
		d.flag = args[1]
		d.fileName = args[4]

	} else if len(args) == 4 {
		d.template = args[2]
		d.flag = args[1]
		d.fileName = args[3]

	} else if len(args) == 3 {
		d.template = args[1]
		d.fileName = args[2]
	}

	return &d, nil
}

// заносит каждую строчку файла в слайс
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
		if err != nil && err != io.EOF {
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

func cut() {
	d, err := newRequest(os.Args)
	if err != nil {
		log.Fatalln(err)
	}

	err = parseFile(d)

}

func main() {

	cut()

}
