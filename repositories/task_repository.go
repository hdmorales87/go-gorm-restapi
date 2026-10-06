package repositories

import (
	"errors"

	"github.com/hdmorales87/go-gorm-restapi/models"
	"gorm.io/gorm"
)

type TaskRepository struct {
	DB *gorm.DB
}

func NewTaskRepository(db *gorm.DB) *TaskRepository {
	return &TaskRepository{DB: db}
}

func (r *TaskRepository) FindAll() ([]models.Task, error) {
	var tasks []models.Task
	result := r.DB.Find(&tasks)
	if result.Error != nil {
		return nil, result.Error
	}
	return tasks, nil
}

func (r *TaskRepository) FindByID(id string) (*models.Task, error) {
	var task models.Task
	result := r.DB.First(&task, id)
	if result.Error != nil {
		return nil, errors.New("Task not found")
	}
	return &task, nil
}

func (r *TaskRepository) Create(task *models.Task) error {
	result := r.DB.Create(task)
	return result.Error
}

func (r *TaskRepository) Update(task *models.Task) error {
	result := r.DB.Save(task)
	return result.Error
}

func (r *TaskRepository) Delete(id string) error {
	var task models.Task
	result := r.DB.First(&task, id)
	if result.Error != nil {
		return errors.New("Task not found")
	}
	result = r.DB.Delete(&task, id)
	return result.Error
}
