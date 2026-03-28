package models

import "time"

type User struct {
	ID           uint      `gorm:"primaryKey" json:"id"`
	UserName     string    `gorm:"unique;not null" json:"user_name"`
	Email        string    `gorm:"unique;not null" json:"email"`
	PasswordHash string    `json:"-"`
	RoleID       uint      `json:"role_id"`
	Phone        string    `gorm:"type:varchar(20)" json:"phone"`
	IsActive     bool      `gorm:"default:true" json:"is_active"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`

	Role    Role     `gorm:"foreignKey:RoleID" json:"role"`
	Patient *Patient `gorm:"foreignKey:ID" json:"patient,omitempty"`

	Notifications []Notifications `gorm:"foreignKey:UserID" json:"notifications"`
}

type Role struct {
	ID       uint   `gorm:"primaryKey" json:"id"`
	RoleName string `gorm:"unique;not null" json:"role_name"`
}

type Notifications struct {
	ID               uint      `gorm:"primaryKey" json:"id"`
	UserID           uint      `json:"user_id"`
	Message          string    `json:"message"`
	NotificationType string    `json:"notification_type"`
	IsRead           bool      `gorm:"default:false" json:"is_read"`
	CreatedAt        time.Time `json:"created_at"`
}

func (User) TableName() string {
	return "users"
}
