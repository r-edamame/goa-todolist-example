package repo

import "todo/domain/model"

type TaskRepo interface {
	GetTaskList() (*model.TaskList, error)
	SaveTaskList(list *model.TaskList) error
}
