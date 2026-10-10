package worker

import (
	"context"
	"sync"
	"testing"
	"time"

	"github.com/dkrest1/taskflow/internal/model"
)

func TestProcessJobCancellation(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	job := model.NewJob(1, "Test cancellation", 5*time.Second)
	cancel()
	processJob(ctx, 1, job)

	actual := job.GetStatus()
	expected := "cancelled"

	if actual != expected {
		t.Errorf("expected %s, got %s", expected, actual)
	}
}

func TestProcessJobSuccess(t *testing.T) {
	ctx := context.Background()

	job := model.NewJob(1, "Process payment", 10*time.Millisecond)
	processJob(ctx, 1, job)

	jobStatus := job.GetStatus()

	if jobStatus != "completed" {
		t.Errorf("expected completed, got %v", jobStatus)
	}

}

func TestProcessJobCancellationDuringProcessing(t *testing.T) {

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var wg sync.WaitGroup

	job := model.NewJob(1, "Process payment", 2*time.Second)

	wg.Go(func() {
		processJob(ctx, 1, job)
	})

	deadline := time.After(500 * time.Millisecond)

	for job.GetStatus() != "processing" {
		select {
			case <-deadline:
				cancel()
				wg.Wait()
				t.Fatal("job did not start processing")
			default:
				time.Sleep(1 * time.Millisecond)
		}
	}

	cancel()
	wg.Wait()

	actual := job.GetStatus()
	expected := "cancelled"

	if actual != expected {
		t.Errorf("expected %s, got %s", expected, actual)
	}

}

func TestWorkerExistWhenQueueClosed(t *testing.T) {
	ctx := context.Background()

	jobQueue := make(chan *model.Job)

	var wg sync.WaitGroup
	
	wg.Add(1)

	go Worker(ctx, 1, jobQueue, &wg)
	close(jobQueue)

	done := make(chan struct{})

	go func() {
		wg.Wait()
		close(done)
	}()

	select {
		case <- done: 
		case <- time.After(500 * time.Millisecond):
			t.Fatal("worker did not exit after queue closed")

	}







}
