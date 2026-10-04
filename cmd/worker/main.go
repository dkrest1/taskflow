package main

import (
	"fmt"

	"github.com/dkrest1/taskflow/internal/model"
)

func main() {
	fmt.Println("Async job processor starting...🚀")

	newJob := model.NewJob(1, "walk the dog")

	fmt.Printf("NewJob: %+v\n", newJob)

	jobQueue := make(chan *model.Job)

	fmt.Printf("Job queue: %v\n", jobQueue)

	go func() {
		job := <- jobQueue

		fmt.Printf("worker received: %+v", job)
	}()

	jobQueue <- newJob

	
}


