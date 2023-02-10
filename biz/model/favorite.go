package model

type Favorite struct {
	ID       int64 `gorm:"primaryKey"`
	UserId   int64 `gorm:"not null"`
	VideoId  int64 `gorm:"not null"`
	Canceled bool  `gorm:"not null"`
}
