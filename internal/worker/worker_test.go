package worker

import (
	"context"
	"testing"
	"time"

	"github.com/dkrest1/taskflow/internal/model"
)


func TestJobProcessCancelation(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
    defer cancel()

	job := model.NewJob(1, "Test cancellation", 5*time.Second)
}