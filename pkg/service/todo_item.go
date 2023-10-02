package service

import (
	"github.com/DimqBBSQO/todo-app"
	"github.com/DimqBBSQO/todo-app/pkg/repository"
)

type TodoItemService struct {
	repo     repository.TodoItem
	listRepo repository.TodoList
}

func NewTodoItemService(repo repository.TodoItem, listRepo repository.TodoList) *TodoItemService {
	return &TodoItemService{repo: repo, listRepo: listRepo}
}

func (i *TodoItemService) Create(userId, listId int, item todo.TodoItem) (int, error) {
	_, err := i.listRepo.GetListById(userId, listId)
	if err != nil {
		return 0, err
	}
	return i.repo.Create(listId, item)
}

func (i *TodoItemService) GetAll(userId, listId int) ([]todo.TodoItem, error) {
	return i.repo.GetAll(userId, listId)
}

func (i *TodoItemService) GetById(userId, itemId int) (todo.TodoItem, error) {
	return i.repo.GetById(userId, itemId)
}

func (i *TodoItemService) Delete(userId, itemId int) error {
	return i.repo.Delete(userId, itemId)
}

func (i *TodoItemService) Update(userId, listId int, input todo.UpdateItemInput) error {
	return i.repo.Update(userId, listId, input)
}
