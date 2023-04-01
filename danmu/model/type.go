package model

import "time"

type Danmu struct {
	ID      int64     `gorm:"primaryKey;not null;unique"`
	VideoID int64     `gorm:"not null"`
	UserID  int64     `gorm:"not null"`
	Content string    `gorm:"not null"`
	BeginAt time.Time `gorm:"not null"`
}
