namespace go core

include "api_data.thrift"

struct FeedRequest{
    1: optional i64 latest_time // 可选参数，限制返回视频的最新投稿时间戳，精确到秒，不填表示当前时间
    2: optional string token // 可选参数，登录用户设置
}

struct FeedResponse{
    1: required i32 status_code // 状态码，0-成功，其他值-失败
    2: optional string status_msg   // 返回状态描述
    3: list<api_data.Video> video_list   // 视频列表
    4: optional i64 next_time   // 本次返回的视频中，发布最早的时间，作为下次请求时的latest_time
}


struct UserRegisterRequest{
    1: required string username // 注册用户名，最长32个字符
    2: required string password // 密码，最长32个字符
}

struct UserRegisterResponse{
    1: required i32 status_code // 状态码，0-成功，其他值-失败
    2: optional string status_msg   // 返回状态描述
    3: required i64 user_id // 用户id
    4: required string token    // 用户鉴权token
}

struct UserLoginRequest{
    1: required string username // 登录用户名
    2: required string password // 登录密码
}

struct UserLoginResponse{
    1: required i32 status_code // 状态码，0-成功，其他值-失败
    2: optional string status_msg   // 返回状态描述
    3: required i64 user_id // 用户id
    4: required string token    // 用户鉴权token
}


struct UserRequest{
    1: required i64 user_id // 用户id
    2: required string token    // 用户鉴权token
}

struct UserResponse{
    1: required i32 status_code // 状态码，0-成功，其他值-失败
    2: optional string status_msg   // 返回状态描述
    3: required api_data.User user   // 用户信息
}

struct PublishActionRequest {
    1: required string token // 用户鉴权token
    // 2: required binary data // 视频数据
    2: required string title // 视频标题
}


struct PublishActionResponse {
    1: required i32 status_code // 状态码，0-成功，其他值-失败
    2: optional string status_msg // 返回状态描述
}

struct PublishListRequest{
    1: required i64 user_id
    2: required string token
}

struct PublishListResponse{
    1: required i32 status_code
    2: optional string status_msg
    3: list<api_data.Video> video_list
}

service Feed{
    FeedResponse Feed(1:FeedRequest req)(api.get="/douyin/feed/")
}
service User{
    UserRegisterResponse UserRegister(1:UserRegisterRequest req) (api.post="/douyin/user/register/")
    UserLoginResponse UserLogin(1:UserLoginRequest req)(api.post="/douyin/user/login/")
    UserResponse UserInfo(1:UserRequest req)(api.get="/douyin/user/")
}

service Publish{
    PublishActionResponse PublishAction(PublishActionRequest req) (api.post="/douyin/publish/action/")
    PublishListResponse PublishList(PublishListRequest req) (api.get="/douyin/publish/list/")
}