package services

import (
	"errors"

	"github.com/hdmorales87/go-gorm-restapi/dto"
	"github.com/hdmorales87/go-gorm-restapi/mappers"
	"github.com/hdmorales87/go-gorm-restapi/models"
	"gorm.io/gorm"
)

type UserService struct {
	DB *gorm.DB
}

func NewUserService(db *gorm.DB) *UserService {
	return &UserService{DB: db}
}

func (s *UserService) GetAllUsers() ([]dto.UserDTO, error) {
	var users []models.User
	result := s.DB.Find(&users)
	if result.Error != nil {
		return nil, result.Error
	}
	return mappers.UsersToDTO(users), nil
}

func (s *UserService) GetUserByID(id string) (*dto.UserWithTasksDTO, error) {
	var user models.User
	result := s.DB.First(&user, id)
	if result.Error != nil {
		return nil, errors.New("User not found")
	}
	s.DB.Model(&user).Association("Tasks").Find(&user.Tasks)
	userDTO := mappers.UserWithTasksToDTO(user)
	return &userDTO, nil
}

func (s *UserService) CreateUser(userDTO dto.CreateUserDTO) (*dto.UserDTO, error) {
	user := mappers.CreateUserDTOToModel(userDTO)
	result := s.DB.Create(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	createdDTO := mappers.UserToDTO(user)
	return &createdDTO, nil
}

func (s *UserService) UpdateUser(id string, userDTO dto.UpdateUserDTO) (*dto.UserDTO, error) {
	var existingUser models.User
	result := s.DB.First(&existingUser, id)
	if result.Error != nil {
		return nil, errors.New("User not found")
	}

	user := mappers.UpdateUserDTOToModel(userDTO)
	user.ID = existingUser.ID
	s.DB.Save(&user)
	updatedDTO := mappers.UserToDTO(user)
	return &updatedDTO, nil
}

func (s *UserService) DeleteUser(id string) error {
	var user models.User
	result := s.DB.First(&user, id)
	if result.Error != nil {
		return errors.New("User not found")
	}

	s.DB.Delete(&user, id)
	return nil
}
