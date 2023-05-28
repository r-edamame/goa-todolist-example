package todoapi

import (
	"context"
	"log"
	domain "todo/domain"
	todo "todo/gen/todo"
)

// todo service example implementation.
// The example methods log the requests and return zero values.
type todosrvc struct {
	logger   *log.Logger
	taskRepo domain.TaskRepo
}

// NewTodo returns the todo service implementation.
func NewTodo(logger *log.Logger) todo.Service {
	return &todosrvc{logger, domain.NewInmemoryTaskRepo()}
}

// ListTasks implements listTasks.
func (s *todosrvc) ListTasks(ctx context.Context) (res []*todo.Task, err error) {
	s.logger.Print("todo.listTasks")

	list, err := s.taskRepo.ListTasks()
	if err != nil {
		return
	}

	tasks := list.All()
	res = make([]*todo.Task, len(tasks))

	for i, task := range tasks {
		res[i] = &todo.Task{
			ID:        &task.Id,
			Name:      &task.Name,
			Completed: &task.Completed,
		}
	}

	return
}

// CreateTask implements createTask.
func (s *todosrvc) CreateTask(ctx context.Context, p *todo.CreateTaskPayload) (res *todo.Task, err error) {
	s.logger.Print("todo.createTask")

	task, err := domain.CreateTask(s.taskRepo, *p.Name)
	if err != nil {
		return
	}

	res = &todo.Task{
		ID:        &task.Id,
		Name:      &task.Name,
		Completed: &task.Completed,
	}
	return
}

// CompleteTask implements completeTask.
func (s *todosrvc) CompleteTask(ctx context.Context, p *todo.CompleteTaskPayload) (res *todo.Task, err error) {
	s.logger.Print("todo.completeTask")

	task, err := domain.CompleteTask(s.taskRepo, *p.ID)
	if err != nil {
		return
	}

	res = &todo.Task{
		ID:        &task.Id,
		Name:      &task.Name,
		Completed: &task.Completed,
	}
	return
}

// RevertTask implements revertTask.
func (s *todosrvc) RevertTask(ctx context.Context, p *todo.RevertTaskPayload) (res *todo.Task, err error) {
	s.logger.Print("todo.revertTask")

	task, err := domain.RevertTask(s.taskRepo, *p.ID)
	if err != nil {
		return
	}

	res = &todo.Task{
		ID:        &task.Id,
		Name:      &task.Id,
		Completed: &task.Completed,
	}
	return
}
