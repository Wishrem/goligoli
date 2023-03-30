package model

import (
	"context"
	"errors"
)

type (
	User struct {
		ID          int64  `gorm:"primarykey;unique;not null"`
		Name        string `gorm:"not null"`
		Password    string `gorm:"type:varchar(256);not null"`
		Email       string `gorm:"unique;not null"`
		PhotoUrl    string `gorm:"not null"`
		Description string
		Ban         bool   `gorm:"not null"`
		Roles       []Role `gorm:"foreignKey:UserID;not null;"`
	}

	Role struct {
		UserID int64  `gorm:"primarykey;unique;not null"`
		Type   string `gorm:"not null"`
	}
)

func (u *User) Get(ctx context.Context) error {
	if u.ID != 0 {
		return db.Model(&User{}).First(&u, u.ID).Error
	}
	if u.Email != "" {
		return db.Model(&User{}).Where("email = ?", u.Email).First(&u).Error
	}
	return errors.New("no index to find")
}

func (u *User) Create(ctx context.Context) error {
	return db.Create(u).Error
}
