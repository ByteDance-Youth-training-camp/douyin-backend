package handler

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
)

type CommentListResponse struct {
	Response
	CommentList []Comment `json:"comment_list,omitempty"`
}

type CommentActionResponse struct {
	Response
	Comment Comment `json:"comment,omitempty"`
}

func CommentAction(ctx context.Context, c *app.RequestContext) {
	// TODO: check comment action
}

func CommentList(ctx context.Context, c *app.RequestContext) {
	// TODO: get comment list from db
}
