package services

import (
	"errors"

	"github.com/hdmorales87/go-gorm-restapi/dto"
	"github.com/hdmorales87/go-gorm-restapi/mappers"
	"github.com/hdmorales87/go-gorm-restapi/models"
	"gorm.io/gorm"
)

type TaskService struct {
	DB *gorm.DB
}

func NewTaskService(db *gorm.DB) *TaskService {
	return &TaskService{DB: db}
}

func (s *TaskService) GetAllTasks() ([]dto.TaskDTO, error) {
	var tasks []models.Task
	result := s.DB.Find(&tasks)
	if result.Error != nil {
		return nil, result.Error
	}
	return mappers.TasksToDTO(tasks), nil
}

func (s *TaskService) GetTaskByID(id string) (*dto.TaskDTO, error) {
	var task models.Task
	result := s.DB.First(&task, id)
	if result.Error != nil {
		return nil, errors.New("Task not found")
	}
	taskDTO := mappers.TaskToDTO(task)
	return &taskDTO, nil
}

func (s *TaskService) CreateTask(taskDTO dto.CreateTaskDTO) (*dto.TaskDTO, error) {
	task := mappers.CreateTaskDTOToModel(taskDTO)
	result := s.DB.Create(&task)
	if result.Error != nil {
		return nil, result.Error
	}
	createdDTO := mappers.TaskToDTO(task)
	return &createdDTO, nil
}

func (s *TaskService) UpdateTask(id string, taskDTO dto.UpdateTaskDTO) (*dto.TaskDTO, error) {
	var existingTask models.Task
	result := s.DB.First(&existingTask, id)
	if result.Error != nil {
		return nil, errors.New("Task not found")
	}

	task := mappers.UpdateTaskDTOToModel(taskDTO)
	task.ID = existingTask.ID
	s.DB.Save(&task)
	updatedDTO := mappers.TaskToDTO(task)
	return &updatedDTO, nil
}

func (s *TaskService) DeleteTask(id string) error {
	var task models.Task
	result := s.DB.First(&task, id)
	if result.Error != nil {
		return errors.New("Task not found")
	}

	s.DB.Delete(&task, id)
	return nil
}
