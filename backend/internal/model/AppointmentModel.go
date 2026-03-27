package models

import "time"

type Appointment struct {
	ID                uint       `gorm:"primaryKey" json:"id"`
	PatientID         uint       `json:"patient_id"`
	DentistScheduleID uint       `json:"dentist_schedule_id"`
	ServiceID         uint       `json:"service_id"`
	AppointmentStart  time.Time  `json:"appointment_start"`
	AppointmentEnd    time.Time  `json:"appointment_end"`
	StatusID          uint       `json:"status_id"`
	IsWalkIn          bool       `gorm:"default:false" json:"is_walk_in"`
	CreatedAt         time.Time  `json:"created_at"`
	UpdatedAt         time.Time  `json:"updated_at"`
	BookedAt          *time.Time `json:"booked_at"`

	Patient           Patient           `json:"patient"`
	ServiceType       ServiceType       `json:"service_type" gorm:"foreignKey:ServiceID"`
	AppointmentStatus AppointmentStatus `json:"status" gorm:"foreignKey:StatusID"`
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
