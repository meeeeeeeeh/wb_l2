// Реализовать функцию, которая будет объединять один или более done-каналов
// в single-канал, если один из его составляющих каналов закроется.
// Очевидным вариантом решения могло бы стать выражение при использованием select,
// которое бы реализовывало эту связь, однако иногда неизвестно общее число done-каналов,
// с которыми вы работаете в рантайме.
// В этом случае удобнее использовать вызов единственной функции, которая,
// приняв на вход один или более or-каналов, реализовывала бы весь функционал.

package main

import (
	"fmt"
	"time"
)

func or(channels ...<-chan interface{}) <-chan interface{} {
	res := make(chan interface{})
	for _, val := range channels {
		go func(ch <-chan interface{}) {
			<-ch
			close(res)
		}(val)
	}
	return res
}

func sig(after time.Duration) <-chan interface{} {
	c := make(chan interface{}) // кладет nil
	go func() {
		defer close(c)
		time.Sleep(after)
	}()
	return c
}

func main() {
	start := time.Now()
	<-or(
		sig(1*time.Microsecond),
		sig(2*time.Second),
		sig(3*time.Second),
		sig(4*time.Second),
	)

	fmt.Printf("done after %v", time.Since(start))

}
