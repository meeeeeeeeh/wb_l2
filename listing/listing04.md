Что выведет программа? Объяснить вывод программы.

```go
package main

func main() {
	ch := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			ch <- i
		}
	}()

	for n := range ch {
		println(n)
	}
}
```

Ответ:
```
Вывод:
0
1
2
3
4
5
6
7
8
9
fatal error: all goroutines are asleep - deadlock!

В программе создается канал и в цикле вызывается корутина, которая записывает в него значения.
После этого идет цикл, в котором выводятся значения из канала.
При этом новое значение не будет записано, пока предыдущее не выведется 
и цикл не будет остановлен, пока канал не закрыт.
После вывода последнего значения будет ожидаться запись нового, но его не поступает,
поэтому все горутины заснут и возникнет deadlock, о чем и сказано в ошибке
```