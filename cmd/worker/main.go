package main

import (
	"fmt"
	"sync"

	"github.com/dkrest1/taskflow/internal/model"
)

func main() {
	fmt.Println("Async job processor starting...🚀")

	newJob1 := model.NewJob(1, "walk the dog")
	newJob2 := model.NewJob(2, "send email")
	newJob3 := model.NewJob(3, "generate report")

	jobQueue := make(chan *model.Job)

	var wg sync.WaitGroup

	wg.Add(1)

	go func() {
		defer wg.Done()
		
		for job := range jobQueue {
			fmt.Printf("worker received: %+v\n", job)
		}

	}()

	jobQueue <- newJob1
	jobQueue <- newJob2
	jobQueue <- newJob3
	close(jobQueue)


	wg.Wait()

	workerCount := 3

	for i := 1; i <= workerCount; i++ {
		wg.Add(3)

		go func ()  {
			defer wg.Done()
		}()
	}
	
}


