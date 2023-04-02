package model

import (
	"context"
	"errors"

	"gorm.io/gorm/clause"
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
	if u.ID == 0 {
		return errors.New("no index to find")
	}
	if err := db.Model(&Ban{}).Where("user_id = ?", u.ID).Updates(u.Ban).Error; err != nil {
		return err
	}

	return db.Model(u).Preload(clause.Associations).Take(u).Error
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
