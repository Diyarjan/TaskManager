package repository

import "github.com/jmoiron/sqlx"

type Taskslist interface {
}

type Repository struct {
	Taskslist
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{}
}
