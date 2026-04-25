package resources

import (
	"time"

	"github.com/BohdanBoriak/boilerplate-go-back/internal/domain"
)

type TaskDto struct {
	Id          uint64            `json:"id"`
	Title       string            `json:"title"`
	Description string            `json:"description"`
	Status      domain.TaskStatus `json:"status"`
	Deadline    *time.Time        `json:"deadline"`
}

type TasksDto struct {
	Items []TasksDto `json:"items"`
	Total uint64     `json:"total"`
	Pages uint       `json:"pages"`
}

func (d TaskDto) DomainToDto(task domain.Task) TaskDto {
	return TaskDto{
		Id:          task.Id,
		Title:       task.Title,
		Description: task.Description,
		Status:      task.Status,
		Deadline:    task.Deadline,
	}
}

func (d TaskDto) DomainToDtoCollection(tasks []domain.Task) []TaskDto {
	result := make([]TaskDto, len(tasks))
	for i, u := range tasks {
		result[i] = d.DomainToDto(u)
	}
	return result
}
