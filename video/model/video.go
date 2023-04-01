package model

import (
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func (v *Video) Create() error {
	return db.Create(v).Error
}

func (v *Video) View() error {
	return db.Model(&Video{}).Preload(clause.Associations).Where("video_url = ?", v.VideoUrl).First(v).Error
}

func (v *Video) Share() error {
	return db.Model(v).Update("shared", gorm.Expr("shared + 1")).Error
}

func (v *Video) Like() error {
	return db.Model(v).Update("liked", gorm.Expr("liked + 1")).Error
}

func (v *Video) JudgeThenGet() error {
	var status Status
	if err := db.Where("video_id = ?", v.ID).Take(&status).Error; err != nil {
		return err
	}
	if err := db.Model(&status).Updates(&Status{
		Reason: v.Status.Reason,
		Passed: v.Status.Passed,
	}).Error; err != nil {
		return err
	}
	return db.Take(&v).Error
}
