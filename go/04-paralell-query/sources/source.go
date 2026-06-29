package sources

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Result struct {
	Source      string
	Value       int
	ElapsedTime time.Duration
}

func GenSource(resultChan chan<- Result, callId int, parentCtx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()
	startTime := time.Now()
	randomSleep := rand.Intn(501)
	randomVal := rand.Intn(367)
	time.Sleep(time.Duration(time.Duration(randomSleep) * time.Millisecond))
	select {
	case <-parentCtx.Done():
		return
	case resultChan <- Result{
		Source:      fmt.Sprintf("Source %d", callId),
		Value:       randomVal,
		ElapsedTime: time.Since(startTime),
	}:
	}
}
