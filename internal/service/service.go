package service

import (
	"ToDoRestApi/internal/domain"
	"ToDoRestApi/internal/repository"
)

type Authorization interface {
	CreateUser(user domain.User) (int, error)
}

type TodoList interface {
	
}

type TodoItem interface {
	
}

type Service struct {
	Authorization
	TodoList
	TodoItem
}

func NewService(rep *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(rep.Authorization),
	}
}