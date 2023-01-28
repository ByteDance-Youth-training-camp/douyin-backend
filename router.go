package main

import (
	handler "douyin_backend/biz/handler"

	"github.com/cloudwego/hertz/pkg/app/server"
)

// customizeRegister registers customize routers.
func customizedRegister(r *server.Hertz) {
	douyin := r.Group("/douyin")

	{
		douyin.GET("/feed", handler.Feed)

		user := douyin.Group("/user")
		{
			user.POST("/register", handler.Register)
			user.POST("/login", handler.Login)
			user.GET("/", handler.UserInfo)
		}

		publish := douyin.Group("/publish")
		{
			publish.POST("/action", handler.Publish)
			publish.GET("/list", handler.PublishList)
		}

		favorite := douyin.Group("/favorite")
		{
			favorite.POST("/action", handler.FavoriteAction)
			favorite.GET("/list", handler.FavoriteList)
		}

		comment := douyin.Group("/comment")
		{
			comment.POST("/action", handler.CommentAction)
			comment.GET("/list", handler.CommentList)
		}

		relation := douyin.Group("/relation")
		{
			relation.POST("/action", handler.RelationAction)
			relation.GET("/follow/list", handler.FollowList)
			relation.GET("/follower/list", handler.FollowerList)
			relation.GET("/friend/list", handler.FriendList)
		}
	}
}
