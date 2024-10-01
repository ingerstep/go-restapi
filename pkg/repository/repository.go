package repository

import (
	gorestapi "github.com/ingerstep/go-restapi"
	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateUser(gorestapi.User) (int, error)
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
