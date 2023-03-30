package model

import (
	"context"
	"errors"
)

type User struct {
	ID          int64  `gorm:"primarykey;uniquel;not null"`
	Name        string `gorm:"not null"`
	Password    string `gorm:"type:varchar(256);not null"`
	Email       string `gorm:"unique;not null"`
	PhotoUrl    string
	Description string
}

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
