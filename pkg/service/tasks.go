package service

import (
	"context"

	"github.com/Diyarjan/TasksManager/models"
	"github.com/Diyarjan/TasksManager/pkg/repository"
)

type tasks struct {
	repo *repository.Repository
}

func New(repo *repository.Repository) Tasks {
	return &tasks{
		repo: repo,
	}
}
func (t *tasks) ListTasks(ctx context.Context) ([]models.ListTask, error) {
	return t.repo.Tasks.ListTasks(ctx)
}
