package main

import (
	"testing"
)

// go test - запуск тестов 
// go test -v - более подробного вывода о процессе тестирования
// файлы с тестами должны заканчиваться на _test.go иначе они не будут найдены 
// тестовые функции должны начинаться с Test иначе тесты могут проходить но они не будут найдены

//покрытие:
// go test -cover ./... - просмотр покрытия
// go test -coverprofile=coverage.out ./... - сохранение покрытия в файл
// go tool cover -html=coverage.out - вывод подробного отчета в html


func TestUnpackage1(t *testing.T) {
	expected := "aaaabccddddde"

	res, err := unpackage("a4bc2d5e")
	if err != nil {
		t.Errorf("unpackage finished with error: %s", err)
	}

	if res != expected {
		t.Errorf("unpackage failed: got %s, expected %s", res, expected)
	}
}

func TestUnpackage2(t *testing.T) {
	expected := "abcd"

	res, err := unpackage("abcd")
	if err != nil {
		t.Errorf("unpackage finished with error: %s", err)
	}

	if res != expected {
		t.Errorf("unpackage failed: got %s, expected %s", res, expected)
	}
}

func TestUnpackage3(t *testing.T) {
	expected := ""

	res, err := unpackage("")
	if err != nil {
		t.Errorf("unpackage finished with error: %s", err)
	}

	if res != expected {
		t.Errorf("unpackage failed: got %s, expected %s", res, expected)
	}
}

func TestUnpackageIncorrectStr(t *testing.T) {
	_, err := unpackage("45")
	if err == nil {
		t.Errorf("unpackage expected error but got nil")
	}

	expectedError := "incorrect string"
	if expectedError != err.Error() {
		t.Errorf("unpackage error mismatch, got: %s, want: %s", err.Error(), expectedError)
	}
}

func TestUnpackageManyNums(t *testing.T) {
	expected := "aaaaaaaaaabbbbbbbbbbb"

	res, err := unpackage("a10b11")
	if err != nil {
		t.Errorf("unpackage finished with error: %s", err)
	}

	if res != expected {
		t.Errorf("unpackage failed: got %s, expected %s", res, expected)
	}
}