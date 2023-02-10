package model

type Follow struct {
	ID         int64 `gorm:"primaryKey"`
	UserId     int64 `gorm:"not null"`
	FollowerId int64 `gorm:"not null"`
	Canceled   bool  `gorm:"not null"`
}
