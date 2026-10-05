package main

import (
	"fmt"
	"sync"
	"time"

	"github.com/dkrest1/taskflow/internal/model"
)

func main() {
	fmt.Println("Async job processor starting...🚀")

	jobQueue := make(chan *model.Job, 3)

	var wg sync.WaitGroup

	workerCount := 3

	wg.Add(workerCount)

	for i := 1; i <= workerCount; i++ {

		go func (workerID int)  {
			defer wg.Done()
			for job := range jobQueue {
				processJob(workerID, job)
			}

		}(i)
	}

	for i := 1; i <= 10; i++ {
		newJob := model.NewJob(i, "walking")
		jobQueue <- newJob
		fmt.Printf("queued job %d: len=%d cap=%d\n", i, len(jobQueue), cap(jobQueue))
	}

	close(jobQueue)

	wg.Wait()
	
}

func processJob(workerID int, job *model.Job) {
	job.Status = "processing"

	fmt.Printf("Worker %d processing Job %d: %s\n",workerID, job.ID,  job.Payload)
	time.Sleep(2 * time.Second)
	job.Status = "completed"

	fmt.Printf("Worker %d completed job %d\n", workerID, job.ID)
}



