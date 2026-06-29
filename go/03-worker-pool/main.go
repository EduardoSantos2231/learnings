package main

import (
	"context"
	"flag"
	"fmt"
	"sync"
	"time"
	"03-worker-pool/pool"
)

type opt struct {
	workers int
	jobs    int
	timeout int
}

func main() {
	var syncGroup sync.WaitGroup
	options := colletctFlags()
	jobsDone := 0
	JobsAmount, workersAmount, timeoutAmount := options.jobs, options.workers, options.timeout
	ctx, cancel := createBackgroundContext(timeoutAmount)
	defer cancel()
	JobsChan := make(chan pool.Job, JobsAmount)
	ResultChan := make(chan pool.Result, JobsAmount)
	syncGroup.Add(workersAmount)
	go pool.GenJobs(JobsChan, JobsAmount, ctx)
	for Id := range workersAmount {
		go pool.GenWorker(JobsChan, ResultChan, Id+1, &syncGroup, ctx)
	}
	syncGroup.Wait()
	close(ResultChan)
	for done := range ResultChan {
		fmt.Println("[OUTPUT] ", done.Output)
		jobsDone++
	}
	fmt.Println("We were suposed to run: ", JobsAmount, "but we did: ", jobsDone)
}

func colletctFlags() *opt {
	var options opt
	flag.IntVar(&options.jobs, "jobs", 10, "the amount of tasks to be done")
	flag.IntVar(&options.jobs, "j", 10, "the amount of tasks to be done")
	flag.IntVar(&options.workers, "workers", 3, "the amount of workers to execute the tasks")
	flag.IntVar(&options.workers, "w", 3, "the amount of workers to execute the tasks")
	flag.IntVar(&options.timeout, "timeout", 0, "set the timout for the tasks be completed")
	flag.IntVar(&options.timeout, "t", 0, "set the timout for the tasks be completed")
	flag.Parse()
	return &options
}

func createBackgroundContext(timeout int) (context.Context, context.CancelFunc) {
	if timeout <= 0 {
		return context.Background(), func() {}
	}
	userTime := time.Duration(timeout)
	return context.WithTimeout(context.Background(), time.Second*userTime)
}
