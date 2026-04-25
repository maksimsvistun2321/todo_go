package controllers

import (
	"log"
	"net/http"

	"github.com/BohdanBoriak/boilerplate-go-back/internal/app"
	"github.com/BohdanBoriak/boilerplate-go-back/internal/domain"
	"github.com/BohdanBoriak/boilerplate-go-back/internal/infra/http/requests"
	"github.com/BohdanBoriak/boilerplate-go-back/internal/infra/http/resources"
)

type TaskController struct {
	taskService app.TaskService
}

func NewTaskController(us app.TaskService) TaskController {
	return TaskController{
		taskService: us,
	}
}

func (c TaskController) FindTask() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		task := r.Context().Value(TaskKey).(domain.Task)
		Success(w, resources.TaskDto{}.DomainToDto(task))
	}
}

func (c TaskController) Update() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		task, err := requests.Bind(r, requests.UpdateTaskRequest{}, domain.Task{})
		if err != nil {
			log.Printf("TaskController: %s", err)
			BadRequest(w, err)
			return
		}

		t := r.Context().Value(TaskKey).(domain.Task)
		t.Title = task.Title
		t.Description = task.Description
		t.Status = task.Status
		task, err = c.taskService.Update(t)
		if err != nil {
			log.Printf("TaskController: %s", err)
			InternalServerError(w, err)
			return
		}

		var taskDto resources.TaskDto
		Success(w, taskDto.DomainToDto(task))
	}
}

func (c TaskController) Delete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		t := r.Context().Value(TaskKey).(domain.Task)

		err := c.taskService.Delete(t.Id)
		if err != nil {
			log.Printf("TaskController: %s", err)
			InternalServerError(w, err)
			return
		}

		Ok(w)
	}
}
