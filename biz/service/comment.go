package service

import (
	"douyin_backend/biz/dal/mysql"
	"douyin_backend/biz/dal/redis"
	"douyin_backend/biz/hertz_gen/model/data"
	"douyin_backend/biz/model"
	"douyin_backend/biz/service/userservice"
	"time"
)

type CommentService struct {
}

func (s *CommentService) CreateComment(userId int64, videoId int64, content *string) (data.Comment, error) {
	// delete cache, update db
	redis.DeleteVideoCommentsCnt(videoId)
	comment, err := mysql.CreateComment(&model.Comment{
		UserId:    userId,
		VideoId:   videoId,
		Content:   *content,
		CreatedAt: time.Now(),
		Canceled:  false,
	})
	if err != nil {
		return data.Comment{}, err
	}

	user, err := userservice.GetUserById(userId)

	if err != nil {
		return data.Comment{}, err
	}

	return data.Comment{
		ID:         comment.ID,
		User:       &user,
		CreateDate: comment.CreatedAt.Format("2006-01-02 15:04:05"),
		Content:    comment.Content,
	}, nil
}

func (s *CommentService) DeleteComment(commentId int64) error {
	return mysql.DeleteCommentByID(commentId)
}

func (s *CommentService) GetCommentList(videoId int64, userId int64) ([]*data.Comment, error) {
	comments, err := mysql.GetCommentList(videoId, userId)

	if err != nil {
		return nil, err
	}

	var commentList []*data.Comment

	for _, comment := range comments {
		user, err := userservice.GetUserById(comment.UserId)

		if err != nil {
			return nil, err
		}

		commentList = append(commentList, &data.Comment{
			ID:         comment.ID,
			User:       &user,
			CreateDate: comment.CreatedAt.Format("2006-01-02 15:04:05"),
			Content:    comment.Content,
		})
	}

	return commentList, nil
}
