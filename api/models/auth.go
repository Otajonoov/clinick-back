package models

type LoginReq struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type LoginRes struct {
	Token       string `json:"token"`
	Username  string `json:"username" binding:"required"`
}
