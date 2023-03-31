package model

import (
	"context"
	"errors"
	"time"

	"gorm.io/gorm/clause"
)

type (
	User struct {
		ID          int64  `gorm:"primaryKey;unique;not null"`
		Name        string `gorm:"not null"`
		Password    string `gorm:"type:varchar(256);not null"`
		Email       string `gorm:"unique;not null"`
		PhotoUrl    string `gorm:"not null"`
		Description string
		Ban         Ban    `gorm:"foreignKey:UserID;references:ID"`
		Roles       []Role `gorm:"foreignKey:UserID;not null;references:ID"`
	}

	Ban struct {
		UserID  int64     `gorm:"primaryKey;unique;not null"`
		Reason  string    `gorm:"not null"`
		BanAt   time.Time `gorm:"not null"`
		UnbanAt time.Time `gorm:"not null"`
	}

	Role struct {
		UserID int64  `gorm:"primaryKey;unique;not null"`
		Type   string `gorm:"not null"`
	}
)

func (u *User) Get(ctx context.Context) error {
	if u.ID != 0 {
		return db.Model(&User{}).Preload(clause.Associations).Where("id = ?", u.ID).First(&u).Error
	} else if u.Email != "" {
		return db.Model(&User{}).Preload(clause.Associations).Where("email = ?", u.Email).First(&u).Error
	}
	return errors.New("no index to find")
}

func (u *User) Create(ctx context.Context) error {
	return db.Create(u).Error
}

func (u *User) UpdateBan(ctx context.Context) error {
	if u.ID != 0 {
		return db.Model(&User{}).Preload("Ban").Select("ban").Where("id = ?", u.ID).Updates(&u.Ban).Error
	}
	return errors.New("no index to find")
}

func (u *User) UpdateInfo(ctx context.Context) error {
	if u.ID != 0 {
		if u.Description != "" && u.PhotoUrl != "" {
			return db.Model(&User{}).Select("description photo_url").Where("id = ?", u.ID).Updates(&User{PhotoUrl: u.PhotoUrl, Description: u.Description}).Error
		} else if u.Description != "" {
			return db.Model(&User{}).Select("description").Where("id = ?", u.ID).Updates(&User{Description: u.Description}).Error
		} else if u.PhotoUrl != "" {
			return db.Model(&User{}).Select("photo_url").Where("id = ?", u.ID).Updates(&User{PhotoUrl: u.PhotoUrl}).Error
		}
	}
	return errors.New("no index to find")
}
