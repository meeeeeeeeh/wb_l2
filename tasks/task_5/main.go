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
	"log"
	"os"
	"strconv"
	"strings"
)

// поддержка только одного шаблона, одного файла, одного флага одновременно
type data struct {
	fileName    string
	flag        string
	template    string
	lineAmount  int
	fileByLines []string
}

func newRequest(args []string) (*data, error) {
	var d data

	if len(args) < 2 || len(args) > 5 {
		err := errors.New("invalid input")
		return nil, err
	}

	if len(args) == 5 && (args[1] == "-A" || args[1] == "-B" || args[1] == "-C") {
		amount, err := strconv.Atoi(args[2])
		if err != nil {
			err := errors.New("invalid input")
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

// заносит каждую строчку файла в слайс и добавляет в структуру
func parseFile(d *data) error {
	f, err := os.Open(d.fileName)
	if err != nil {
		return err
	}
	defer f.Close()

	reader := bufio.NewReader(f)
	for {
		line, err := reader.ReadString('\n')
		if err != nil && err != io.EOF {
			return err
		}
		if err == io.EOF {
			line += "\n"
			d.fileByLines = append(d.fileByLines, line)
			break
		}
		d.fileByLines = append(d.fileByLines, line)
	}

	return nil
}

func findMatches(d *data) {
	var newFileData []string

	for i, line := range d.fileByLines {
		if strings.Contains(line, d.template) {
			if d.flag == "-A" {

				for idx := i; idx <= len(d.fileByLines)-1 && idx <= d.lineAmount+i; idx++ {
					newFileData = append(newFileData, d.fileByLines[idx])
				}
				i += d.lineAmount

			} else if d.flag == "-B" {
				idx := i - d.lineAmount
				if idx < 0 {
					newFileData = append(newFileData, d.fileByLines[:i+1]...)

				} else {

					for ; idx <= i; idx++ {
						newFileData = append(newFileData, d.fileByLines[idx])
					}
				}

			} else if d.flag == "-C" {
				idx := i - d.lineAmount
				if idx < 0 {
					newFileData = append(newFileData, d.fileByLines[:i+1]...)

				} else {
					for idx := i - d.lineAmount; idx <= len(d.fileByLines)-1 && idx >= 0 && idx <= d.lineAmount+i; idx++ {
						newFileData = append(newFileData, d.fileByLines[idx])
					}

				}
				i += d.lineAmount
			}

		}
	}

	d.fileByLines = newFileData
}

func printLines(d data) {
	if d.flag == "-A" || d.flag == "-B" || d.flag == "-C" {
		findMatches(&d)
		for _, line := range d.fileByLines {
			fmt.Print(line)
		}
	} else {
		count := 0
		for i, line := range d.fileByLines {
			match := false

			if d.flag == "-i" {
				match = strings.Contains(strings.ToLower(line), strings.ToLower(d.template))
			} else if d.flag == "-F" {
				lineTrim := strings.Replace(line, " ", "", -1)
				lineTrim = strings.Replace(lineTrim, "\n", "", -1)

				templTrim := strings.Replace(d.template, " ", "", -1)
				templTrim = strings.Replace(templTrim, "\n", "", -1)

				match = lineTrim == templTrim
			} else {
				match = strings.Contains(line, d.template)
			}

			if match && d.flag != "-v" {
				if d.flag == "-n" {
					fmt.Println(i + 1)
				} else if d.flag != "-c" {
					fmt.Print(line)
				}

				count++
			}

			if d.flag == "-v" && !match {
				fmt.Print(line)
			}

		}

		if d.flag == "-c" {
			fmt.Print(count)
		}
	}

}

func grep() {
	d, err := newRequest(os.Args)
	if err != nil {
		log.Fatalln(err)
	}

	err = parseFile(d)
	if err != nil {
		log.Fatalln(err)
	}

	printLines(*d)

}

func main() {

	grep()

}
