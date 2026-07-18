package main

import (
	"context"
	"flag"
	"fmt"
	"04-paralell-query/sources"
	"sync"
	"time"
)

type options struct {
	sources int
	timeout int
}

func main() {
	var syncGroup sync.WaitGroup
	userOptions := colletFlags()
	qtSources, timeout := userOptions.sources, userOptions.timeout
	ctx, cancel := genContext(timeout)
	resultChan := make(chan sources.Result, qtSources)
	syncGroup.Add(qtSources)
	defer cancel()
	go func() {
		syncGroup.Wait()
		close(resultChan)
	}()

	for calls := range userOptions.sources {
		go sources.GenSource(resultChan, calls, ctx, &syncGroup)
	}
	for doneSrc := range resultChan {
		fmt.Printf("[RECEIVED] %s -> %d\n[TIME] %v seconds\n", doneSrc.Source, doneSrc.Value, doneSrc.ElapsedTime.Seconds())
	}
}

func colletFlags() options {
	var opts options
	flag.IntVar(&opts.sources, "sources", 5, "sources to be consulted")
	flag.IntVar(&opts.sources, "s", 5, "sources to be consulted")
	flag.IntVar(&opts.timeout, "timeout", 2, "sources to be consulted")
	flag.IntVar(&opts.timeout, "t", 2, "sources to be consulted")
	flag.Parse()
	return opts
}

func genContext(timeout int) (context.Context, context.CancelFunc) {
	var result = context.Background()
	if timeout <= 0 {
		return result, func() {}
	}
	userTime := time.Duration(timeout)
	return context.WithTimeout(result, time.Second*userTime)
}
