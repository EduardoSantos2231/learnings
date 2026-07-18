package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	chSender := make(chan string, 1)
	go sendOrCancel(chSender, "end", ctx)
	select {
	case msg := <-chSender:
		fmt.Println(msg)
	case <-time.After(time.Second * 5):
		cancel()
	}
}

func sendOrCancel[T any](ch chan<- T, val T, ctx context.Context) error {
	select {
	case ch <- val:
		return nil
	case <-ctx.Done():
		fmt.Println("cancelei")
		return ctx.Err()
	}
}

func fanIn[T any](ctx context.Context, chans ...<-chan T) <-chan T {
	outputChan := make(chan T)
	var wg sync.WaitGroup
	for _, chanToRead := range chans {
		wg.Add(1)
		go readFrom(chanToRead, outputChan, &wg, ctx)
	}
	wg.Wait()
	go func(wg *sync.WaitGroup) {
		wg.Wait()
		close(outputChan)
	}(&wg)
	return outputChan
}

func readFrom[T any](chanToRead <-chan T, chanToWrite chan<- T, wg *sync.WaitGroup, ctx context.Context) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			return
		case msg, ok := <-chanToRead:
			if !ok {
				return
			}
			select {
			case chanToWrite <- msg:
			case <-ctx.Done():
			}
		}
	}
}
