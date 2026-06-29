package pool

import (
	"context"
	"fmt"
	"sync"
	"time"
)

type Job struct {
	Id      int
	Payload string
}

type Result struct {
	JobId    int
	Output   string
	workerId int
}

func GenJobs(JobsChan chan Job, JobsAmount int, parentCtx context.Context) {
	defer close(JobsChan)

	for Id := range JobsAmount {
		select {
		case <-parentCtx.Done():
			fmt.Println("[TASKS]: tasks stopped due timeout")
			return

		case JobsChan <- Job{Id: Id, Payload: "Job generated"}:
			time.Sleep(time.Millisecond * 50)
		}

	}
}
func GenWorker(JobsChan <-chan Job, ResultChan chan<- Result, workerId int, sg *sync.WaitGroup, parentCtx context.Context) {
	defer sg.Done()
	for task := range JobsChan {
		select {
		case <-parentCtx.Done():
			fmt.Println("[Workers]: workers stopped due timeout")
			return
		case ResultChan <- Result{
			JobId:    task.Id,
			Output:   fmt.Sprintf("Job %d finished by Worker with Id: %d", task.Id, workerId),
			workerId: workerId,
		}:
		}

	}
}
