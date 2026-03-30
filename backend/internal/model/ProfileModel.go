package models

import "time"

type Patient struct {
	ID                 uint      `gorm:"primaryKey" json:"id"` // FK to User.ID
	BirthDate          time.Time `json:"birth_date"`
	MedicalCondition   string    `gorm:"type:text" json:"medical_condition"`
	AllergicMedication string    `gorm:"type:text" json:"allergic_medication"`
	UpdatedAt          time.Time `json:"updated_at"`
	User               User      `gorm:"foreignKey:ID;references:ID" json:"-"`
}

// type Staff struct {
// 	ID        uint      `gorm:"primaryKey" json:"id"` // FK to User.ID
// 	Phone     string    `json:"phone"`
// 	Position  string    `json:"position"`
// 	UpdatedAt time.Time `json:"updated_at"`
// 	User      User      `gorm:"foreignKey:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"user_info"`
// }

// func (Patient) TableName() string {
// 	return "patients"
// }
