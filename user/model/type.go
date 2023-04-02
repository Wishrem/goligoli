package model

import "time"

type (
	User struct {
		ID          int64  `gorm:"primaryKey;unique;not null"`
		Name        string `gorm:"not null"`
		Password    string `gorm:"type:varchar(256);not null"`
		Email       string `gorm:"unique;not null"`
		PhotoUrl    string `gorm:"not null"`
		Description string
		Ban         *Ban    `gorm:"foreignKey:UserID;references:ID"`
		Roles       []*Role `gorm:"foreignKey:UserID;not null;references:ID"`
	}

	Ban struct {
		ID      int64     `gorm:"primaryKey;unique;not null"`
		UserID  int64     `gorm:"unique;not null"`
		Reason  string    `gorm:"not null"`
		BanAt   time.Time `gorm:"not null"`
		UnbanAt time.Time `gorm:"not null"`
	}

	Role struct {
		ID     int64  `gorm:"primaryKey;unique;not null"`
		UserID int64  `gorm:"unique;not null"`
		Type   string `gorm:"not null"`
	}
)
