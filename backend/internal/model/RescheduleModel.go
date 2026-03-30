package models

import "time"

type RescheduleRequest struct {
	ID            uint       `gorm:"primaryKey" json:"id"`
	AppointmentID uint       `json:"appointment_id"`
	RequestedDate time.Time  `gorm:"type:date" json:"requested_date"`
	RequestedTime time.Time  `json:"requested_time"`
	Reason        string     `json:"reason"`
	StatusID      uint       `json:"status_id"`
	CreatedAt     time.Time  `json:"created_at"`
	ApprovedBy    *uint      `json:"approved_by"`
	ApprovedAt    *time.Time `json:"approved_at"`

	Status RescheduleStatus `json:"status" gorm:"foreignKey:StatusID"`
}

type RescheduleStatus struct {
	ID          uint   `gorm:"primaryKey" json:"id"`
	StatusName  string `json:"status_name"`
	Description string `json:"description"`
}
