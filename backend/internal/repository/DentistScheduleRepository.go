package repository

import (
	models "backend/internal/model"

	"gorm.io/gorm"
)

type DentistScheduleRepository struct {
	db *gorm.DB
}

func NewDentistScheduleRepository(db *gorm.DB) *DentistScheduleRepository {
	return &DentistScheduleRepository{db: db}
}

func (r *DentistScheduleRepository) GetByDate(date string) ([]models.DentistSchedule, error) {
	var schedules []models.DentistSchedule

	// ใช้ .Select() เพื่อดึงเฉพาะฟิลด์ที่จำเป็นต่อการแสดงผล/คำนวณคิว
	err := r.db.Model(&models.DentistSchedule{}).
		Select("id", "staff_id", "start_time", "end_time").
		Where("clinic_date = ?", date).
		Find(&schedules).Error

	return schedules, err
}

// ตรวจสอบเวรซ้ำ (ใช้เฉพาะ ID ตัวเดียวมานับจำนวนก็พอ ไม่ต้องเอาทั้งแถว)
func (r *DentistScheduleRepository) CheckOverlap(staffID uint, date string, start, end string) (bool, error) {
	var count int64
	err := r.db.Model(&models.DentistSchedule{}).
		Where("staff_id = ? AND clinic_date = ?", staffID, date).
		Where("NOT (end_time <= ? OR start_time >= ?)", start, end). // Logic Overlap ที่กระชับขึ้น
		Count(&count).Error
	return count > 0, err
}
