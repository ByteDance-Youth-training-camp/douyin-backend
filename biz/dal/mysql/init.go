package mysql

import (
	"douyin_backend/biz/model"
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var dsn = "hairline:hairline@tcp(localhost:33306)/douyin?charset=utf8&parseTime=True&loc=Local"

var DB *gorm.DB

func Init() {

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			IgnoreRecordNotFoundError: true,
		},
	)

	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		SkipDefaultTransaction: true,
		PrepareStmt:            true,
		Logger:                 newLogger,
	})
	if err != nil {
		panic(err)
	}
	migrate(DB)
}

func migrate(db *gorm.DB) {
	db.AutoMigrate(&model.User{})
	db.AutoMigrate(&model.Video{})
	db.AutoMigrate(&model.Comment{})
	db.AutoMigrate(&model.Favorite{})
	db.AutoMigrate(&model.Follow{})
}
