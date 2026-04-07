package wire

import (
	"backend/internal/controller"
	"backend/internal/repository"
	"backend/internal/service"

	"gorm.io/gorm"
)

type Container struct {
	UserController        *controller.UserController
	AuthController        *controller.AuthController
	DentalAddController   *controller.DentalAddController
	ClinicController      *controller.ClinicController
	AppointmentController *controller.AppointmentController
}

func NewContainer(db *gorm.DB) *Container {
	userRepo := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepo)
	userController := controller.NewUserController(userService)
	authService := service.NewAuthService(userRepo)
	authController := controller.NewAuthController(authService)
	ServiceRepo := repository.NewServiceRepository(db)
	dentalAddService := service.NewDentalAddService(ServiceRepo)
	dentalAddController := controller.NewDentalAddController(dentalAddService)
	appointmentRepo := repository.NewAppointmentRepository(db)

	clinicRepo := repository.NewClinicRepository(db)
	clinicService := service.NewClinicService(clinicRepo)
	clinicController := controller.NewClinicController(clinicService)
	appointmentService := service.NewAppointmentService(appointmentRepo, clinicService)

	appointmentController := controller.NewAppointmentController(appointmentService)

	return &Container{
		UserController:        userController,
		AuthController:        authController,
		DentalAddController:   dentalAddController,
		ClinicController:      clinicController,
		AppointmentController: appointmentController,
	}
}
