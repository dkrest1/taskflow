package worker

import (
	"context"
	"testing"
	"time"

	"github.com/dkrest1/taskflow/internal/model"
)

func TestProcessJobCancellation(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	job := model.NewJob(1, "Test cancellation", 5*time.Second)
	cancel()
	processJob(ctx, 1 ,job)

	actual := job.GetStatus()
	expected := "cancelled"

	if actual != expected {
    	 t.Errorf("expected %s, got %s", expected, actual)
	}
}

func TestProcessJobSuccess(t *testing.T) {
	ctx := context.Background()


	job := model.NewJob(1, "Process payment", 10 * time.Millisecond)
	processJob(ctx, 1, job)

	jobStatus := job.GetStatus()

	if jobStatus != "completed" {
		t.Errorf("expected completed, got %v", jobStatus)
	}


}
