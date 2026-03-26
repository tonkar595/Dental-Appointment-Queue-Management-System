package service

import (
	models "backend/internal/model"
	"backend/internal/repository"
)

type UserService struct {
	repo *repository.UserRepository
}

func NewUserService(repo *repository.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) CreateUser(user *models.User) error {
	return s.repo.Create(user)

}
func (s *UserService) GetAllUsers() ([]models.User, error) {
	return s.repo.GetAll()
}
func (s *UserService) GetUserByID(id int) (*models.User, error) {
	return s.repo.GetByID(id)
}
func (s *UserService) UpdateUser(id int, data *models.User) error {
	return s.repo.Update(id, data)
}
func (s *UserService) DeleteUser(id int) error {
	return s.repo.Delete(id)
}
