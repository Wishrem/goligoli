package handler

type Token struct {
	SS string `json:"token" form:"token" binding:"required"`
}
