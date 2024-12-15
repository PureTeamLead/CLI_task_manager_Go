package data

import "time"

type Status int

//enum for statuses
const (
	Todo Status = iota 
	In_progress 
	Done
)

type Task struct {
	ID int64 `json:"ID"`
	Description string `json:"Description"`
	Status Status `json:"Status"`
	CreatedAt time.Time `json:"CreatedAt"`
	UpdatedAt time.Time `json:"UpdatedAt"`
}

var Tasks []Task 