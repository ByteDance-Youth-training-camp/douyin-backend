package minio

import (
	"douyin_backend/biz/config"
	"log"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"golang.org/x/net/context"
)

var Cli *minio.Client
var ctx = context.Background()

func Init() {
	cfg := &config.Cfg.MinIO

	var err error
	Cli, err = minio.New(cfg.Endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(cfg.AccessKeyID, cfg.SecretAccessKey, ""),
		Secure: false,
	})

	if err != nil {
		log.Fatalln(err)
	}

	InitBucket(VideoBucket)
	InitBucket(ImageBucket)
}

func InitBucket(bucket string) {
	ok, err := Cli.BucketExists(ctx, bucket)
	if err != nil {
		log.Fatal(err)
	}
	if !ok {
		err = Cli.MakeBucket(ctx, bucket, minio.MakeBucketOptions{})
		if err != nil {
			log.Fatal(err)
		}
	}

}
