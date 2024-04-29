// Необходимо реализовать свой собственный UNIX-шелл-утилиту с поддержкой ряда простейших команд:

// - cd <args> - смена директории (в качестве аргумента могут быть то-то и то)
// - pwd - показать путь до текущего каталога
// - echo <args> - вывод аргумента в STDOUT
// - kill <args> - "убить" процесс, переданный в качесте аргумента (пример: такой-то пример)
// - ps - выводит общую информацию по запущенным процессам в формате *такой-то формат*

// Так же требуется поддерживать функционал fork/exec-команд

// Дополнительно необходимо поддерживать конвейер на пайпах (linux pipes, пример cmd1 | cmd2 | .... | cmdN).

// *Шелл — это обычная консольная программа, которая будучи запущенной, в интерактивном сеансе выводит некое приглашение
// в STDOUT и ожидает ввода пользователя через STDIN. Дождавшись ввода, обрабатывает команду согласно своей логике
// и при необходимости выводит результат на экран. Интерактивный сеанс поддерживается до тех пор, пока не будет введена команда выхода (например \quit).

package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/shirou/gopsutil/process"
)

func kill(args []string) {
	if len(args) > 2 || len(args) <= 1 {
		err := errors.New("invalid arguments")
		log.Fatalln(err)
	}

	pid, err := strconv.Atoi(args[1])
	if err != nil {
		log.Fatalln(err)
	}

	proc, err := os.FindProcess(pid)
	if err != nil {
		log.Fatalln(err)
	}

	err = proc.Kill()
	if err != nil {
		log.Fatalln(err)
	}
}

func ps() {
	proc, err := process.Processes()
	if err != nil {
		log.Fatalln(err)
	}

	for _, val := range proc {
		fmt.Println(val)
	}
}

func echo(args []string) {
	for _, val := range args {
		fmt.Print(val)
	}
}

func pwd() {
	pwd, err := os.Getwd()
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(pwd)
}

func cd(args []string) {
	if len(args) > 2 || len(args) <= 1 {
		err := errors.New("invalid arguments")
		log.Fatalln(err)
	}

	err := os.Chdir(args[1])
	if err != nil {
		log.Fatalln(err)
	}

}

func main() {

	//cd(os.Args)
	// pwd()

	// ps()
	// p := []string{"1", "78"}
	kill(os.Args)
	//ps()
}
