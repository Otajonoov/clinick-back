package models

type Customer struct {
	ID        int64  `json:"id"`
	Fullname  string `json:"fullname"`
	Stars     string `json:"stars"`
	About     string `json:"about"`
	ImageUrl  string `json:"image_url"`
	CreatedAt string `json:"created_at"`
}

type CreateCustomer struct {
	Fullname string `json:"fullname" binding:"required"`
	Stars    string `json:"stars" binding:"required"`
	About    string `json:"about" binding:"required"`
	ImageUrl string `json:"image_url" binding:"required"`
}

type AllCustomers struct {
	Customers []*Customer `json:"customers"`
}

type CustomersResp struct {
	Customers []*Customer `json:"customers"`
	Count     int64       `json:"count"`
}

type CustomersFindReq struct {
	Limit int64 `json:"limit" binding:"required" default:"10"`
	Page  int64 `json:"page" binding:"required" default:"1"`
}
