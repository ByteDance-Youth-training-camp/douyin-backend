namespace go data

struct Video{
    1: required i64 id  // 视频唯一标识
    2: required User author // 视频作者信息
    3: required string play_url // 视频播放地址
    4: required string cover_url    // 视频封面地址
    5: required i64 favorite_count  // 视频的点赞总数
    6: required i64 comment_count   // 视频的评论总数
    7: required bool is_favorite    // true-已点赞，false-未点赞
    8: required string title    // 视频标题
}

struct User{
    1: required i64 id  // 用户id
    2: required string name // 用户名称
    3: optional i64 follow_count    // 关注总数
    4: optional i64 follower_count  // 粉丝总数
    5: required bool is_follow  // true-已关注，false-未关注
}

struct Comment{
    1: required i64 id
    2: required User user
    3: required string content
    4: required string create_date // 评论发布日期 格式 mm-dd
}

const i64 MsgTypeSent = 1
const i64 MsgTypeReceived = 2

struct FriendUser{
    1: optional string message
    2: required i64 msgType
}

struct Message{
    1: required i64 id
    2: required i64 to_user_id
    3: required i64 from_user_id
    4: required string content
    5: optional string create_time
}