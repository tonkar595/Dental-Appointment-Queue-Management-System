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
	ID               uint       `json:"id"`
	AppointmentStart time.Time  `json:"appointment_start"`
	AppointmentEnd   time.Time  `json:"appointment_end"`
	TreatmentNote    string     `json:"treatment_note"`
	IsWalkIn         bool       `json:"is_walk_in"`
	Patient          PatientDTO `json:"patient"`
	Staff            StaffDTO   `json:"staff"`
	Service          ServiceDTO `json:"service"`
	Status           StatusDTO  `json:"status"`
}

type PatientDTO struct {
	ID               uint   `json:"id"`
	Name             string `json:"name"`
	MedicalCondition string `json:"medical_condition"`
}

type StaffDTO struct {
	ID    uint   `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Phone string `json:"phone"`
}

type ServiceDTO struct {
	ID       uint   `json:"id"`
	Name     string `json:"name"`
	Duration int    `json:"duration_minutes"`
}

type StatusDTO struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}
