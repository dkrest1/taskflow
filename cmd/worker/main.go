package main

import (
	"fmt"
	"sync"

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

	for i := 1; i <= 10; i++ {
		newJob := model.NewJob(i, "walking")
		jobQueue <- newJob
		fmt.Printf("queued job %d: len=%d cap=%d\n", i, len(jobQueue), cap(jobQueue))
	}

	close(jobQueue)

	wg.Wait()
	
}


