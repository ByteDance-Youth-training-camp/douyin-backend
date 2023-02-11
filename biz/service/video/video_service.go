package video

import (
	"douyin_backend/biz/dal/minio"
	"douyin_backend/biz/dal/mysql"
	"fmt"
	"image/jpeg"
	"io"

	"github.com/bakape/thumbnailer/v2"
	"github.com/cloudwego/hertz/pkg/common/hlog"
)

type Video struct {
	FReader io.Reader
	Size    int64
	Vid     int64
}

func init() {
	go coverGenerator()
}

type coverGenTask struct {
	vid  int64
	vkey string
}

var coverQueSize = 1024
var coverQue = make(chan coverGenTask, coverQueSize)

func UploadVideo(video *Video) error {
	key := GenVKey(video.Vid)

	err := minio.UploadVideo(key, video.FReader, video.Size)
	if err != nil {
		return err
	}
	coverQue <- coverGenTask{video.Vid, key}
	return nil
}

func GenVKey(vid int64) string {
	return fmt.Sprintf("video%d", vid)
}

func coverGenerator() {
	hlog.Info("Start cover generator ")
	for task := range coverQue {
		func() {
			obj, err := minio.GetVideoObject(task.vkey)
			if err != nil {
				hlog.Error(err)
				return
			}
			defer obj.Close()

			_, thumb, err := thumbnailer.Process(obj, thumbnailer.Options{
				// 0,0 => default size 150 by 150
				ThumbDims: thumbnailer.Dims{Width: 320, Height: 320},
			})
			if err != nil && err != io.EOF {
				hlog.Error(err)
				return
			}

			imgReader, imgWriter := io.Pipe()
			defer imgReader.Close()
			go func() {
				defer imgWriter.Close()
				if err := jpeg.Encode(imgWriter, thumb, nil); err != nil {
					hlog.Error(err)
				}
			}()
			// use same key among different buckets
			imgKey := task.vkey

			if err := minio.UploadImage(imgKey, imgReader, -1); err != nil {
				hlog.Error(err)
				return
			}

			// update database
			if err := mysql.UpdateVideoKeys(task.vid, task.vkey, imgKey); err != nil {
				hlog.Error(err)
				return
			}
		}()

	}
}
