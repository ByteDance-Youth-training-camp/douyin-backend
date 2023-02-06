package service

import (
	"douyin_backend/biz/dal/mysql"
	"douyin_backend/biz/hertz_gen/model/core"
	"douyin_backend/biz/model"

	"golang.org/x/crypto/bcrypt"
)

var (
	msgInvalidLength = "The length of Username or Password should be in range [5, 32]"
	msgInternalError = "Internal Error"
	msgDbError       = "Database Error"
	msgUserExist     = "Username already exists!"
)

func RegisterUser(req *core.UserRegisterRequest, resp *core.UserRegisterResponse) {
	ulen, plen := len(req.Username), len(req.Password)
	if ulen > 32 || plen > 32 || ulen < 1 || plen <= 5 {
		resp.StatusCode = -1
		resp.StatusMsg = &msgInvalidLength
		return
	}
	// Check if username exist in database
	if user, _ := mysql.FindUserByName(req.Username); user != nil {
		resp.StatusCode = -1
		resp.StatusMsg = &msgUserExist
		return
	}

	// generate hash for password
	hash, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		resp.StatusCode = -1
		resp.StatusMsg = &msgInternalError
		return
	}

	user, err := mysql.CreateUser(&model.User{Username: req.Username, Password: string(hash)})
	if err != nil {
		resp.StatusCode = -1
		resp.StatusMsg = &msgDbError
		return
	}
	resp.UserID = user.ID

	// TODO( generate token )

	return
}
