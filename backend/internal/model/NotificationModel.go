package models

import "time"

type Notification struct {
	ID               uint      `gorm:"primaryKey" json:"id"`
	UserID           uint      `gorm:"not null" json:"user_id"`
	Message          string    `gorm:"type:text;not null" json:"message"`
	NotificationType string    `gorm:"size:50" json:"notification_type"`
	IsRead           bool      `gorm:"default:false" json:"is_read"`
	CreatedAt        time.Time `json:"created_at"`

	User User `gorm:"foreignKey:UserID"`
}
