package models

type Doctor struct {
	ID       int64  `json:"id"`
	Fullname string `json:"fullname"`
	Type     string `json:"type"`
	About    string `json:"about"`
	ImageUrl string `json:"image_url"`
}

type CreateDoctor struct {
	Fullname string `json:"fullname" binding:"required"`
	Type     string `json:"type" binding:"required"`
	About    string `json:"about" binding:"required"`
	ImageUrl string `json:"image_url" binding:"required"`
}

type AllDoctors struct {
	Doctors []*Doctor `json:"doctors"`
}

type ResponseError struct {
	Message string `json:"message"`
}

type ResponseOK struct {
	Message string `json:"message"`
}
