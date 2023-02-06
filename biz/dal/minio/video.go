package minio

import (
	"bytes"
	"log"
	"net/url"
	"time"

	"github.com/minio/minio-go/v7"
)

var VideoBucket = "videos"

func InitVideoBucket() {
	ok, err := Cli.BucketExists(ctx, "test-bucket")
	if err != nil {
		log.Fatal(err)
	}
	if !ok {
		err = Cli.MakeBucket(ctx, "test-bucket", minio.MakeBucketOptions{})
		if err != nil {
			log.Fatal(err)
		}
	}

}

func UploadVideo(key string, video []byte) error {
	_, err := Cli.PutObject(ctx, VideoBucket, key, bytes.NewReader(video), int64(len(video)), minio.PutObjectOptions{ContentType: "video/mp4"})
	if err != nil {
		return err
	}
	return nil
}

func GetVideoUrl(key string, expired time.Duration) (url *url.URL, err error) {
	return Cli.PresignedGetObject(ctx, VideoBucket, key, expired, nil)
}
