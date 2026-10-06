package mappers

import (
	"github.com/hdmorales87/go-gorm-restapi/dto"
	"github.com/hdmorales87/go-gorm-restapi/models"
)

func UserToDTO(user models.User) dto.UserDTO {
	return dto.UserDTO{
		ID:        user.ID,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
	}
}

func UsersToDTO(users []models.User) []dto.UserDTO {
	dtos := make([]dto.UserDTO, len(users))
	for i, user := range users {
		dtos[i] = UserToDTO(user)
	}
	return dtos
}

func UserWithTasksToDTO(user models.User) dto.UserWithTasksDTO {
	taskDTOs := TasksToDTO(user.Tasks)
	return dto.UserWithTasksDTO{
		ID:        user.ID,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
		Tasks:     taskDTOs,
	}
}

func CreateUserDTOToModel(dto dto.CreateUserDTO) models.User {
	return models.User{
		FirstName: dto.FirstName,
		LastName:  dto.LastName,
		Email:     dto.Email,
	}
}

func UpdateUserDTOToModel(dto dto.UpdateUserDTO) models.User {
	return models.User{
		FirstName: dto.FirstName,
		LastName:  dto.LastName,
		Email:     dto.Email,
	}
}
