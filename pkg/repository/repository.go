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
	}
}
