package services

import (
	"github.com/hdmorales87/go-gorm-restapi/dto"
)

type ITaskService interface {
	GetAllTasks() ([]dto.TaskDTO, error)
	GetTaskByID(id string) (*dto.TaskDTO, error)
	CreateTask(taskDTO dto.CreateTaskDTO) (*dto.TaskDTO, error)
	UpdateTask(id string, taskDTO dto.UpdateTaskDTO) (*dto.TaskDTO, error)
	DeleteTask(id string) error
}

type IUserService interface {
	GetAllUsers() ([]dto.UserDTO, error)
	GetUserByID(id string) (*dto.UserWithTasksDTO, error)
	CreateUser(userDTO dto.CreateUserDTO) (*dto.UserDTO, error)
	UpdateUser(id string, userDTO dto.UpdateUserDTO) (*dto.UserDTO, error)
	DeleteUser(id string) error
}
