package dto

// LoginRequest ใช้สำหรับรับข้อมูลจาก Client
type LoginRequest struct {
	Identity string `json:"identity" validate:"required"` // จะส่งเป็น Username หรือ Email ก็ได้
	Password string `json:"password" validate:"required"`
}

// LoginResponse ใช้สำหรับส่งข้อมูลกลับ (ถ้าต้องการซ่อนบางฟิลด์)
type LoginResponse struct {
	Token string `json:"token"`
	Type  string `json:"type"`
}

// type RegisterRequest struct {
// 	UserName  string    `json:"user_name" validate:"required"`
// 	Email     string    `json:"email" validate:"required,email"`
// 	Password  string    `json:"password" validate:"required,min=6"`
// 	FirstName string    `json:"first_name"`
// 	LastName  string    `json:"last_name"`
// 	Phone     string    `json:"phone"`
// 	BirthDate time.Time `json:"birth_date"`
// }
type RegisterRequest struct {
	UserName string `json:"user_name" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
}
