// Создать программу печатающую точное время с использованием NTP -библиотеки.
//Инициализировать как go module. Использовать библиотеку github.com/beevik/ntp.
//Написать программу печатающую текущее время / точное время с использованием этой библиотеки.

// Требования:
// 1) Программа должна быть оформлена как go module
// 2) Программа должна корректно обрабатывать ошибки библиотеки:
//выводить их в STDERR и возвращать ненулевой код выхода в OS

package main

import (
	"fmt"
	"log"
	"time"

	"github.com/beevik/ntp"
)

//по умолчанию логи выводятся в stderr
func printTime() {
	time, err := ntp.Time("pool.ntp.org")
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(time)
}

func main() {
	printTime()
	fmt.Println(time.Now())
}