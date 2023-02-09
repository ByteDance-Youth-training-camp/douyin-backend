package model

type Video struct {
	ID         int64 `gorm:"primaryKey"`
	PlayUrl    string
	CoverUrl   string
	UploadTime int64 `gorm:"autoCreateTime; index"`
	UserID     int64
	User       User
}
