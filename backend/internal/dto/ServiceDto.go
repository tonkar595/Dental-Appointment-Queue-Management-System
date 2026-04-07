package dto

type CreateServiceRequest struct {
	ServiceName     string `json:"service_name" validate:"required"`
	Description     string `json:"description"`
	DurationMinutes int    `json:"duration_minutes" validate:"required,min=1"`
}
