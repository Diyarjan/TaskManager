package repository

import (
	"context"

	"github.com/Diyarjan/TasksManager/models"
	"github.com/jmoiron/sqlx"
)

type tasks struct {
	db *sqlx.DB
}

func New(db *sqlx.DB) Tasks {
	return &tasks{
		db: db,
	}
}

func (r *tasks) ListTasks(ctx context.Context) ([]models.ListTask, error) {
	// TODO Database interaction
	return nil, nil
}
