package service

import (
	"backend/internal/dto"
	models "backend/internal/model"
	"backend/internal/repository"
	"time"
)

type ClinicService struct {
	repo *repository.ClinicRepository
}

func NewClinicService(repo *repository.ClinicRepository) *ClinicService {
	return &ClinicService{repo: repo}
}

func (s *ClinicService) GetEffectiveSchedule(dateStr string) dto.ClinicAvailabilityResponse {
	parsedDate, err := time.Parse("2006-01-02", dateStr)
	if err != nil {
		// กรณี Format วันที่ส่งมาผิด ให้ตีเป็นปิดไปก่อนเพื่อความปลอดภัย
		return dto.ClinicAvailabilityResponse{IsOpen: false, Note: "รูปแบบวันที่ไม่ถูกต้อง"}
	}

	// 1. เช็คใน Database ก่อน (Override Logic)
	schedule, err := s.repo.GetScheduleByDate(dateStr)
	if err == nil && schedule != nil {
		return dto.ClinicAvailabilityResponse{
			IsOpen:    schedule.IsOpen,
			OpenTime:  schedule.OpenTime,
			CloseTime: schedule.CloseTime,
			Note:      schedule.Note,
			IsCustom:  true,
		}
	}

	// 2. ถ้าไม่มีใน DB ให้ใช้ค่า Default (จ-ศ เปิด, ส-อา ปิด)
	weekday := parsedDate.Weekday()
	if weekday == time.Saturday || weekday == time.Sunday {
		return dto.ClinicAvailabilityResponse{
			IsOpen:   false,
			Note:     "วันหยุดประจำสัปดาห์",
			IsCustom: false,
		}
	}

	return dto.ClinicAvailabilityResponse{
		IsOpen:    true,
		OpenTime:  "09:00",
		CloseTime: "17:00",
		Note:      "วันทำการปกติ",
		IsCustom:  false,
	}
}

func (s *ClinicService) UpdateCustomSchedule(req dto.UpdateClinicScheduleRequest) error {
	parsedDate, _ := time.Parse("2006-01-02", req.ClinicDate)

	// เตรียมเวลา (ถ้า IsOpen เป็น false เวลาจะเป็นค่าว่าง)
	// var oTime, cTime time.Time
	// if req.IsOpen {
	// 	oTime, _ = time.Parse("15:04", req.OpenTime)
	// 	cTime, _ = time.Parse("15:04", req.CloseTime)
	// }

	schedule := &models.ClinicSchedule{
		ClinicDate: parsedDate,
		OpenTime:   req.OpenTime,
		CloseTime:  req.CloseTime,
		IsOpen:     req.IsOpen,
		Note:       req.Note,
	}

	return s.repo.SaveSchedule(schedule)
}
