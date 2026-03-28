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

func (r *UserRepository) RegisterPatient(user *models.User, patient *models.Patient) error {
	// เริ่ม Transaction
	return r.db.Transaction(func(tx *gorm.DB) error {
		// 1. บันทึก User ก่อน เพื่อให้ได้ ID มา
		if err := tx.Create(user).Error; err != nil {
			return err
		}

		// 2. เอา ID ที่ได้จาก User มาใส่ใน Patient (1:1 Relationship)
		patient.ID = user.ID
		if err := tx.Create(patient).Error; err != nil {
			return err // ถ้าพังตรงนี้ GORM จะ Rollback การสร้าง User ให้เอง
		}

		return nil
	})
}
