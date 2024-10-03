package service

import (
	gorestapi "github.com/ingerstep/go-restapi"
	"github.com/ingerstep/go-restapi/pkg/repository"
)

type Authorization interface {
	CreateUser(user gorestapi.User) (int, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(token string) (int, error)
}

type TodoList interface {
	Create(userId int, list gorestapi.TodoList) (int, error)
	GetAll(userId int) ([]gorestapi.TodoList, error)
	GetById(userId, listId int) (gorestapi.TodoList, error)
	Delete(userId, listId int) error
	Update(userId, listId int, input gorestapi.UpdateListInput) error
}

type TodoItem interface {
	Create(userId, listId int, item gorestapi.TodoItem) (int, error)
	GetAll(userId, listId int) ([]gorestapi.TodoItem, error)
	GetById(userId, itemId int) (gorestapi.TodoItem, error)
	Delete(userId, itemId int) error
	Update(userId, itemId int, input gorestapi.UpdateItemInput) error
}

type Service struct {
	Authorization
	TodoItem
	TodoList
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
		TodoList:      NewTodoListService(repos.TodoList),
		TodoItem:      NewTodoItemService(repos.TodoItem, repos.TodoList),
	}
}
