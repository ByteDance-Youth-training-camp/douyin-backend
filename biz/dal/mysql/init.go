package mysql

import (
	"douyin_backend/biz/model"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var dsn = "hairline:hairline@tcp(localhost:3306)/douyin?charset=utf8&parseTime=True&loc=Local"

var DB *gorm.DB

func Init() {
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		SkipDefaultTransaction: true,
		PrepareStmt:            true,
	})
	if err != nil {
		panic(err)
	}
	migrate(DB)
}

func migrate(db *gorm.DB){
	db.AutoMigrate(&model.User{})
}
