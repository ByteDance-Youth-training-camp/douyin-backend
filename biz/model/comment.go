package model

import "time"

type Comment struct {
	ID        int64     `gorm:"primaryKey"`
	UserId    int64     `gorm:"not null"`
	VideoId   int64     `gorm:"not null"`
	Content   string    `gorm:"type:text; not null"`
	CreatedAt time.Time `gorm:"not null"`
	Canceled  bool      `gorm:"not null"`
}
