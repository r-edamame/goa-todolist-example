package domain

import (
	"errors"

	"github.com/google/uuid"
)

// ========== Types

// Task
type Task struct {
	Id        string
	Name      string
	Completed bool
}

func newTask(name string) Task {
	id, _ := uuid.NewUUID()

	return Task{
		Id:        id.String(),
		Name:      name,
		Completed: false,
	}
}

// TaskList
type TaskList struct {
	list []*Task
}

func (list *TaskList) Append(task Task) {
	(*list).list = append(list.list, &task)
}
func (list *TaskList) All() []*Task {
	return list.list
}
func (list *TaskList) Find(id string) (*Task, bool) {
	var i int
	for ; list.list[i].Id != id; i++ {
	}

	if i == len(list.list) {
		return nil, false
	}

	return list.list[i], true
}

// ========== Repo

type TaskRepo interface {
	ListTasks() (TaskList, error)
	SaveTasks(list TaskList) error
}

// ========= Usecases

func CreateTask(repo TaskRepo, name string) (res *Task, err error) {
	list, err := repo.ListTasks()
	if err != nil {
		return
	}

	task := newTask(name)
	list.Append(task)
	err = repo.SaveTasks(list)
	if err != nil {
		return
	}

	return &task, nil
}

func CompleteTask(repo TaskRepo, id string) (res *Task, err error) {
	list, err := repo.ListTasks()
	if err != nil {
		return
	}

	task, ok := list.Find(id)
	if !ok {
		return nil, errors.New("task not found")
	}

	task.Completed = true
	repo.SaveTasks(list)

	return task, nil
}

func RevertTask(repo TaskRepo, id string) (res *Task, err error) {
	list, err := repo.ListTasks()
	if err != nil {
		return
	}

	task, ok := list.Find(id)
	if !ok {
		return nil, errors.New("task not found")
	}

	task.Completed = false
	repo.SaveTasks(list)

	return task, nil
}

// ========= repo impl

type InmemoryTaskRepo struct {
	tasklist *TaskList
}

func NewInmemoryTaskRepo() InmemoryTaskRepo {
	return InmemoryTaskRepo{
		tasklist: &TaskList{},
	}
}
func (repo InmemoryTaskRepo) ListTasks() (TaskList, error) {
	return *repo.tasklist, nil
}
func (repo InmemoryTaskRepo) SaveTasks(list TaskList) error {
	*(repo.tasklist) = list
	return nil
}
