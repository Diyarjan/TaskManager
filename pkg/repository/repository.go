package repository

type Taskslist interface {
}

type Repository struct {
	Taskslist
}

func NewRepository() *Repository {
	return &Repository{}
}
