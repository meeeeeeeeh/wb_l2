package main

import (
	"log"
	"os"
	"os/exec"
	"testing"
)

func TestWget(t *testing.T) {
	url := "https://en.wikipedia.org/wiki/Facade_pattern"

	//создаем директорию где будет лежать файл созданный оригинальной утилитой
	cmdDir := exec.Command("mkdir", "origin_file_directory")
	err := cmdDir.Run()
	if err != nil {
		log.Fatalln("test error: creating directory 'origin_file_directory' failed")
	}

	//переходим в директорию
	err = os.Chdir("origin_file_directory")
	if err != nil {
		log.Fatalln("test error: changing directory failed")
	}

	//вызываем оригинальную  утилиту
	cmd := exec.Command("wget", url)
	// //Устанавливаем директорию, в которой будет выполняться команда
	// cmd.Dir = "test_file_directory"
	_, err = cmd.CombinedOutput()
	if err != nil {
		log.Fatalln("test error: original wget utility failed")
		t.Error()
	}

	//переходим на директорию выше
	err = os.Chdir("../")
	if err != nil {
		log.Fatalln("test error: changing directory failed")
	}

	//создаем директорию где будет лежать файл созданный нашей программой
	cmdDir = exec.Command("mkdir", "test_file_directory")
	//Устанавливаем директорию, в которой будет выполняться команда
	//тк mkdir и chdir выполняются в изолированных процессах и без установки директории
	//chdir не увидит созданную директорию
	err = cmdDir.Run()
	if err != nil {
		log.Fatalln("test error: creating directory 'test_file_directory' failed")
	}

	// //переходим в директорию
	// err = os.Chdir("test_file_directory")
	// if err != nil {
	// 	log.Fatalln("test error: changing directory failed")
	// }

	//вызываем нашу программу
	cmd = exec.Command("./wget_impl", url)
	//Устанавливаем директорию, в которой будет выполняться команда
	cmd.Dir = "test_file_directory"
	_, err = cmd.CombinedOutput()
	if err != nil {
		log.Fatalln("test error: test wget utility failed")
		t.Error()
	}

	// //удаляем созданные директории
	// err = os.RemoveAll("test_file_directory")
	// err = os.RemoveAll("origin_file_directory")
}
