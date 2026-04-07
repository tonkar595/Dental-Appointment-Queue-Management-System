package dto

type CreateDentistScheduleRequest struct {
	StaffID    uint   `json:"staff_id" validate:"required"`
	ClinicDate string `json:"clinic_date" validate:"required"` // Format: 2006-01-02
	StartTime  string `json:"start_time" validate:"required"`  // Format: 15:04
	EndTime    string `json:"end_time" validate:"required"`    // Format: 15:04
}

// สำหรับส่งข้อมูลออก (เลือกเฉพาะที่จำเป็น)
type DentistScheduleResponse struct {
	ID        uint   `json:"id"`
	StaffID   uint   `json:"staff_id"`
	StartTime string `json:"start_time"`
	EndTime   string `json:"end_time"`
}

type UpdateClinicScheduleRequest struct {
	ClinicDate string `json:"clinic_date" validate:"required"` // "2026-04-13"
	OpenTime   string `json:"open_time"`                       // "10:00"
	CloseTime  string `json:"close_time"`                      // "19:00"
	IsOpen     bool   `json:"is_open"`                         // true=เปิดพิเศษ, false=หยุดพิเศษ
	Note       string `json:"note"`                            // "วันสงกรานต์"
}

type ClinicAvailabilityResponse struct {
	IsOpen    bool   `json:"is_open"`
	OpenTime  string `json:"open_time"`
	CloseTime string `json:"close_time"`
	Note      string `json:"note"`
	IsCustom  bool   `json:"is_custom"` // บอกว่าอันนี้คือค่าจาก DB หรือ Default
}
