package handler

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
)

type UserListResponse struct {
	Response
	UserList []User `json:"user_list"`
}

func RelationAction(ctx context.Context, c *app.RequestContext) {
	// TODO: check relation action
}

func FollowList(ctx context.Context, c *app.RequestContext) {
	// TODO: get follow list from db
}

func FollowerList(ctx context.Context, c *app.RequestContext) {
	// TODO: get follower list from db
}

func FriendList(ctx context.Context, c *app.RequestContext) {
	// TODO: get friend list from db
}
