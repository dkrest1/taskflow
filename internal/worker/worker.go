package worker

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/dkrest1/taskflow/internal/model"
)

func Worker(ctx context.Context, workerID int, jobQueue <-chan *model.Job, wg *sync.WaitGroup) {
	defer wg.Done()

	for {
		select {
			case <- ctx.Done():
				fmt.Println("Worker cancel...")
				return

			case job, ok := <- jobQueue:

				if !ok {
					fmt.Printf("Worker %d: job queue closed\n", workerID)
					return
				}

				processJob(ctx ,workerID, job)
		}
			
	}
}

func processJob(ctx context.Context, workerID int, job *model.Job) {
	job.SetStatus("processing")

	fmt.Printf("Worker %d processing Job %d: %s(duration: %v)\n", workerID, job.ID, job.Payload, job.Duration)

	select {
		case <- ctx.Done():
			fmt.Println("Worker cancel...")
			job.SetStatus("cancelled")
			fmt.Printf("Worker %d cancelled job %d\n", workerID, job.ID)
			return
		case <- time.After(job.Duration):
			job.SetStatus("completed")
			fmt.Printf("Worker %d completed job %d\n", workerID, job.ID)

	}
}
