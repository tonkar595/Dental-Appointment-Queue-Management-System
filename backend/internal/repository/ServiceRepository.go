package repository

import (
	models "backend/internal/model"

	"gorm.io/gorm"
)

type ServiceRepository struct {
	db *gorm.DB
}

func NewServiceRepository(db *gorm.DB) *ServiceRepository {
	return &ServiceRepository{db: db}
}

func (r *ServiceRepository) Create(service *models.ServiceType) error {
	return r.db.Create(service).Error
}

func (r *ServiceRepository) GetAll() ([]models.ServiceType, error) {
	var services []models.ServiceType
	// ดึงเฉพาะตัวที่ IsActive = true (หรือจะดึงหมดก็ได้ถ้าเป็นหน้า Admin)
	err := r.db.Find(&services).Error
	return services, err
}

func (r *ServiceRepository) GetByID(id uint) (*models.ServiceType, error) {
	var service models.ServiceType
	if err := r.db.First(&service, id).Error; err != nil {
		return nil, err
	}
	return &service, nil
}
