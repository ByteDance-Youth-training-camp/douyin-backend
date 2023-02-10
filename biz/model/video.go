package model

import "time"

type Video struct {
	ID          int64     `gorm:"primaryKey"`
	UserId      int64     `gorm:"not null"`
	Title       string    `gorm:"type:varchar(255); not null"`
	Url         string    `gorm:"type:varchar(255); not null"`
	Thumb       string    `gorm:"type:varchar(255); not null"`
	PublishedAt time.Time `gorm:"not null"`
}
