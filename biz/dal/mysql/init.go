package mysql

import (
	"douyin_backend/biz/config"
	"douyin_backend/biz/model"
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func Init() {

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			IgnoreRecordNotFoundError: true,
		},
	)
	dsn := config.Cfg.MySql.Dsn
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
}
