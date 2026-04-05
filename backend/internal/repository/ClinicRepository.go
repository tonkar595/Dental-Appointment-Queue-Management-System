package repository

import (
	models "backend/internal/model"

	"gorm.io/gorm"
)

type ClinicRepository struct {
	db *gorm.DB
}

func NewClinicRepository(db *gorm.DB) *ClinicRepository {
	return &ClinicRepository{db: db}
}

// บันทึกหรืออัปเดตวันเปิด-ปิด (Upsert)
func (r *ClinicRepository) SaveSchedule(schedule *models.ClinicSchedule) error {
	return r.db.Save(schedule).Error
}

// ดึงข้อมูลวันเปิด-ปิด เฉพาะวันที่ต้องการ
func (r *ClinicRepository) GetScheduleByDate(date string) (*models.ClinicSchedule, error) {
	var schedule models.ClinicSchedule
	result := r.db.Select("clinic_date", "open_time", "close_time", "is_open", "note").
		Where("clinic_date = ?", date).
		Limit(1).
		Find(&schedule)
	if result.Error != nil {
		return nil, result.Error
	}

	// เช็คว่าเจอข้อมูลไหม (ถ้า RowsAffected เป็น 0 แปลว่าไม่เจอ)
	if result.RowsAffected == 0 {
		return nil, gorm.ErrRecordNotFound
	}

	return &schedule, nil
}

// ดึงรายการวันหยุดทั้งหมด (เอาไว้ทำปฏิทินหน้าบ้าน)
func (r *ClinicRepository) GetHolidays() ([]models.ClinicSchedule, error) {
	var holidays []models.ClinicSchedule
	err := r.db.Select("clinic_date", "note").
		Where("is_open = ?", false).
		Find(&holidays).Error
	return holidays, err
}
