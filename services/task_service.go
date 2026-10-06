package services

import (
	"errors"

	"github.com/hdmorales87/go-gorm-restapi/dto"
	"github.com/hdmorales87/go-gorm-restapi/mappers"
	"github.com/hdmorales87/go-gorm-restapi/repositories"
)

type TaskService struct {
	TaskRepository repositories.ITaskRepository
}

func NewTaskService(taskRepository repositories.ITaskRepository) *TaskService {
	return &TaskService{TaskRepository: taskRepository}
}

func (s *TaskService) GetAllTasks() ([]dto.TaskDTO, error) {
	tasks, err := s.TaskRepository.FindAll()
	if err != nil {
		return nil, err
	}
	return mappers.TasksToDTO(tasks), nil
}

func (s *TaskService) GetTaskByID(id string) (*dto.TaskDTO, error) {
	task, err := s.TaskRepository.FindByID(id)
	if err != nil {
		return nil, errors.New("Task not found")
	}
	taskDTO := mappers.TaskToDTO(*task)
	return &taskDTO, nil
}

func (s *TaskService) CreateTask(taskDTO dto.CreateTaskDTO) (*dto.TaskDTO, error) {
	task := mappers.CreateTaskDTOToModel(taskDTO)
	err := s.TaskRepository.Create(&task)
	if err != nil {
		return nil, err
	}
	createdDTO := mappers.TaskToDTO(task)
	return &createdDTO, nil
}

func (s *TaskService) UpdateTask(id string, taskDTO dto.UpdateTaskDTO) (*dto.TaskDTO, error) {
	existingTask, err := s.TaskRepository.FindByID(id)
	if err != nil {
		return nil, errors.New("Task not found")
	}

	task := mappers.UpdateTaskDTOToModel(taskDTO)
	task.ID = existingTask.ID
	err = s.TaskRepository.Update(&task)
	if err != nil {
		return nil, err
	}
	updatedDTO := mappers.TaskToDTO(task)
	return &updatedDTO, nil
}

func (s *TaskService) DeleteTask(id string) error {
	_, err := s.TaskRepository.FindByID(id)
	if err != nil {
		return errors.New("Task not found")
	}

	err = s.TaskRepository.Delete(id)
	return err
}
