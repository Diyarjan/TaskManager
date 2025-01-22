package repository

import (
	"context"

	"github.com/Diyarjan/TasksManager/models"
	"github.com/jmoiron/sqlx"
)

type Tasks interface {
	ListTasks(ctx context.Context) ([]models.ListTask, error)
}

type Repository struct {
	Tasks
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{}
}
