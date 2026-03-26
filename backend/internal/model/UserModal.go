package models

import "time"

type User struct {
	ID           uint      `gorm:"primaryKey" json:"id"`
	UserName     string    `gorm:"size:100;unique;not null" json:"user_name"`
	FirstName    string    `gorm:"size:100;not null" json:"first_name"`
	LastName     string    `gorm:"size:100;not null" json:"last_name"`
	Email        string    `gorm:"size:255;unique;not null" json:"email"`
	PasswordHash string    `gorm:"type:text;not null" json:"password_hash"`
	RoleID       uint      `json:"role_id"`
	IsActive     bool      `gorm:"default:true" json:"is_active"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`

	Role    Role
	Patient *Patient
	Staff   *Staff
}

type Role struct {
	ID       uint   `gorm:"primaryKey" json:"id"`
	RoleName string `gorm:"size:50;unique;not null" json:"role_name"`
}

type Patient struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Phone     string    `gorm:"size:20" json:"phone"`
	BirthDate time.Time `json:"birth_date"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Staff struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Phone     string    `gorm:"size:20" json:"phone"`
	Position  string    `gorm:"size:100" json:"position"`
	UpdatedAt time.Time `json:"updated_at"`
}
