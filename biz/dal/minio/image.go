package minio

import (
	"io"
	"net/url"
	"time"

	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/minio/minio-go/v7"
)

var ImageBucket = "videocovers"

func UploadImage(key string, img io.Reader, size int64) error {

	_, err := Cli.PutObject(ctx, ImageBucket, key, img, size, minio.PutObjectOptions{ContentType: "image/jpeg"})
	if err != nil {
		hlog.Error(err)
		return err
	}
	return nil
}

func GetImageUrl(key string, expired time.Duration) (url *url.URL, err error) {

	return Cli.PresignedGetObject(ctx, ImageBucket, key, expired, nil)
}
