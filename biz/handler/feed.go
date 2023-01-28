package handler

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
)

type FeedResponse struct {
	Response
	VideoList []Video `json:"video_list,omitempty"`
	NextTime  int64   `json:"next_time,omitempty"`
}

func Feed(ctx context.Context, c *app.RequestContext) {
	// TODO: get video list from db
}
