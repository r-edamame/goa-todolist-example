package usecases

import (
	"errors"
	"todo/domain/model"
	"todo/domain/repo"
)

func CreateTask(repo repo.TaskRepo, name string) (res *model.Task, err error) {
	list, err := repo.GetTaskList()
	if err != nil {
		return
	}

	task := model.NewTask(name)
	list.Append(task)
	err = repo.SaveTaskList(list)
	if err != nil {
		return
	}

	return &task, nil
}

func CompleteTask(repo repo.TaskRepo, id string) (res *model.Task, err error) {
	list, err := repo.GetTaskList()
	if err != nil {
		return
	}

	task, ok := list.Find(id)
	if !ok {
		return nil, errors.New("task not found")
	}

	task.Complete()
	err = repo.SaveTaskList(list)
	if err != nil {
		return
	}

	return task, nil
}

func RevertTask(repo repo.TaskRepo, id string) (res *model.Task, err error) {
	list, err := repo.GetTaskList()
	if err != nil {
		return
	}

	task, ok := list.Find(id)
	if !ok {
		return nil, errors.New("task not found")
	}

	task.Revert()
	err = repo.SaveTaskList(list)
	if err != nil {
		return
	}

	return task, nil
}
