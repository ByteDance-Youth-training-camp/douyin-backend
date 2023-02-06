namespace go socialize
include "api_data.thrift"

struct RelationActionRequest{
    1: required string token
    2: required i64 to_user_id
    3: required i32 action_type
}

struct RelationActionResponse{
    1: required i32 status_code
    2: optional string status_msg
}

struct RelationFollowListRequest{
    1: required i64 user_id
    2: required string token
}

struct RelationFollowListResponse{
    1: required i32 status_code
    2: optional string status_msg
    3: list<api_data.User> user_list
}

struct RelationFollowerListRequest{
    1: required i64 user_id
    2: required string token
}

struct RelationFollowerListResponse{
    1: required i32 status_code
    2: optional string status_msg
    3: list<api_data.User> user_list
}


struct RelationFriendListRequest{
    1: required i64 user_id
    2: required string token
}

struct RelationFriendListResponse{
    1: required i32 status_code
    2: optional string status_msg
    3: list<api_data.FriendUser> user_list
}

struct MessageChatRequest{
    1: required string token
    2: required i64 to_user_id
}

struct MessageChatResponse{
    1: required i32 status_code
    2: optional string status_msg
    3: list<api_data.Message> message_list
}

struct MessageActionRequest{
    1: required string token
    2: required i64 to_user_id
    3: required i32 action_type
    4: required string content
}

struct MessageActionResponse{
    1: required i32 status_code
    2: optional string status_msg
}

service Relation{
    RelationActionResponse RelationAction(RelationActionRequest req) (api.post="/douyin/relation/action/")
    RelationFollowListResponse FollowList(RelationFollowListRequest req) (api.get="/douyin/relation/follow/list/")
    RelationFollowerListResponse FollowerList(RelationFollowerListRequest req) (api.get="/douyin/relation/follower/list/")
    RelationFriendListResponse FriendList(RelationFriendListRequest req) (api.get="/douyin/relation/friend/list/")
}

service Message{
    MessageActionResponse MessageAction(MessageActionRequest req) (api.post="/douyin/message/action/")
    MessageChatResponse MessageChat(MessageChatRequest req) (api.get="/douyin/message/chat/")
}
