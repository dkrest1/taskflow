package model

import "time"

type Job struct {
	ID       int           `json:"id"`
	Payload  string        `json:"payload"`
	Status   string        `json:"status"`
	Duration time.Duration `json:"duration"`
}

func NewJob(id int, payload string, duration time.Duration) *Job {
	return &Job{
		ID:       id,
		Payload:  payload,
		Status:   "pending",
		Duration: duration,
	}
}
