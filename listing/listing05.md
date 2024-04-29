Что выведет программа? Объяснить вывод программы.

```go
package main

type customError struct {
	msg string
}

func (e *customError) Error() string {
	return e.msg
}

func test() *customError {
	{
		// do something
	}
	return nil
}

func main() {
	var err error
	err = test()
	if err != nil {
		println("error")
		return
	}
	println("ok")
}
```

Ответ:
```
Вывод: 
error

Как и в 5 примере, здесь в поле itab инерфейса err записана информация о типе ошибке 
(customError). В значение err лежит nil, но сам интерфейс не будет равен nil

```