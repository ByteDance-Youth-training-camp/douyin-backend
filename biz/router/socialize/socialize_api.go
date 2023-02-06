// Code generated by hertz generator. DO NOT EDIT.

package Socialize

import (
	socialize "douyin_backend/biz/handler/socialize"
	"github.com/cloudwego/hertz/pkg/app/server"
)

/*
 This file will register all the routes of the services in the master idl.
 And it will update automatically when you use the "update" command for the idl.
 So don't modify the contents of the file, or your code will be deleted when it is updated.
*/

// Register register routes based on the IDL 'api.${HTTP Method}' annotation.
func Register(r *server.Hertz) {

	root := r.Group("/", rootMw()...)
	{
		_douyin := root.Group("/douyin", _douyinMw()...)
		{
			_message := _douyin.Group("/message", _messageMw()...)
			{
				_action := _message.Group("/action", _actionMw()...)
				_action.POST("/", append(_message_ctionMw(), socialize.MessageAction)...)
			}
			{
				_chat := _message.Group("/chat", _chatMw()...)
				_chat.GET("/", append(_messagechatMw(), socialize.MessageChat)...)
			}
		}
		{
			_relation := _douyin.Group("/relation", _relationMw()...)
			{
				_action0 := _relation.Group("/action", _action0Mw()...)
				_action0.POST("/", append(_relation_ctionMw(), socialize.RelationAction)...)
			}
			{
				_follow := _relation.Group("/follow", _followMw()...)
				{
					_list := _follow.Group("/list", _listMw()...)
					_list.GET("/", append(_followlistMw(), socialize.FollowList)...)
				}
			}
			{
				_follower := _relation.Group("/follower", _followerMw()...)
				{
					_list0 := _follower.Group("/list", _list0Mw()...)
					_list0.GET("/", append(_followerlistMw(), socialize.FollowerList)...)
				}
			}
			{
				_friend := _relation.Group("/friend", _friendMw()...)
				{
					_list1 := _friend.Group("/list", _list1Mw()...)
					_list1.GET("/", append(_friendlistMw(), socialize.FriendList)...)
				}
			}
		}
	}
}
