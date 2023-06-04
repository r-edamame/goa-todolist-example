package repo

import "todo/domain/model"

type InmemoryTaskRepo struct {
	list *model.TaskList
}

func NewInmemoryTaskRepo() InmemoryTaskRepo {
	list := model.EmptyTaskList()
	return InmemoryTaskRepo{
		list: &list,
	}
}

func (repo *InmemoryTaskRepo) GetTaskList() (*model.TaskList, error) {
	return repo.list, nil
}

func (repo *InmemoryTaskRepo) SaveTaskList(list *model.TaskList) error {
	repo.list = list
	return nil
}
