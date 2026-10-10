package model

import (
	"sync"
	"time"
)

type Job struct {
	ID       int           `json:"id"`
	Payload  string        `json:"payload"`
	Status   string        `json:"status"`
	Duration time.Duration `json:"duration"`

	mu sync.Mutex
}

func NewJob(id int, payload string, duration time.Duration) *Job {
	return &Job{
		ID:       id,
		Payload:  payload,
		Status:   "pending",
		Duration: duration,
	}
}

func (j *Job) SetStatus(status string) {
	j.mu.Lock()
	defer j.mu.Unlock()

	j.Status = status
}

func (j *Job) GetStatus() string {
	j.mu.Lock()
	defer j.mu.Unlock()

	return j.Status
}

