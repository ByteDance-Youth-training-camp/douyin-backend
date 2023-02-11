package minio

import (
	"io"
	"net/url"
	"time"

	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/minio/minio-go/v7"
)

var VideoBucket = "videos"

func GetVideoObject(key string) (*minio.Object, error) {
	return Cli.GetObject(ctx, VideoBucket, key, minio.GetObjectOptions{})
}

func UploadVideo(key string, video io.Reader, size int64) error {

	_, err := Cli.PutObject(ctx, VideoBucket, key, video, size, minio.PutObjectOptions{ContentType: "video/mp4"})
	if err != nil {
		hlog.Error(err)
		return err
	}
	hlog.Info("uploaded video with key ", key)
	return nil
}

func GetVideoUrl(key string, expired time.Duration) (url *url.URL, err error) {
	return Cli.PresignedGetObject(ctx, VideoBucket, key, expired, nil)
}
