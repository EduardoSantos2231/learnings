package producers

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func ProducerA(ctx context.Context, dataChan chan<- string, wg *sync.WaitGroup) {
	defer wg.Done()
	counter := 0
	for {
		select {
		case <-ctx.Done():
			return
		default:
			counter++
			time.Sleep(time.Millisecond * 100)
			dataChan <- fmt.Sprintf("{A}-%d", counter)
		}
	}
}
