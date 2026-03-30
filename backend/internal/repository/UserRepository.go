package repository

import (
	models "backend/internal/model"
	"errors"
	"strings"

	"github.com/jackc/pgx/v5/pgconn"
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
		Preload("Role").Preload("Patient").
		Where("user_name = ? AND is_active = ?", username, true).
		First(&user).Error
	return &user, err
}

// ค้นหาด้วย Email เท่านั้น
func (r *UserRepository) FindByEmail(email string) (*models.User, error) {
	var user models.User
	err := r.db.Model(&models.User{}).
		Preload("Role").Preload("Patient").
		Where("email = ? AND is_active = ?", email, true).
		First(&user).Error
	return &user, err
}

func (r *UserRepository) FindByPhone(phone string) (*models.User, error) {
	var user models.User
	err := r.db.Model(&models.User{}).
		Preload("Role").Preload("Patient").
		Where("phone = ? AND is_active = ?", phone, true).
		First(&user).Error
	return &user, err
}

func (r *UserRepository) RegisterPatient(user *models.User, patient *models.Patient) error {
	return r.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(user).Error; err != nil {
			// ดักจับ Error จาก Postgres (Unique Violation)
			var pgErr *pgconn.PgError
			if errors.As(err, &pgErr) && pgErr.Code == "23505" {
				// พยายามระบุว่าอะไรซ้ำ (Email, Username หรือ Phone)
				if strings.Contains(pgErr.Detail, "email") {
					return errors.New("email already exists")
				}
				if strings.Contains(pgErr.Detail, "user_name") {
					return errors.New("username already exists")
				}
				if strings.Contains(pgErr.Detail, "phone") {
					return errors.New("phone number already exists")
				}
				return errors.New("duplicate entry")
			}
			return err
		}

		patient.ID = user.ID
		if err := tx.Create(patient).Error; err != nil {
			return err
		}
		return nil
	})
}
