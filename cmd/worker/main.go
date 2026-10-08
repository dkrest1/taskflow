package main

import (
	"fmt"
	"sync"
	"time"

	"github.com/dkrest1/taskflow/internal/model"
	"github.com/dkrest1/taskflow/internal/worker"
)

func main() {
	fmt.Println("Async job processor starting...🚀")

	jobQueue := make(chan *model.Job, 3)

	var wg sync.WaitGroup

	workerCount := 3

	wg.Add(workerCount)

	for i := 1; i <= workerCount; i++ {

		go worker.Worker(i, jobQueue, &wg)

	}

	var jobs []*model.Job

	for i := 1; i <= 10; i++ {

		durationSeconds := (i-1)%3 + 1
		duration := time.Duration(durationSeconds) * time.Second

		newJob := model.NewJob(i, "walking", duration)

		jobs = append(jobs, newJob)

		jobQueue <- newJob
		fmt.Printf("queued job %d: len=%d cap=%d\n", i, len(jobQueue), cap(jobQueue))
	}

	close(jobQueue)

	stopMonitor := make(chan struct{})

	var monitorWG sync.WaitGroup
	monitorWG.Add(1)

	go func() {
		defer monitorWG.Done()
		for {
			select {
				case <- stopMonitor:
					return
				default:
					for _, job := range jobs {
						fmt.Printf("JobID: %v, Job Status: %v\n", job.ID, job.GetStatus())
					} 

			}

			time.Sleep(100 * time.Millisecond)
		}
	}()

	wg.Wait()
	close(stopMonitor)
	monitorWG.Wait()

}
