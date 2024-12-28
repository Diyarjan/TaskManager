package service

import "github.com/Diyarjan/TasksManager/pkg/repository"

type Taskslist interface {
}

type Service struct {
	Taskslist
}

func NewService(repos *repository.Repository) *Service {
	return &Service{}
}
