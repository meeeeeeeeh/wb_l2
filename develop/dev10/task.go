/*
Реализовать простейший telnet-клиент.

Примеры вызовов:
go-telnet --timeout=10s host port go-telnet mysite.ru 8080 go-telnet --timeout=3s 1.1.1.1 123

Требования:
1. Программа должна подключаться к указанному хосту (ip или доменное имя + порт)
по протоколу TCP. После подключения STDIN программы должен записываться в сокет, а данные полученные и
сокета должны выводиться в STDOUT
2. Опционально в программу можно передать таймаут на подключение к серверу
(через аргумент --timeout, по умолчанию 10s)
3. При нажатии Ctrl+D программа должна закрывать сокет и завершаться. Если сокет закрывается
со стороны сервера, программа должна также завершаться. При подключении к несуществующему сервер,
программа должна завершаться через timeout
*/

//вначале вызывается netcat для создания сервера кторый будет прослушивать определенный порт
// nc -lp 1234

package main

import (
	"bufio"
	"context"
	"errors"
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"time"
)

type client struct {
	time time.Duration
	host string
	port string
}

func newClient() (*client, error) {
	//парсит строку и ищет там лаг timeout, если не будет найден то знавение 10сек
	time := flag.Duration("timeout", 10*time.Second, "getting time of server working")
	flag.Parse()

	if len(os.Args) < 2 {
		err := errors.New("host or port name wasn't specified")
		return nil, err
	}

	//возвращает аргументы командной строки котрык не являются флагами
	args := flag.Args()

	return &client{
		time: *time,
		host: args[0],
		port: args[1],
	}, nil
}

func newConnection() error {
	client, err := newClient()
	if err != nil {
		return err
	}

	connect, err := net.DialTimeout("tcp", client.host+":"+client.port, client.time)
	if err != nil {
		return err
	}

	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, client.time)
	defer cancel()

	for {
		select {
		case <-ctx.Done():
			fmt.Println("Connection time is out")
			connect.Close()
			return nil
		default:
			//ввод из stdin записывается в соединение
			reader := bufio.NewReader(os.Stdin)
			text, err := reader.ReadString('\n')
			if err != nil {
				return err
			}
			fmt.Fprintln(connect, text)

			//ввод из конекта записывается в stdin
			reader = bufio.NewReader(connect)
			text, err = reader.ReadString('\n')
			if err != nil {
				return err
			}
			fmt.Println(text)

		}
	}

}

func main() {
	err := newConnection()
	if err != nil {
		log.Fatalln(err)
	}
}
