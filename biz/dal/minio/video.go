package minio

import (
	"io"
	"log"
	"net/url"
	"time"

	"github.com/minio/minio-go/v7"
)

var VideoBucket = "videos"

func InitVideoBucket() {
	ok, err := Cli.BucketExists(ctx, VideoBucket)
	if err != nil {
		log.Fatal(err)
	}
	if !ok {
		err = Cli.MakeBucket(ctx, VideoBucket, minio.MakeBucketOptions{})
		if err != nil {
			log.Fatal(err)
		}
	}

}

func UploadVideo(key string, video io.Reader, size int64) error {

	_, err := Cli.PutObject(ctx, VideoBucket, key, video, size, minio.PutObjectOptions{ContentType: "video/mp4"})
	if err != nil {
		return err
	}
	return nil
}

func GetVideoUrl(key string, expired time.Duration) (url *url.URL, err error) {
	return Cli.PresignedGetObject(ctx, VideoBucket, key, expired, nil)
}
