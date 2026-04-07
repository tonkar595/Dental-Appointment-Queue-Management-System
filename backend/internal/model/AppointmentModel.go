package models

import "time"

type Appointment struct {
	ID        uint `gorm:"primaryKey" json:"id"`
	PatientID uint `json:"patient_id"`
	// DentistScheduleID uint       `json:"dentist_schedule_id"`
	StaffID          uint       `json:"staff_id"`
	ServiceID        uint       `json:"service_id"`
	AppointmentStart time.Time  `json:"appointment_start"`
	AppointmentEnd   time.Time  `json:"appointment_end"`
	TreatmentNote    string     `gorm:"type:text" json:"treatment_note"`
	StatusID         uint       `json:"status_id"`
	IsWalkIn         bool       `gorm:"default:false" json:"is_walk_in"`
	CreatedAt        time.Time  `json:"-"`
	UpdatedAt        time.Time  `json:"-"`
	BookedAt         *time.Time `json:"-"`

	Patient Patient           `gorm:"foreignKey:PatientID" json:"patient,omitempty"`
	Staff   User              `gorm:"foreignKey:StaffID" json:"staff,omitempty"` // ชี้ไปที่ User/Staff
	Service ServiceType       `gorm:"foreignKey:ServiceID" json:"service,omitempty"`
	Status  AppointmentStatus `gorm:"foreignKey:StatusID" json:"status,omitempty"`
}

type AppointmentStatus struct {
	ID          uint   `gorm:"primaryKey" json:"id"`
	StatusName  string `json:"status_name"`
	Description string `json:"description"`
}

type ServiceType struct {
	ID              uint   `gorm:"primaryKey" json:"id"`
	ServiceName     string `json:"service_name"`
	Description     string `json:"description"`
	DurationMinutes int    `json:"duration_minutes"`
	IsActive        bool   `gorm:"default:true" json:"is_active"`
}

func (AppointmentStatus) TableName() string {
	return "appointment_status"
}
