package handler

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
)

type UserLoginResponse struct {
	Response
	UserId int64  `json:"user_id,omitempty"`
	Token  string `json:"token"`
}

type UserResponse struct {
	Response
	User User `json:"user"`
}

func Register(ctx context.Context, c *app.RequestContext) {
	// TODO: check username and password
}

func Login(ctx context.Context, c *app.RequestContext) {
	// TODO: check username and password
}

func UserInfo(ctx context.Context, c *app.RequestContext) {
	// TODO: get user info from db
}
