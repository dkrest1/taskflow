package model


type Job struct {
	ID int `json:"id"`
	Payload string `json:"payload"`
	Status string `json:"status"`
}

func NewJob(id int, payload string) *Job {
	return &Job{
		ID: id,
		Payload: payload,
		Status: "pending",
	}
}