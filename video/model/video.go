package model

import (
	"reflect"
	"strings"

	"github.com/wishrem/goligoli/erp"
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

func (s *Search) SearchVideos(dst *[]Video) error {
	sql := make([]string, 0)
	params := make([]interface{}, 0)
	sT := reflect.TypeOf(*s)
	sV := reflect.ValueOf(*s)
	for i := 0; i < sT.NumField(); i++ {
		if tag := sT.Field(i).Tag.Get("sql"); tag != "" && !sV.Field(i).IsZero() {
			sql = append(sql, tag)
			v := sV.Field(i)
			if !v.CanInterface() {
				return erp.New(erp.INTERNAL_ERROR, "the filed of 'Search': %v can't be used as 'Interface'", v)
			}
			params = append(params, v.Interface())
		}
	}
	if len(sql) == 0 {
		return erp.New(erp.INTERNAL_ERROR, "video searching option is empty")
	}

	return db.Model(&Video{}).Preload(clause.Associations).Where(strings.Join(sql, " and "), params...).Find(dst).Error
}
