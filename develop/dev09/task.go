// Реализовать утилиту wget с возможностью скачивать сайты целиком.

package main

import (
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

// утилита без флагов
func getUrl() (string, error) {
	if len(os.Args) > 2 {
		err := errors.New("invalid arguments")
		return "", err
	}
	return os.Args[1], nil
}

func createFile(url string) (*os.File, error) {
	temp := strings.Split(url, "/")
	filename := temp[len(temp)-1]
	file, err := os.Create(filename)
	if err != nil {
		return nil, err
	}
	return file, nil
}

func writeToFile(url string, file *os.File) error {
	defer file.Close()
	request, err := http.Get(url)
	if err != nil {
		return err
	}
	if request.StatusCode != 200 {
		err := fmt.Errorf("request failed status %d", request.StatusCode)
		return err
	}
	defer request.Body.Close()

	_, err = io.Copy(file, request.Body)
	if err != nil {
		return err
	}

	return nil
}

func wget() {
	url, err := getUrl()
	if err != nil {
		log.Panic(err)
	}

	file, err := createFile(url)
	if err != nil {
		log.Panic(err)
	}

	err = writeToFile(url, file)
	if err != nil {
		log.Panic(err)
	}
}

func main() {
	wget()
}
