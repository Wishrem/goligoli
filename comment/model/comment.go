package model

import "gorm.io/gorm/clause"

func (cm *Comment) CreateComment() error {
	return db.Create(cm).Error
}

func (rp *Response) CreateResponse() error {
	return db.Create(rp).Error
}

func (cm *Comment) GetCommentAndResponds() error {
	return db.Model(&Comment{}).Preload(clause.Associations).Take(cm).Error
}
