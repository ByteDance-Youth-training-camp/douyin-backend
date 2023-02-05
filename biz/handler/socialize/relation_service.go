// Code generated by hertz generator.

package socialize

import (
	"context"

	socialize "douyin_backend/biz/hertz_gen/model/socialize"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

// RelationAction .
// @router /douyin/relation/action/ [POST]
func RelationAction(ctx context.Context, c *app.RequestContext) {
	var err error
	var req socialize.RelationActionRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(socialize.RelationActionResponse)

	c.JSON(consts.StatusOK, resp)
}

// FollowList .
// @router /douyin/relation/follow/list/ [GET]
func FollowList(ctx context.Context, c *app.RequestContext) {
	var err error
	var req socialize.RelationFollowListRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(socialize.RelationFollowListResponse)

	c.JSON(consts.StatusOK, resp)
}

// FollowerList .
// @router /douyin/relation/follower/list/ [GET]
func FollowerList(ctx context.Context, c *app.RequestContext) {
	var err error
	var req socialize.RelationFollowerListRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(socialize.RelationFollowerListResponse)

	c.JSON(consts.StatusOK, resp)
}
