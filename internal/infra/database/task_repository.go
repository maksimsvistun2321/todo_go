package database

import (
	"time"

	"github.com/BohdanBoriak/boilerplate-go-back/internal/domain"
	"github.com/upper/db/v4"
)

const TasksTableName = "tasks"

type task struct {
	Id          uint64            `db:"id,omitempty"`
	UserId      uint64            `db:"user_id"`
	Title       string            `db:"title"`
	Description string            `db:"description"`
	Status      domain.TaskStatus `db:"status"`
	Deadline    *time.Time        `db:"deadline"`
	CreatedDate time.Time         `db:"created_date,omitempty"`
	UpdatedDate time.Time         `db:"updated_date,omitempty"`
	DeletedDate *time.Time        `db:"deleted_date,omitempty"`
}

type TaskRepository interface {
	FindByTitle(title string) (domain.Task, error)
	FindById(id uint64) (domain.Task, error)
	Find(id uint64) (interface{}, error)
	Update(task domain.Task) (domain.Task, error)
	Delete(id uint64) error
}

type taskRepository struct {
	coll db.Collection
	sess db.Session
}

func NewTaskRepository(dbSession db.Session) TaskRepository {
	return taskRepository{
		coll: dbSession.Collection(TasksTableName),
		sess: dbSession,
	}
}

func (r taskRepository) FindByTitle(title string) (domain.Task, error) {
	var t task
	err := r.coll.Find(db.Cond{"title": title, "deleted_date": nil}).One(&t)
	if err != nil {
		return domain.Task{}, err
	}
	return r.mapModelToDomain(t), nil
}

func (r taskRepository) FindById(id uint64) (domain.Task, error) {
	var tsk task
	err := r.coll.Find(db.Cond{"id": id}).One(&tsk)
	if err != nil {
		return domain.Task{}, err
	}

	return r.mapModelToDomain(tsk), nil
}

func (r taskRepository) Find(id uint64) (interface{}, error) {
	var tsk task
	err := r.coll.Find(db.Cond{"id": id}).One(&tsk)
	if err != nil {
		return domain.Task{}, err
	}

	return r.mapModelToDomain(tsk), nil
}

func (r taskRepository) Save(tasks domain.Task) (domain.Task, error) {
	t := r.mapDomainToModel(tasks)
	t.CreatedDate, t.UpdatedDate = time.Now(), time.Now()
	err := r.coll.InsertReturning(&t)
	if err != nil {
		return domain.Task{}, err
	}
	return r.mapModelToDomain(t), nil
}

func (r taskRepository) Update(tasks domain.Task) (domain.Task, error) {
	t := r.mapDomainToModel(tasks)
	t.UpdatedDate = time.Now()
	err := r.coll.Find(db.Cond{"id": t.Id, "deleted_date": nil}).Update(&t)
	if err != nil {
		return domain.Task{}, err
	}
	return r.mapModelToDomain(t), nil
}

func (r taskRepository) Delete(id uint64) error {
	return r.coll.Find(db.Cond{"id": id, "deleted_date": nil}).Update(map[string]interface{}{"deleted_date": time.Now()})
}

func (r taskRepository) mapDomainToModel(d domain.Task) task {
	return task{
		Id:          d.Id,
		Title:       d.Title,
		Description: d.Description,
		Status:      d.Status,
		Deadline:    d.Deadline,
		CreatedDate: d.CreatedDate,
		UpdatedDate: d.UpdatedDate,
		DeletedDate: d.DeletedDate,
	}
}

func (r taskRepository) mapModelToDomain(m task) domain.Task {
	return domain.Task{
		Id:          m.Id,
		Title:       m.Title,
		Description: m.Description,
		Status:      m.Status,
		Deadline:    m.Deadline,
		CreatedDate: m.CreatedDate,
		UpdatedDate: m.UpdatedDate,
		DeletedDate: m.DeletedDate,
	}
}
