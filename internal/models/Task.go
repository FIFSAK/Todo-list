package models

import (
	"database/sql"
	"log"
	"time"
)

type Task struct {
	ID       int       `json:"id"`
	Status   string    `json:"status"`
	Title    string    `json:"title"`
	ActiveAt time.Time `json:"activeAt"`
}
type TaskModel struct {
	DB *sql.DB
}

func NewTask(id int, status string, title string, activeAt string) *Task {
	layout := "2006-01-02"
	activeAtTime, err := time.Parse(layout, activeAt)
	if err != nil {
		log.Printf("Error parsing date: %v", err)
		return nil
	}
	return &Task{
		ID:       id,
		Status:   status,
		Title:    title,
		ActiveAt: activeAtTime,
	}
}

func CreateTask(t *TaskModel) (int, error) {

	return 0, nil
}
