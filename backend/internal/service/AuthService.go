package service

import (
	"backend/internal/dto"
	models "backend/internal/model"
	"backend/internal/repository"
	"errors"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	repo *repository.UserRepository
}

func NewAuthService(repo *repository.UserRepository) *AuthService {
	return &AuthService{repo: repo}
}

func isEmail(identifier string) bool {
	return strings.Contains(identifier, "@")
}

func (s *AuthService) Login(identity, password string) (string, error) {
	var user *models.User
	var err error

	// 1. เลือกใช้ function ใน repo ตามรูปแบบ identity
	if isEmail(identity) {
		user, err = s.repo.FindByEmail(identity)
	} else {
		user, err = s.repo.FindByUsername(identity)
	}

	// ถ้า query แล้วเกิด error หรือไม่พบ user
	if err != nil {
		return "", errors.New("invalid credentials")
	}

	// 2. ตรวจสอบ Password (เปรียบเทียบ password ที่รับมากับ hash ใน DB)
	// user.PasswordHash คือตัวแปรที่เราเก็บไว้ใน model User
	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password))
	if err != nil {
		return "", errors.New("invalid credentials")
	}

	// 3. สร้าง JWT Token เมื่อตรวจสอบผ่านแล้ว
	// กำหนด Claims (ข้อมูลที่จะฝังไปกับ Token)
	claims := jwt.MapClaims{
		"user_id": user.ID,
		"role":    user.Role.RoleName,                    // ดึงจาก Preload Role
		"exp":     time.Now().Add(time.Hour * 24).Unix(), // หมดอายุใน 24 ชม.
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// 4. Sign Token ด้วย Secret Key
	// หมายเหตุ: ควรดึง "your_secret_key" มาจาก os.Getenv("JWT_SECRET") เพื่อความปลอดภัย
	tokenString, err := token.SignedString([]byte("your_secret_key"))
	if err != nil {
		return "", errors.New("could not generate token")
	}

	return tokenString, nil
}
func (s *AuthService) Register(req dto.RegisterRequest) error {
	// 1. Hash Password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	// 2. เตรียม User Model (Phone ย้ายมาอยู่ที่นี่แล้ว)
	user := &models.User{
		UserName:     req.UserName,
		Email:        req.Email,
		PasswordHash: string(hashedPassword),
		Phone:        req.Phone, // บันทึกเบอร์โทรที่ตาราง User
		RoleID:       2,         // 2 คือ Patient ตามตาราง Roles
		IsActive:     true,
	}

	// 3. เตรียม Patient Model (เก็บข้อมูลสุขภาพ)
	patient := &models.Patient{
		UpdatedAt: time.Now(),
	}

	// 4. ส่งไปบันทึกที่ Repo
	return s.repo.RegisterPatient(user, patient)
}
