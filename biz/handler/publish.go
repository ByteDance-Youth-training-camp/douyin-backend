package handler

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
)

type VideoListResponse struct {
	Response
	VideoList []Video `json:"video_list"`
}

func Publish(ctx context.Context, c *app.RequestContext) {
	// TODO: check token
}

func PublishList(ctx context.Context, c *app.RequestContext) {
	// TODO: get video list from db
}
