package video

import (
	"douyin_backend/biz/dal/minio"
	"io"

	"github.com/cloudwego/hertz/pkg/common/hlog"
)

type Video struct {
	FReader  io.Reader
	Size     int64
	Key      string
	CoverKey string
	Vid      int64
	retry    int
}

func UploadVideo(video *Video) {
	upload_que <- video
}

var upload_que = make(chan *Video)

func VideoService() {
	for v := range upload_que {
		err := minio.UploadVideo(v.Key, v.FReader, v.Size)
		if err != nil {
			hlog.Debugf("upload video %s failed: %e\n", v.Key, err)
			v.retry++
			if v.retry < 3 {
				upload_que <- v
			} else {
				hlog.Infof("failed to upload video %s after 3 retries.\n", v.Key)
			}
		}
	}
}
