package service

import (
	"fitness/internal/auth"
	"fitness/internal/model"
	"fitness/internal/repository"

	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	CreateUser(user *model.User) error
	GetUserByID(id int) (*model.User, error)
	GetUserByEmail(email string) (*model.User, error)
	UpdateUser(user *model.User) error
	DeleteUser(id int) error
}

type userService struct {
	repo repository.UserRepository
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func NewUserService(repo repository.UserRepository) UserService {
	return &userService{repo: repo}
}

func (s *userService) CreateUser(user *model.User) error {
	hashedPassword, err := auth.HashPassword(user.Password)
	if err != nil {
		return err
	}
	user.Password = hashedPassword
	return s.repo.Create(user)
}

func (s *userService) GetUserByID(id int) (*model.User, error) {
	return s.repo.GetByID(id)
}

func (s *userService) GetUserByEmail(email string) (*model.User, error) {
	return s.repo.GetByEmail(email)
}

func (s *userService) UpdateUser(user *model.User) error {
	return s.repo.Update(user)
}

func (s *userService) DeleteUser(id int) error {
	return s.repo.Delete(id)
}
