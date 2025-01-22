package service

import (
	"context"

	"github.com/Diyarjan/TasksManager/models"
	"github.com/Diyarjan/TasksManager/pkg/repository"
)

type Tasks interface {
	ListTasks(ctx context.Context) ([]models.ListTask, error)
}

type Service struct {
	Tasks
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Tasks: New(repos),
	}
}
