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

type AllAppointmentResponse struct {
	ID               uint      `json:"id"`
	AppointmentStart time.Time `json:"appointment_start"`
	AppointmentEnd   time.Time `json:"appointment_end"`
	TreatmentNote    string    `json:"treatment_note"`
	IsWalkIn         bool      `json:"is_walk_in"`
	PatientID        uint      `json:"patient_id"`
	PatientName      string    `json:"patient_name"`
	StaffID          uint      `json:"staff_id"`
	StaffName        string    `json:"staff_name"`
	StaffEmail       string    `json:"staff_email"`
	StaffPhone       string    `json:"staff_phone"`
	ServiceID        uint      `json:"service_id"`
	ServiceName      string    `json:"service_name"`
	Duration         int       `json:"duration_minutes"`
	StatusID         uint      `json:"status_id"`
	StatusName       string    `json:"status_name"`
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

type UpdateAppointmentRequest struct {
	AppointmentStart *string `json:"appointment_start"` // รับเป็น string "2026-04-10 10:00"
	ServiceID        *uint   `json:"service_id"`
	StatusID         *uint   `json:"status_id"`
	TreatmentNote    *string `json:"treatment_note"`
}

// Response ทั่วไปสำหรับ Update/Delete
type MessageResponse struct {
	Status  string      `json:"status"`  // เช่น "success" หรือ "error"
	Message string      `json:"message"` // เช่น "อัปเดตนัดหมายสำเร็จ"
	Data    interface{} `json:"data,omitempty"`
}
