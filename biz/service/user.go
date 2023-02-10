package service

import (
	"douyin_backend/biz/dal/mysql"
	"douyin_backend/biz/hertz_gen/model/data"
	"douyin_backend/biz/model"
	"errors"

	"github.com/cloudwego/hertz/pkg/common/hlog"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
}

func (s *UserService) GetUserById(id int64) (data.User, error) {
	user, err := mysql.UserInfoByID(id)
	if err != nil {
		return data.User{}, err
	}
	return data.User{
		ID:   user.ID,
		Name: user.Username,
	}, nil
}

func (s *UserService) GetUserByName(name string) (model.User, error) {
	user, err := mysql.FindUserByName(name)
	if err != nil {
		return model.User{}, err
	}
	return *user, nil
}

func (s *UserService) CreateUser(user *model.User) (*data.User, error) {
	ulen, plen := len(user.Username), len(user.Password)
	if ulen > 32 || plen > 32 || ulen < 1 || plen <= 5 {
		return nil, errors.New("invalid username or password length")
	}
	// Check if username exist in database
	if user, _ := mysql.FindUserByName(user.Username); user != nil {
		return nil, errors.New("username already exists")
	}

	// generate hash for password
	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		hlog.DefaultLogger().Debug(err)
		return nil, errors.New("internal error")
	}

	user, err = mysql.CreateUser(&model.User{Username: user.Username, Password: string(hash)})
	if err != nil {
		hlog.DefaultLogger().Debug(err)
		return nil, errors.New("internal error")
	}

	return &data.User{
		ID:   user.ID,
		Name: user.Username,
	}, nil
}
