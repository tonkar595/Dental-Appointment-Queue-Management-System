package service

import (
	"backend/internal/dto"
	models "backend/internal/model"
	"backend/internal/repository"
	"errors"
	"time"
)

type AppointmentService struct {
	repo          *repository.AppointmentRepository
	clinicService *ClinicService // เรียกใช้ Service ที่เราทำไว้ก่อนหน้า
}

func NewAppointmentService(repo *repository.AppointmentRepository, clinic *ClinicService) *AppointmentService {
	return &AppointmentService{repo: repo, clinicService: clinic}
}

func (s *AppointmentService) CreateAppointment(req dto.CreateAppointmentRequest) error {

	isDentistAsPatient, err := s.repo.ValidatePatient(req.PatientID)
	if err != nil {
		return errors.New("ไม่สามารถตรวจสอบข้อมูลคนไข้ได้")
	}

	if isDentistAsPatient {
		return errors.New("ไม่สามารถจองนัดหมายได้: ผู้ใช้งานที่มีสิทธิ์เป็นทันตแพทย์ไม่สามารถรับนัดในฐานะคนไข้ได้")
	}

	isDentist, err := s.repo.IsUserDentist(req.StaffID)
	if err != nil || !isDentist {
		return errors.New("ไม่สามารถจองได้: รหัสเจ้าหน้าที่นี้ไม่ใช่ทันตแพทย์")
	}
	// 1. เช็ค Clinic Open/Close (Default + Override Logic)
	availability := s.clinicService.GetEffectiveSchedule(req.Date)
	if !availability.IsOpen {
		return errors.New("ไม่สามารถจองได้: คลินิกปิดทำการ (" + availability.Note + ")")
	}

	// 2. ดึง Duration จาก Database เพื่อคำนวณเวลาจบอัตโนมัติ
	duration, err := s.repo.GetServiceDuration(req.ServiceID)
	if err != nil {
		return errors.New("ประเภทบริการไม่ถูกต้อง")
	}

	// 3. รวมร่างวันที่และเวลา (Parse String to time.Time)
	loc, _ := time.LoadLocation("Asia/Bangkok")
	startTime, errS := time.ParseInLocation("2006-01-02 15:04", req.Date+" "+req.StartTime, loc)
	if errS != nil {
		return errors.New("รูปแบบวันที่หรือเวลาไม่ถูกต้อง")
	}

	// คำนวณ EndTime อัตโนมัติ
	endTime := startTime.Add(time.Duration(duration) * time.Minute)

	// 4. เช็คว่าเวลาที่จองอยู่นอกเวลาทำการไหม (เช่น จอง 16:30 แต่คลินิกปิด 17:00 และบริการนี้ใช้เวลา 60 นาที)
	// (ขั้นตอนนี้แนะนำให้เพิ่มเพื่อให้ตารางนัดไม่เลยเวลาปิดคลินิก)

	// 5. เช็คว่าหมอว่างไหม (Overlap Check)
	busy, err := s.repo.IsSlotBusy(req.StaffID, startTime, endTime)
	if err != nil {
		return err
	}
	if busy {
		return errors.New("คุณหมอติดนัดหมายอื่นในช่วงเวลานี้ กรุณาเลือกเวลาใหม่")
	}

	// 6. บันทึกข้อมูล
	now := time.Now()
	appointment := &models.Appointment{
		PatientID:        req.PatientID,
		StaffID:          req.StaffID,
		ServiceID:        req.ServiceID,
		AppointmentStart: startTime,
		AppointmentEnd:   endTime,
		TreatmentNote:    req.Note,
		StatusID:         1, // 1 = Pending/Confirmed
		IsWalkIn:         req.IsWalkIn,
		BookedAt:         &now,
	}

	return s.repo.Create(appointment)
}

func (s *AppointmentService) GetAppointments(date string) ([]dto.AllAppointmentResponse, error) {
	appointments, err := s.repo.GetByDate(date)
	if err != nil {
		return nil, err
	}

	var response []dto.AllAppointmentResponse
	for _, a := range appointments {
		res := dto.AllAppointmentResponse{
			ID:               a.ID,
			AppointmentStart: a.AppointmentStart,
			AppointmentEnd:   a.AppointmentEnd,
			TreatmentNote:    a.TreatmentNote,
			IsWalkIn:         a.IsWalkIn,

			PatientID:   a.PatientID,
			PatientName: a.Patient.User.UserName, // ดึงจาก Nested User

			StaffID:    a.StaffID,
			StaffName:  a.Staff.UserName, // Staff ใน Model คือ User
			StaffEmail: a.Staff.Email,
			StaffPhone: a.Staff.Phone,

			ServiceID:   a.ServiceID,
			ServiceName: a.Service.ServiceName,
			Duration:    a.Service.DurationMinutes,

			StatusID:   a.StatusID,
			StatusName: a.Status.StatusName,
		}
		response = append(response, res)
	}

	return response, nil
}

// ดูประวัตินัดหมายของคนไข้
func (s *AppointmentService) GetPatientHistory(patientID uint) ([]dto.AppointmentResponse, error) {
	appointments, err := s.repo.GetByPatientID(patientID)
	if err != nil {
		return nil, err
	}

	var response []dto.AppointmentResponse
	for _, a := range appointments {
		res := dto.AppointmentResponse{
			ID:               a.ID,
			AppointmentStart: a.AppointmentStart,
			AppointmentEnd:   a.AppointmentEnd,
			TreatmentNote:    a.TreatmentNote,
			IsWalkIn:         a.IsWalkIn,

			// Mapping Nested: Patient
			Patient: dto.PatientDTO{
				ID:               a.PatientID,
				Name:             a.Patient.User.UserName,
				MedicalCondition: a.Patient.MedicalCondition,
			},

			// Mapping Nested: Staff (ข้อมูลจะมาครบตามที่เรา Query ไว้)
			Staff: dto.StaffDTO{
				ID:    a.StaffID,
				Name:  a.Staff.UserName,
				Email: a.Staff.Email,
				Phone: a.Staff.Phone,
			},

			// Mapping Nested: Service
			Service: dto.ServiceDTO{
				ID:       a.ServiceID,
				Name:     a.Service.ServiceName,
				Duration: a.Service.DurationMinutes,
			},

			// Mapping Nested: Status
			Status: dto.StatusDTO{
				ID:   a.StatusID,
				Name: a.Status.StatusName,
			},
		}
		response = append(response, res)
	}

	return response, nil
}

func (s *AppointmentService) PatchAppointment(id uint, req dto.UpdateAppointmentRequest) error {
	updateData := make(map[string]interface{})

	// 1. ถ้ามีการเปลี่ยน Service หรือ เวลา ต้องคำนวณเวลาจบใหม่
	if req.AppointmentStart != nil || req.ServiceID != nil {
		// ดึงข้อมูลเดิมมาตั้งต้นก่อน
		existing, _ := s.repo.GetByID(id)

		newStart := existing.AppointmentStart
		newServiceID := existing.ServiceID

		if req.AppointmentStart != nil {
			loc, _ := time.LoadLocation("Asia/Bangkok")
			newStart, _ = time.ParseInLocation("2006-01-02 15:04", *req.AppointmentStart, loc)
			updateData["appointment_start"] = newStart
		}
		if req.ServiceID != nil {
			newServiceID = *req.ServiceID
			updateData["service_id"] = newServiceID
		}

		// คำนวณ EndTime ใหม่ตาม Duration ของ Service
		duration, _ := s.repo.GetServiceDuration(newServiceID)
		updateData["appointment_end"] = newStart.Add(time.Duration(duration) * time.Minute)
	}

	// 2. ฟิลด์อื่นๆ ที่ไม่ต้องคำนวณ
	if req.StatusID != nil {
		updateData["status_id"] = *req.StatusID
	}
	if req.TreatmentNote != nil {
		updateData["treatment_note"] = *req.TreatmentNote
	}

	return s.repo.Update(id, updateData)
}

func (s *AppointmentService) DeleteAppointment(id uint) error {
	// เช็คก่อนว่านัดหมายมีตัวตนไหม
	_, err := s.repo.GetByID(id)
	if err != nil {
		return errors.New("ไม่พบข้อมูลนัดหมายที่ต้องการลบ")
	}

	// สั่งลบ
	return s.repo.Delete(id)
}
