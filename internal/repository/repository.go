package repository

import (
	"ToDoRestApi/internal/domain"
	"github.com/jmoiron/sqlx"
)

const (
	userTable = "users"
)

type Authorization interface {
	CreateUser(user domain.User) (int, error)
}

type TodoList interface {

}

type TodoItem interface {

}

type Repository struct {
	Authorization
	TodoList
	TodoItem
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
	}
}
