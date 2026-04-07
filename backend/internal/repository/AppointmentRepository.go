package repository

import (
	models "backend/internal/model"
	"time"

	"gorm.io/gorm"
)

type AppointmentRepository struct {
	db *gorm.DB
}

func NewAppointmentRepository(db *gorm.DB) *AppointmentRepository {
	return &AppointmentRepository{db: db}
}

func (r *AppointmentRepository) IsSlotBusy(staffID uint, start, end time.Time) (bool, error) {
	var count int64
	// สูตร: (จองใหม่เริ่มก่อนนัดเก่าจบ) และ (จองใหม่จบหลังนัดเก่าเริ่ม)
	err := r.db.Model(&models.Appointment{}).
		Where("staff_id = ? AND status_id != ?", staffID, 6). // ไม่นับนัดที่ถูกยกเลิก (6 คือ Cancelled)
		Where("appointment_start < ? AND appointment_end > ?", end, start).
		Count(&count).Error

	return count > 0, err
}

func (r *AppointmentRepository) Create(appointment *models.Appointment) error {
	return r.db.Create(appointment).Error
}

func (r *AppointmentRepository) GetServiceDuration(serviceID uint) (int, error) {
	var service models.ServiceType
	err := r.db.First(&service, serviceID).Error
	return service.DurationMinutes, err
}
func (r *AppointmentRepository) IsUserDentist(staffID uint) (bool, error) {
	var roleName string
	// Join ตาราง users กับ roles เพื่อดูชื่อ Role
	err := r.db.Table("users").
		Select("roles.role_name").
		Joins("join roles on roles.id = users.role_id").
		Where("users.id = ?", staffID).
		Row().Scan(&roleName)

	if err != nil {
		return false, err
	}
	// ดักเฉพาะคนที่เป็น Dentist (หรือชื่อ Role ที่คุณตั้งไว้)
	return roleName == "Dentist", nil
}
func (r *AppointmentRepository) ValidatePatient(patientID uint) (bool, error) {
	var roleName string
	// เจาะจงเช็คจากตาราง patients -> users -> roles
	err := r.db.Table("patients").
		Select("roles.role_name").
		Joins("join users on users.id = patients.user_id").
		Joins("join roles on roles.id = users.role_id").
		Where("patients.id = ?", patientID).
		Row().Scan(&roleName)

	if err != nil {
		return false, err
	}

	// ถ้าชื่อ Role เป็น Dentist แสดงว่า ID นี้คือหมอปลอมตัวมาเป็นคนไข้
	return roleName == "Dentist", nil
}
