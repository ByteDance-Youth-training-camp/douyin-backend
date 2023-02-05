namespace go interact
include "api_data.thrift"

struct FavoriteActionRequest{
    1: required string token
    2: required i64 video_id
    3: required i32 action_type
}

struct FavoriteActionResponse{
    1: required i32 status_code
    2: optional string status_msg
}

struct FavoriteListRequest{
    1: required i64 user_id
    2: required string token
}

struct FavoriteListResponse{
    1: required i32 status_code
    2: optional string status_msg
    3: list<api_data.Video> video_list
}

struct CommentActionRequest{
    1: required string token
    2: required i64 video_id
    3: required i32 action_type
    4: optional string comment_text // 用户填写的评论内容，在action_type=1的时候使用
    5: optional i64 comment_id  // 要删除的评论id，在action_type=2的时候使用
}

struct CommentActionResponse{
    1: required i32 status_code
    2: optional string status_msg
    3: optional api_data.Comment comment
}

struct CommentListRequest{
    1: required string token
    2: required i64 video_id
}

struct CommentListResponse{
    1: required i32 status_code
    2: optional string status_msg
    3: list<api_data.Comment> comment_list
}


service FavoriteService{
    FavoriteActionResponse FavoriteAction(FavoriteActionRequest req) (api.post="/douyin/favorite/action/")
    FavoriteListResponse FavoriteList(FavoriteListRequest req)(api.get="/douyin/favorite/list/")
}

service CommentService{
    CommentActionResponse CommentAction(CommentActionRequest req)(api.post="/douyin/comment/action/")
    CommentListResponse CommentList(CommentListRequest req)(api.get="/douyin/comment/list/")
}