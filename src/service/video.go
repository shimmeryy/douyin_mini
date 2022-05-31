package service

import (
	"context"
	"io/ioutil"
	"mime/multipart"
	"sync"
	"tiktok/src/errno"
	"tiktok/src/utils/ossUtil"
)

type VideoService interface {
	CreateVideo()
}

var (
	videoService     *VideoServiceImpl
	videoServiceOnce sync.Once
)

func VideoServiceInstance() *VideoServiceImpl {
	videoServiceOnce.Do(
		func() {
			videoService = &VideoServiceImpl{}
		})
	return videoService
}

type VideoServiceImpl struct {
}

func (this *VideoServiceImpl) CreateVideo(ctx context.Context, file *multipart.FileHeader) (url string, err error) {
	fileHandle, err := file.Open()
	if err != nil {
		return "", errno.ServiceErr.WithMessage("打开文件错误")
	}
	defer fileHandle.Close()
	fileByte, err := ioutil.ReadAll(fileHandle)
	url, err = ossUtil.UploadFile(file.Filename, fileByte)
	if err != nil {
		return "", errno.ServiceErr.WithMessage("上传失败")
	}
	return url, nil
}
