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
}

type TodoItem interface {
}

type Service struct {
	Authorization
	TodoItem
	TodoList
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
	}
}
