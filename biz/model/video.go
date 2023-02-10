package model

type Video struct {
	ID         int64 `gorm:"primaryKey"`
	Title      string
	VideoKey   string
	CoverKey   string
	Ready      bool
	UploadTime int64 `gorm:"autoCreateTime; index"`
	UserID     int64
	User       User
}
