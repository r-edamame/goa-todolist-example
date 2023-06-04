package model

import (
	"encoding/json"
	"errors"

	"github.com/google/uuid"
)

type Task struct {
	id        string
	name      string
	completed bool
}

func (task Task) MarshalJSON() ([]byte, error) {
	m := map[string]interface{}{
		"id":        task.id,
		"name":      task.name,
		"completed": task.completed,
	}
	return json.Marshal(m)
}

func NewTask(name string) Task {
	id, _ := uuid.NewUUID()
	return Task{
		id:        id.String(),
		name:      name,
		completed: false,
	}
}

// getters
func (task *Task) Id() string {
	return task.id
}
func (task *Task) Name() string {
	return task.name
}
func (task *Task) Completed() bool {
	return task.completed
}

// methods
func (task *Task) Complete() error {
	if task.completed {
		return errors.New("task is already completed")
	}
	task.completed = true
	return nil
}
func (task *Task) Revert() error {
	if !task.completed {
		return errors.New("task is not completed")
	}
	task.completed = false
	return nil
}

type TaskList struct {
	list []*Task
}

func EmptyTaskList() TaskList {
	return TaskList{
		list: make([]*Task, 0),
	}
}

func (list *TaskList) All() []*Task {
	return list.list
}
func (list *TaskList) Append(task Task) {
	list.list = append(list.list, &task)
}
func (list *TaskList) Find(id string) (*Task, bool) {
	for _, task := range list.All() {
		if task.Id() == id {
			return task, true
		}
	}
	return nil, false
}
