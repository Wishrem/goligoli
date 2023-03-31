package model

type Video struct {
	ID          int64  `gorm:"primaryKey;not null;unique"`
	UserID      int64  `gorm:"not null"`
	Title       string `gorm:"unique;not null"`
	Description string `gorm:"not null"`
	Liked       int64  `gorm:"not null"`
	Shared      int64  `gorm:"not null"`
	VideoUrl    string `gorm:"not null"`
	Status      Status `gorm:"foreignKey:VideoID;references:ID"`
}

type Status struct {
	VideoID int64 `gorm:"primaryKey;not null;unique;"`
	Passed  bool  `gorm:"not null"`
	Reason  string
}
