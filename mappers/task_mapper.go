package mappers

import (
	"github.com/hdmorales87/go-gorm-restapi/dto"
	"github.com/hdmorales87/go-gorm-restapi/models"
)

func TaskToDTO(task models.Task) dto.TaskDTO {
	return dto.TaskDTO{
		ID:          task.ID,
		Title:       task.Title,
		Description: task.Description,
		Done:        task.Done,
		UserID:      task.UserID,
	}
}

func TasksToDTO(tasks []models.Task) []dto.TaskDTO {
	dtos := make([]dto.TaskDTO, len(tasks))
	for i, task := range tasks {
		dtos[i] = TaskToDTO(task)
	}
	return dtos
}

func CreateTaskDTOToModel(dto dto.CreateTaskDTO) models.Task {
	done := dto.Done
	return models.Task{
		Title:       dto.Title,
		Description: dto.Description,
		Done:        done,
		UserID:      dto.UserID,
	}
}

func UpdateTaskDTOToModel(dto dto.UpdateTaskDTO) models.Task {
	done := false
	if dto.Done != nil {
		done = *dto.Done
	}
	return models.Task{
		Title:       dto.Title,
		Description: dto.Description,
		Done:        done,
		UserID:      dto.UserID,
	}
}
