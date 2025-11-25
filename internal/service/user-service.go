package service

import (
	"errors"
	"pharmacy-team/internal/models"
	"pharmacy-team/internal/repository"
)

var ErrUserNotFound = errors.New("студент не найден")

type UserService interface {
	CreateUser(req models.UserCreate) (*models.User, error)
	GetByUser(id uint) (*models.User, error)
	UpdateUser(id uint, req models.UserUpdate) (*models.User, error)
	DeleteUser(id uint) error
	GetUsers() ([]models.User, error)
}
type userService struct {
	repository repository.UserRepository
}

func NewUserService(reposit repository.UserRepository) UserService {
	return &userService{repository: reposit}
}

func validateUser(req models.UserCreate) error {
	if req.FullName == "" {
		return errors.New("ФИО не должно быть пустым")
	}
	if req.Email == "" {
		return errors.New("email не должно быть пустым")
	}
	if req.DefaultAddress == "" {
		return errors.New("адрес не должен быть пустым")
	}
	return nil
}

func (s *userService) CreateUser(req models.UserCreate) (*models.User, error) {

	if err := validateUser(req); err != nil {
		return nil, err
	}

	user := &models.User{
		FullName:       req.FullName,
		Email:          req.Email,
		Phone:          req.Phone,
		DefaultAddress: req.DefaultAddress,
	}

	if err := s.repository.Create(user); err != nil {
		return nil, err
	}
	return user, nil
}

func (s *userService) GetByUser(id uint) (*models.User, error) {
	users, err := s.repository.GetByID(id)
	if err != nil {
		return nil, errors.New("пользователь не найден")
	}

	return users, err
}

func (s *userService) UpdateUser(id uint, req models.UserUpdate) (*models.User, error) {
	user, err := s.repository.GetByID(id)
	if err != nil {
		return nil, errors.New("пользователь не найден")
	}

	if req.FullName != nil {
		user.FullName = *req.FullName
	}
	if req.Email != nil {
		user.Email = *req.Email
	}
	if req.Phone != nil {
		user.Phone = *req.Phone
	}
	if req.DefaultAddress != nil {
		user.DefaultAddress = *req.DefaultAddress
	}

	if err = s.repository.Update(user); err != nil {
		return nil, errors.New("ошибка при попытки изминения пользователя")
	}

	return user, nil
}

func (s *userService) DeleteUser(id uint) error {

	if _, err := s.repository.GetByID(id); err != nil {
		return errors.New("пользователь не найден")
	}

	if err := s.repository.Delete(id); err != nil {
		return errors.New("ошибка при удалении пользователя")
	}

	return nil
}

func (s *userService) GetUsers() ([]models.User, error) {

	users, err := s.repository.GetAll()
	if err != nil {
		return nil, errors.New("не удалось загрузить список пользователей")
	}

	return users, nil
}
