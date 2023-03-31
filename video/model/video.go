package model

import "gorm.io/gorm/clause"

func (v *Video) Create() error {
	return db.Create(v).Error
}

func (v *Video) View() error {
	return db.Model(&Video{}).Where("video_url = ?", v.VideoUrl).First(v).Error
}

func (v *Video) GetThenShare() error {
	if err := db.First(v).Error; err != nil {
		return err
	}
	return db.Model(&Video{}).Where("id = ?", v.ID).Select("shared").Updates(&Video{Shared: v.Shared + 1}).Error
}

func (v *Video) Like() error {
	if err := db.Model(&Video{}).Where("id = ?", v.ID).First(v).Error; err != nil {
		return err
	}
	return db.Model(&Video{}).Select("linked").Updates(Video{Shared: v.Liked + 1}).Error
}

func (v *Video) JudgeThenGet() error {
	if err := db.Model(v).Select("Status").Updates(&Video{
		Status: Status{
			VideoID: v.ID,
			Passed:  v.Status.Passed,
			Reason:  v.Status.Reason,
		},
	}).Error; err != nil {
		return err
	}
	return db.Model(&Video{}).Where("id = ?", v.ID).Preload(clause.Associations).First(&v).Error
}
