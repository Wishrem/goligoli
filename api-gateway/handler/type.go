package handler

type Token struct {
	SS string `json:"token" form:"token" binding:"required"`
}

type File struct {
	FileName string `uri:"name" binding:"required"`
	ShareID  string `form:"share_id"`
}
