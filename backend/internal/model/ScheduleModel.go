package models

import "time"

type ClinicSchedule struct {
	ClinicDate time.Time `gorm:"primaryKey;type:date" json:"clinic_date"`
	OpenTime   time.Time `json:"open_time"`
	CloseTime  time.Time `json:"close_time"`
	IsOpen     bool      `json:"is_open"`
	Note       string    `json:"note"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

type DentistSchedule struct {
	ID         uint      `gorm:"primaryKey" json:"id"`
	StaffID    uint      `json:"staff_id"`
	ClinicDate time.Time `gorm:"type:date" json:"clinic_date"`
	StartTime  time.Time `json:"start_time"`
	EndTime    time.Time `json:"end_time"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`

	Staff Staff `json:"staff"`
}
