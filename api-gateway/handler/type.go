package handler

type Token struct {
	SS string `json:"token" form:"token" binding:"required"`
}

type File struct {
	FileName string `uri:"name" binding:"required"`
	ShareID  string `form:"share_id"`
}

type VideoID struct {
	ID int64 `uri:"video_id" binding:"required"`
}

type CommentID struct {
	ID int64 `uri:"comment_id" binding:"required"`
}
