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

// поддержка только одного файла и сепаратора одновременно
// все флаги и значения пишется через пробел
// номера колонок которые нужно вывести пишутся вместе через запятую
type data struct {
	fileName string
	f        bool
	d        bool
	s        bool
	sep      string
	fields   []int
}

const (
	defaultField = 1
	defaultSep   = "\t"
)

func newRequest(args []string) (*data, error) {
	var d data

	for i := range args {
		if args[i][0] == '-' {
			if args[i][1] == 'f' {
				d.f = true
				fields := strings.Split(args[i+1], ",")
				for _, val := range fields {
					field, err := strconv.Atoi(string(val))
					if err != nil {
						return nil, err
					}
					d.fields = append(d.fields, field)
				}
				i++
			} else if args[i][1] == 'd' {
				d.d = true
				d.sep = args[i+1]
				i++
			} else if args[i][1] == 's' {
				d.s = true
			}
		} else if strings.Contains(args[i], ".txt") {
			d.fileName = args[i]
		}
	}

	if !d.f {
		d.fields = append(d.fields, defaultField)
	}
	if !d.d {
		d.sep = defaultSep
	}
	if d.fileName == "" {
		err := errors.New("no file was inputted")
		return nil, err
	}

	return &d, nil
}

func parseFile(d data) error {
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

		if strings.Contains(line, d.sep) {
			output := strings.Split(line, d.sep)

			for idx, val := range d.fields {
				if val > len(output) {
					continue
				}
				if idx != len(d.fields) && idx > 0 {
					fmt.Print(d.sep)
				}
				fmt.Print(output[val-1])

			}
		} else if !d.s {
			fmt.Print(line)
		}

		if err == io.EOF {
			fmt.Print("\n")
			break
		}

	}

	return nil
}

func cut() {
	d, err := newRequest(os.Args)
	if err != nil {
		log.Fatalln(err)
	}

	err = parseFile(*d)
	if err != nil {
		log.Fatalln(err)
	}

}

func main() {
	cut()
}
