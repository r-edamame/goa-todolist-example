package todoapi

import (
	"context"
	"log"
	todo "todo/gen/todo"
)

// todo service example implementation.
// The example methods log the requests and return zero values.
type todosrvc struct {
	logger *log.Logger
}

// NewTodo returns the todo service implementation.
func NewTodo(logger *log.Logger) todo.Service {
	return &todosrvc{logger}
}

// ListTasks implements listTasks.
func (s *todosrvc) ListTasks(ctx context.Context) (res []*todo.Task, err error) {
	s.logger.Print("todo.listTasks")

	tid1 := "taskid1"
	tname1 := "task1"
	incompleted := false
	task1 := todo.Task{
		ID:        &tid1,
		Name:      &tname1,
		Completed: &incompleted,
	}

	tid2 := "taskid2"
	tname2 := "task2"
	completed := true
	task2 := todo.Task{
		ID:        &tid2,
		Name:      &tname2,
		Completed: &completed,
	}

	res = []*todo.Task{
		&task1,
		&task2,
	}
	return
}

// CreateTask implements createTask.
func (s *todosrvc) CreateTask(ctx context.Context, p *todo.CreateTaskPayload) (res *todo.Task, err error) {
	id := "testid1"
	completed := false
	res = &todo.Task{
		ID:        &id,
		Name:      p.Name,
		Completed: &completed,
	}
	s.logger.Print("todo.createTask")
	return
}

// CompleteTask implements completeTask.
func (s *todosrvc) CompleteTask(ctx context.Context, p *todo.CompleteTaskPayload) (res *todo.Task, err error) {
	name := "newtask"
	completed := true
	res = &todo.Task{
		ID:        p.ID,
		Name:      &name,
		Completed: &completed,
	}
	s.logger.Print("todo.completeTask")
	return
}

// RevertTask implements revertTask.
func (s *todosrvc) RevertTask(ctx context.Context, p *todo.RevertTaskPayload) (res *todo.Task, err error) {
	name := "newtask"
	completed := false
	res = &todo.Task{
		ID:        p.ID,
		Name:      &name,
		Completed: &completed,
	}
	s.logger.Print("todo.revertTask")
	return
}
