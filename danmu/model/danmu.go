package model

func (dm *Danmu) Create() error {
	return db.Create(dm).Error
}

func (dm *Danmu) GetAllByVideoID(dms *[]Danmu) error {
	return db.Model(&Danmu{}).Where("video_id = ?", dm.VideoID).Find(dms).Error
}
