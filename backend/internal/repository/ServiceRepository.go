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
func (r *ServiceRepository) Update(id uint, service *models.ServiceType) error {
	// ใช้ Updates เพื่ออัปเดตเฉพาะฟิลด์ที่ส่งมา
	updateData := map[string]interface{}{
		"service_name":     service.ServiceName,
		"description":      service.Description,
		"duration_minutes": service.DurationMinutes,
	}

	return r.db.Model(&models.ServiceType{}).Where("id = ?", id).Updates(updateData).Error

}

// ฟังก์ชันสำหรับ soft detelete หรือ toggle active/inactive
func (r *ServiceRepository) ToggleActive(id uint, status bool) error {
	return r.db.Model(&models.ServiceType{}).Where("id = ?", id).Update("is_active", status).Error
}

func (r *ServiceRepository) Restore(id uint) error {
	// กู้คืนโดยการตั้งค่า is_active = true
	return r.db.Model(&models.ServiceType{}).
		Where("id = ?", id).
		Update("is_active", true).Error
}
