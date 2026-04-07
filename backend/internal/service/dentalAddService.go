package service

import (
	"backend/internal/dto"
	models "backend/internal/model"
	"backend/internal/repository"
)

type DentalAddService struct {
	repo *repository.ServiceRepository
}

func NewDentalAddService(repo *repository.ServiceRepository) *DentalAddService {
	return &DentalAddService{repo: repo}
}

func (s *DentalAddService) CreateService(req dto.CreateServiceRequest) error {
	// เตรียม Model
	addService := &models.ServiceType{
		ServiceName:     req.ServiceName,
		Description:     req.Description,
		DurationMinutes: req.DurationMinutes,
		IsActive:        true,
	}

	return s.repo.Create(addService)
}

func (s *DentalAddService) GetAllServices() ([]models.ServiceType, error) {
	return s.repo.GetAll()
}

func (s *DentalAddService) GetServiceByID(id uint) (*models.ServiceType, error) {
	return s.repo.GetByID(id)
}

func (s *DentalAddService) UpdateService(id uint, req dto.CreateServiceRequest) error {
	updateData := &models.ServiceType{
		ServiceName:     req.ServiceName,
		Description:     req.Description,
		DurationMinutes: req.DurationMinutes,
	}
	return s.repo.Update(id, updateData)
}

func (s *DentalAddService) DeactivateService(id uint) error {
	return s.repo.ToggleActive(id, false)
}

func (s *DentalAddService) RestoreService(id uint) error {
	// เราอาจจะเช็คก่อนก็ได้ว่า ID นี้มีอยู่จริงไหม หรือปล่อยให้ Repo จัดการเลย
	return s.repo.Restore(id)
}
