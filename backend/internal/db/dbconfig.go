package db

import (
	models "backend/internal/model"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func DatabaseConnected() *gorm.DB {
	godotenv.Load()

	dsn := os.Getenv("DATABASE_URL")

	if dsn == "" {
		panic("DATABASE_URL environment variable is not set")
	}

	dbConnection, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(fmt.Sprintf("failed to connect database: %v", err))
	}

	DB = dbConnection

	fmt.Println("Database connected successfully")

	return DB
}

func AutoMigrate(database *gorm.DB) {
	database.AutoMigrate(
		// &models.Role{},
		&models.User{},
		// &models.Patient{},
		// &models.Staff{},
		// &models.ServiceType{},
		// &models.AppointmentStatus{},
		// &models.ClinicSchedule{},
		// &models.DentistSchedule{},
		// &models.Appointment{},
		// &models.RescheduleRequest{},
		// &models.RescheduleStatus{},
		// &models.Notifications{},
	)
}

func SeedStatuses(db *gorm.DB) {

	// Seed AppointmentStatus
	appStatuses := []models.AppointmentStatus{
		{ID: 1, StatusName: "Pending", Description: "รอการยืนยัน"},
		{ID: 2, StatusName: "Confirmed", Description: "ยืนยันนัดแล้ว"},
		{ID: 3, StatusName: "Checked-in", Description: "คนไข้มาถึงคลินิกแล้ว/รอเข้าพบ"},
		{ID: 4, StatusName: "In-Progress", Description: "กำลังดำเนินการรักษา"},
		{ID: 5, StatusName: "Completed", Description: "เสร็จสิ้นการรักษา"},
		{ID: 6, StatusName: "Cancelled", Description: "ยกเลิกแล้ว"},
	}
	for _, s := range appStatuses {
		db.Table("appointment_status").Where("id = ?", s.ID).FirstOrCreate(&s)
	}

	// Seed RescheduleStatus
	reStatuses := []models.RescheduleStatus{
		{ID: 1, StatusName: "Waiting", Description: "รอหมอพิจารณา"},
		{ID: 2, StatusName: "Approved", Description: "อนุมัติให้เลื่อนนัดแล้ว"},
		{ID: 3, StatusName: "Rejected", Description: "ไม่อนุมัติ"},
	}
	for _, s := range reStatuses {
		db.Table("reschedule_status").Where("id = ?", s.ID).FirstOrCreate(&s)
	}
}
