package mysql

import "douyin_backend/biz/model"

func FindCommentByID(id int64) (*model.Comment, error) {
	comment := model.Comment{ID: id}
	result := DB.First(&comment)
	if result.Error != nil {
		return nil, result.Error
	}
	return &comment, nil
}

func CreateComment(comment *model.Comment) (*model.Comment, error) {
	result := DB.Create(&comment)
	return comment, result.Error
}

func DeleteCommentByID(id int64) error {
	comment := model.Comment{ID: id}
	result := DB.Delete(&comment)
	return result.Error
}

func GetCommentList(videoId int64, userId int64) ([]*model.Comment, error) {
	var comments []*model.Comment
	result := DB.Where("video_id = ? and user_id = ? and canceled = ?", videoId, userId, false).Find(&comments)
	return comments, result.Error
}
