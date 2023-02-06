package minio

import (
	"log"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"golang.org/x/net/context"
)

var Cli *minio.Client
var ctx = context.Background()
var (
	Endpoint        = "192.168.193.1:9000"
	AccessKeyID     = "dy_access_key"
	SecretAccessKey = "892dcf10-a575-11ed-ab7e-fb82e1d9a6b7"
)

func Init() {
	var err error
	Cli, err = minio.New(Endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(AccessKeyID, SecretAccessKey, ""),
		Secure: false,
	})

	if err != nil {
		log.Fatalln(err)
	}

	InitVideoBucket()
}
