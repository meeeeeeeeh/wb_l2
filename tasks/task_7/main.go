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
	"sync"
	"time"
)

func or(channels ...<-chan interface{}) <-chan interface{} {
	c := make(chan interface{})
	var wg sync.WaitGroup
	for _, val := range channels {
		wg.Add(1)
		go func(ch <-chan interface{}) {
			defer wg.Done()
			<-ch
			close(c)

		}(val)
	}

	wg.Wait()
	return c
}

func sig(after time.Duration) chan interface{} {
	c := make(chan interface{})
	go func() {
		defer close(c)
		time.Sleep(after)

	}()
	return c
}

func main() {
	ch1 := sig(1 * time.Second)
	ch1 <- 1
	ch2 := sig(1 * time.Second)
	ch2 <- 2
	start := time.Now()

	<-or(ch1, ch2)

	// <-or(
	// 	sig(1*time.Second),
	// 	sig(2*time.Second),
	// 	sig(3*time.Second),
	// 	sig(4*time.Second),
	// 	sig(5*time.Second),
	// )

	fmt.Printf("done after %v", time.Since(start))
}
