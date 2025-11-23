package repository

import (
	"pharmacy-team/internal/models"

	"gorm.io/gorm"
)

type UserRepository interface {
	Create(user *models.User) error
	GetByID(id uint) (*models.User, error)
	GetAll() ([]models.User, error)
	Update(*models.User) error
	Delete(id uint) error
}

type gormUserRepository struct {
	db *gorm.DB // подключение к базе данных через GORM
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &gormUserRepository{db: db}
}

func (r *gormUserRepository) Create(user *models.User) error {
	if user == nil {
		return nil
	}

	return r.db.Create(user).Error

}

func (r *gormUserRepository) GetByID(id uint) (*models.User, error) {
	var user models.User

	if err := r.db.First(&user, id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *gormUserRepository) GetAll() ([]models.User, error) {
	var user []models.User

	if err := r.db.Find(&user).Error; err != nil {
		return nil, err
	}

	return user, nil
}

func (r *gormUserRepository) Update(pharmacy *models.User) error {
	if pharmacy == nil {
		return nil
	}
	return r.db.Save(pharmacy).Error
}

func (r *gormUserRepository) Delete(id uint) error {
	return r.db.Delete(&models.User{}, id).Error
}
