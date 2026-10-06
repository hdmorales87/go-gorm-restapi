package services

import (
	"errors"

	"github.com/hdmorales87/go-gorm-restapi/dto"
	"github.com/hdmorales87/go-gorm-restapi/mappers"
	"github.com/hdmorales87/go-gorm-restapi/repositories"
)

type UserService struct {
	UserRepository repositories.IUserRepository
}

func NewUserService(userRepository repositories.IUserRepository) *UserService {
	return &UserService{UserRepository: userRepository}
}

func (s *UserService) GetAllUsers() ([]dto.UserDTO, error) {
	users, err := s.UserRepository.FindAll()
	if err != nil {
		return nil, err
	}
	return mappers.UsersToDTO(users), nil
}

func (s *UserService) GetUserByID(id string) (*dto.UserWithTasksDTO, error) {
	user, err := s.UserRepository.FindWithTasks(id)
	if err != nil {
		return nil, errors.New("User not found")
	}
	userDTO := mappers.UserWithTasksToDTO(*user)
	return &userDTO, nil
}

func (s *UserService) CreateUser(userDTO dto.CreateUserDTO) (*dto.UserDTO, error) {
	user := mappers.CreateUserDTOToModel(userDTO)
	err := s.UserRepository.Create(&user)
	if err != nil {
		return nil, err
	}
	createdDTO := mappers.UserToDTO(user)
	return &createdDTO, nil
}

func (s *UserService) UpdateUser(id string, userDTO dto.UpdateUserDTO) (*dto.UserDTO, error) {
	existingUser, err := s.UserRepository.FindByID(id)
	if err != nil {
		return nil, errors.New("User not found")
	}

	user := mappers.UpdateUserDTOToModel(userDTO)
	user.ID = existingUser.ID
	err = s.UserRepository.Update(&user)
	if err != nil {
		return nil, err
	}
	updatedDTO := mappers.UserToDTO(user)
	return &updatedDTO, nil
}

func (s *UserService) DeleteUser(id string) error {
	_, err := s.UserRepository.FindByID(id)
	if err != nil {
		return errors.New("User not found")
	}

	err = s.UserRepository.Delete(id)
	return err
}
