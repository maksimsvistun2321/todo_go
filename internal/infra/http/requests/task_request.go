package requests

import (
	"time"

	"github.com/BohdanBoriak/boilerplate-go-back/internal/domain"
)

type TaskRequest struct {
	Title       string `json:"title" validate:"required,gte=1,max=50"`
	Description string `json:"description"`
	Deadline    *int64 `json:"deadline"`
}

type TaskStatusRequest struct {
	Status domain.TaskStatus `json:"status"`
}

func (r TaskRequest) ToDomainModel() (interface{}, error) {
	var timeUnix int64
	if r.Deadline != nil {
		timeUnix = *r.Deadline
	}
	var deadline *time.Time
	if timeUnix != 0 {
		dl := time.Unix(timeUnix, 0)
		deadline = &dl
	}
	return domain.Task{
		Title:       r.Title,
		Description: r.Description,
		Deadline:    deadline,
	}, nil
}

func (r TaskStatusRequest) ToDomainModel() (interface{}, error) {
	return domain.Task{
		Status: r.Status,
	}, nil
}
