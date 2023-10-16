package models

type Service struct {
	ID          int64  `json:"id"`
	ServiceName string `json:"service_name"`
	About       string `json:"about"`
	ImageUrl    string `json:"image_url"`
}

type CreateService struct {
	ServiceName string `json:"service_name" binding:"required"`
	About       string `json:"about" binding:"required"`
	ImageUrl    string `json:"image_url" binding:"required"`
}

type AllServices struct {
	Services []*Service `json:"services"`
}
