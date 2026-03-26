package repository

import (
	models "backend/internal/model"

	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db}
}

func (r *UserRepository) Create(user *models.User) error {
	return r.db.Create(user).Error

}
func (r *UserRepository) GetAll() ([]models.User, error) {
	var users []models.User
	err := r.db.Find(&users).Error
	return users, err
}
func (r *UserRepository) GetByID(id int) (*models.User, error) {
	var user models.User
	err := r.db.First(&user).Error
	return &user, err
}
func (r *UserRepository) Update(id int, data *models.User) error {
	return r.db.Model(&models.User{}).
		Where("user_id = ?", id).
		Updates(data).Error
}

func (r *UserRepository) Delete(id int) error {
	return r.db.Delete(&models.User{}, id).Error
}
