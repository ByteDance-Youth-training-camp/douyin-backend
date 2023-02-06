package mysql

import (
	"douyin_backend/biz/model"
)

func CreateUser(user *model.User) (*model.User, error) {
	result := DB.Create(&user)
	return user, result.Error
}

func FindUserByName(username string) (*model.User, error) {
	user := model.User{Username: username}
	result := DB.First(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}
