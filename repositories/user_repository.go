package repositories

import (
	"errors"

	"github.com/hdmorales87/go-gorm-restapi/models"
	"gorm.io/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{DB: db}
}

func (r *UserRepository) FindAll() ([]models.User, error) {
	var users []models.User
	result := r.DB.Find(&users)
	if result.Error != nil {
		return nil, result.Error
	}
	return users, nil
}

func (r *UserRepository) FindByID(id string) (*models.User, error) {
	var user models.User
	result := r.DB.First(&user, id)
	if result.Error != nil {
		return nil, errors.New("User not found")
	}
	return &user, nil
}

func (r *UserRepository) FindWithTasks(id string) (*models.User, error) {
	var user models.User
	result := r.DB.First(&user, id)
	if result.Error != nil {
		return nil, errors.New("User not found")
	}
	r.DB.Model(&user).Association("Tasks").Find(&user.Tasks)
	return &user, nil
}

func (r *UserRepository) Create(user *models.User) error {
	result := r.DB.Create(user)
	return result.Error
}

func (r *UserRepository) Update(user *models.User) error {
	result := r.DB.Save(user)
	return result.Error
}

func (r *UserRepository) Delete(id string) error {
	var user models.User
	result := r.DB.First(&user, id)
	if result.Error != nil {
		return errors.New("User not found")
	}
	result = r.DB.Delete(&user, id)
	return result.Error
}
