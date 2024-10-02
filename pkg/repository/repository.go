package repository

import (
	gorestapi "github.com/ingerstep/go-restapi"
	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateUser(user gorestapi.User) (int, error)
	GetUser(username, password string) (gorestapi.User, error)
}

type TodoList interface {
	Create(userId int, list gorestapi.TodoList) (int, error)
	GetAll(userId int) ([]gorestapi.TodoList, error)
	GetById(userId, listId int) (gorestapi.TodoList, error)
}

type TodoItem interface {
}

type Repository struct {
	Authorization
	TodoItem
	TodoList
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		TodoList:      NewTodoListPostgres(db),
	}
}
