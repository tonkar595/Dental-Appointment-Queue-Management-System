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

func (r *UserRepository) FindByUsername(username string) (*models.User, error) {
	var user models.User
	err := r.db.Model(&models.User{}).
		Preload("Role").Preload("Patient").Preload("Staff").
		Where("user_name = ? AND is_active = ?", username, true).
		First(&user).Error
	return &user, err
}

// ค้นหาด้วย Email เท่านั้น
func (r *UserRepository) FindByEmail(email string) (*models.User, error) {
	var user models.User
	err := r.db.Model(&models.User{}).
		Preload("Role").Preload("Patient").Preload("Staff").
		Where("email = ? AND is_active = ?", email, true).
		First(&user).Error
	return &user, err
}
