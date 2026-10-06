package repositories

import (
	"github.com/hdmorales87/go-gorm-restapi/models"
)

type ITaskRepository interface {
	FindAll() ([]models.Task, error)
	FindByID(id string) (*models.Task, error)
	Create(task *models.Task) error
	Update(task *models.Task) error
	Delete(id string) error
}

type IUserRepository interface {
	FindAll() ([]models.User, error)
	FindByID(id string) (*models.User, error)
	Create(user *models.User) error
	Update(user *models.User) error
	Delete(id string) error
	FindWithTasks(id string) (*models.User, error)
}
