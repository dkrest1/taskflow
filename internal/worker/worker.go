package worker

import (
	"fmt"
	"sync"
	"time"

	"github.com/dkrest1/taskflow/internal/model"
)

func Worker(workerID int, jobQueue <-chan *model.Job, wg *sync.WaitGroup) {
	defer wg.Done()
	for job := range jobQueue {
		processJob(workerID, job)
	}
}

func processJob(workerID int, job *model.Job) {
	job.Status = "processing"

	fmt.Printf("Worker %d processing Job %d: %s(duration: %v)\n", workerID, job.ID, job.Payload, job.Duration)
	time.Sleep(job.Duration)
	job.Status = "completed"

	fmt.Printf("Worker %d completed job %d\n", workerID, job.ID)
}
