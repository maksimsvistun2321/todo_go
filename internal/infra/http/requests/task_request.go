package requests

import (
	"time"

	"github.com/BohdanBoriak/boilerplate-go-back/internal/domain"
)

type CreateTaskRequest struct {
	Title       string            `json:"title" validate:"required,gte=1,max=50"`
	Description string            `json:"description" validate:"required,gte=1,max=255"`
	Status      domain.TaskStatus `json:"status" validate:"required"`
	Deadline    *time.Time        `json:"deadline" validate:"required"`
}

type UpdateTaskRequest struct {
	Title       string            `json:"title" validate:"required,gte=1,max=50"`
	Description string            `json:"description" validate:"required,gte=1,max=255"`
	Status      domain.TaskStatus `json:"status" validate:"required"`
}

func (r CreateTaskRequest) ToDomainModel() (interface{}, error) {
	return domain.Task{
		Title:       r.Title,
		Description: r.Description,
		Status:      r.Status,
		Deadline:    r.Deadline,
	}, nil
}

func (r UpdateTaskRequest) ToDomainModel() (interface{}, error) {
	return domain.Task{
		Title:       r.Title,
		Description: r.Description,
		Status:      r.Status,
	}, nil
}
