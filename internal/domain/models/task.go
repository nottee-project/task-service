package models

import (
	"time"

	"github.com/google/uuid"
)

type Task struct {
	Id        string    `db:"id" json:"id"`
	UserId    uuid.UUID `db:"user_id" json:"user_id"`
	Title     string    `db:"title" json:"title"`
	Body      string    `db:"body" json:"body"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
	UpdatedAt time.Time `db:"updated_at" json:"updated_at"`
}

type CreateTaskRequest struct {
	UserId uuid.UUID `json:"user_id"`
	Title  string    `json:"title"`
	Body   string    `json:"body"`
}

type CreateTaskResponse struct {
	Id        string    `db:"id" json:"id"`
	UserId    uuid.UUID `db:"user_id" json:"user_id"`
	Title     string    `db:"title" json:"title"`
	Body      string    `db:"body" json:"body"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
	UpdatedAt time.Time `db:"updated_at" json:"updated_at"`
}

type Tasks struct {
	Tasks []Task
}

type ListTasksRequest struct {
	UserId uuid.UUID `json:"user_id"`
}

type ListTasksResponse struct {
	Tasks Tasks
}

type GetTaskRequest struct {
	TaskId string
	UserId string
}

type GetTaskResponse struct {
	ID     string    `db:"id"`
	UserID uuid.UUID `db:"user_id"`
	Title  string    `db:"title"`
	Body   string    `db:"body"`
}

type UpdateTaskRequest struct {
	Id     string
	UserId uuid.UUID
	// Title     string
	// Body      string
	// UpdatedAt time.Time
}

type UpdateTaskResponse struct {
	Id        string
	UserId    uuid.UUID
	Title     string
	Body      string
	UpdatedAt time.Time
}

type DeleteTaskRequest struct {
	TaskId string
	UserId string
}

type DeleteTaskResponse struct {
	Message string
}
