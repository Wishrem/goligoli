package model

import "time"

type (
	Comment struct {
		ID       int64      `gorm:"primaryKey;not null;unique"`
		Content  string     `gorm:"not null"`
		SentAt   time.Time  `gorm:"not null"`
		Responds []Response `gorm:"foreignKey:CommentID;references:ID"`
	}

	Response struct {
		ID        int64     `gorm:"primaryKey;not null;unique"`
		CommentID int64     `gorm:"not null"`
		Content   string    `gorm:"not null"`
		SentAt    time.Time `gorm:"not null"`
	}
)
