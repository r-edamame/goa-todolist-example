package todoapi

import (
	"context"
	"log"

	"todo/domain/model"
	"todo/domain/repo"
	"todo/domain/usecases"
	todo "todo/gen/todo"
)

// todo service example implementation.
// The example methods log the requests and return zero values.
type todosrvc struct {
	logger   *log.Logger
	taskRepo repo.TaskRepo
}

// NewTodo returns the todo service implementation.
func NewTodo(logger *log.Logger) todo.Service {
	taskRepo := repo.NewInmemoryTaskRepo()
	return &todosrvc{logger, &taskRepo}
}

func convertTaskForResult(task *model.Task) *todo.Task {
	id, name, completed := task.Id(), task.Name(), task.Completed()
	return &todo.Task{
		ID:        &id,
		Name:      &name,
		Completed: &completed,
	}
}

// ListTasks implements listTasks.
func (s *todosrvc) ListTasks(ctx context.Context) (res []*todo.Task, err error) {
	s.logger.Print("todo.listTasks")

	list, err := s.taskRepo.GetTaskList()
	if err != nil {
		return
	}

	tasks := list.All()
	s.logger.Println(tasks)
	res = make([]*todo.Task, len(tasks))

	for i, task := range tasks {
		res[i] = convertTaskForResult(task)
	}

	return
}

// CreateTask implements createTask.
func (s *todosrvc) CreateTask(ctx context.Context, p *todo.CreateTaskPayload) (res *todo.Task, err error) {
	s.logger.Print("todo.createTask")

	task, err := usecases.CreateTask(s.taskRepo, *p.Name)
	if err != nil {
		return
	}

	res = convertTaskForResult(task)
	return
}

// CompleteTask implements completeTask.
func (s *todosrvc) CompleteTask(ctx context.Context, p *todo.CompleteTaskPayload) (res *todo.Task, err error) {
	s.logger.Print("todo.completeTask")

	task, err := usecases.CompleteTask(s.taskRepo, *p.ID)
	if err != nil {
		return
	}

	res = convertTaskForResult(task)
	return
}

// RevertTask implements revertTask.
func (s *todosrvc) RevertTask(ctx context.Context, p *todo.RevertTaskPayload) (res *todo.Task, err error) {
	s.logger.Print("todo.revertTask")

	task, err := usecases.RevertTask(s.taskRepo, *p.ID)
	if err != nil {
		return
	}

	res = convertTaskForResult(task)
	return
}
