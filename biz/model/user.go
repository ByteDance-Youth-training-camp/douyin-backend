package model

type User struct {
	ID       int64  `gorm:"primaryKey"`
	Username string `gorm:"size:32; unique; not null"`
	Password string `gorm:"not null"`
}
