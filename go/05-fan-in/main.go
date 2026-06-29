package main

import (
	"context"
	"05-fan-in/producers"
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	dataChan := make(chan string, 100)
	chanA := make(chan string, 10)
	chanB := make(chan string, 10)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	wg.Add(3)
	go gracefulShutdown(dataChan, &wg)
	go producers.ProducerA(ctx, chanA, &wg)
	go producers.ProducerB(ctx, chanB, &wg)
	go merger(ctx, chanA, chanB, dataChan, &wg)
	for data := range dataChan {
		fmt.Println(data)
	}
}

func gracefulShutdown(dataChan chan string, wg *sync.WaitGroup) {
	wg.Wait()
	close(dataChan)
}

func merger(ctx context.Context, chanA <-chan string, chanB <-chan string, chanMerged chan<- string, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			return
		case data := <-chanA:
			chanMerged <- data

		case data := <-chanB:
			chanMerged <- data
		}
	}

}
