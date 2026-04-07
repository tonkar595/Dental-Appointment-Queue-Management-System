package dto

import "time"

type CreateAppointmentRequest struct {
	PatientID uint   `json:"patient_id" validate:"required"`
	StaffID   uint   `json:"staff_id" validate:"required"`
	ServiceID uint   `json:"service_id" validate:"required"`
	Date      string `json:"date" validate:"required"`       // เช่น "2026-04-10"
	StartTime string `json:"start_time" validate:"required"` // เช่น "10:30"
	Note      string `json:"treatment_note"`
	IsWalkIn  bool   `json:"is_walk_in"`
}

type AppointmentResponse struct {
	ID               uint      `json:"id"`
	PatientName      string    `json:"patient_name"`
	StaffName        string    `json:"staff_name"`
	ServiceName      string    `json:"service_name"`
	AppointmentStart time.Time `json:"appointment_start"`
	AppointmentEnd   time.Time `json:"appointment_end"`
	Status           string    `json:"status"`
}
