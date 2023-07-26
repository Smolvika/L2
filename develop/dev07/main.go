package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var sig = func(after time.Duration) <-chan interface{} {
		c := make(chan interface{})
		go func() {
			defer close(c)
			time.Sleep(after)
		}()
		return c
	}
	start := time.Now()
	// Ожидаем запись в канал, что основанная горутина не завершилась всех остальных
	<-joinChannels(
		sig(1*time.Second),
		sig(2*time.Hour),
		sig(5*time.Minute),
		sig(1*time.Hour),
		sig(1*time.Minute),
	)

	fmt.Printf("fone after %v", time.Since(start))
}

func joinChannels(channels ...<-chan interface{}) <-chan interface{} {
	res := make(chan interface{})
	var wg sync.WaitGroup
	wg.Add(len(channels))
	for _, channel := range channels {
		channel := channel
		go func() {
			for ch := range channel {
				res <- ch
			}
			wg.Done()
		}()
	}
	wg.Wait()
	close(res)
	return res
}
