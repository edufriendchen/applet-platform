package storage

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/edufriendchen/applet-platform/model"
	"github.com/qiniu/api.v7/v7/auth/qbox"
	"github.com/qiniu/api.v7/v7/storage"
)

func (s *QiNiuStorage) UploadFile(ctx context.Context, req model.UploadFileRequest) (*model.UploadFileResponse, error) {
	putPolicy := storage.PutPolicy{Scope: s.bucket}
	upToken := putPolicy.UploadToken(s.mac)
	cfg := s.getConfig()
	formUploader := storage.NewFormUploader(cfg)
	ret := storage.PutRet{}
	var putExtra storage.PutExtra

	if req.Extra != nil {
		putExtra = storage.PutExtra{Params: req.Extra}
	}

	f, err := req.File.Open()
	if err != nil {
		hlog.Error("QiNiuStorage UploadFile file.Open() Filed ", err)

		return nil, errors.New("function file.Open() Filed, err:" + err.Error())
	}

	defer f.Close()
	fileKey := fmt.Sprintf("%d%s", time.Now().Unix(), req.File.Filename)
	err = formUploader.Put(ctx, &ret, upToken, fileKey, f, req.File.Size, &putExtra)
	if err != nil {
		hlog.Error("QiNiuStorage formUploader.Put() Filed ", err)

		return nil, errors.New("function formUploader.Put() Filed, err:" + err.Error())
	}

	return &model.UploadFileResponse{
		Key:  ret.Key,
		Path: s.BaseURI + "/" + ret.Key,
		Size: req.File.Size,
	}, nil
}

func (s *QiNiuStorage) DeleteFile(ctx context.Context, key string) error {
	mac := qbox.NewMac(s.AccessKey, s.SecretKey)
	cfg := s.getConfig()
	bucketManager := storage.NewBucketManager(mac, cfg)
	if err := bucketManager.Delete(s.bucket, key); err != nil {
		hlog.Error("QiNiuStorage formUploader.Put() Filed ", err)

		return errors.New("function bucketManager.Delete() Filed, err:" + err.Error())
	}
	return nil
}

func (s *QiNiuStorage) getConfig() *storage.Config {
	cfg := storage.Config{
		UseHTTPS:      s.UseHTTPS,
		UseCdnDomains: s.UseCdnDomains,
	}
	switch s.Zone {
	case "ZoneHuadong":
		cfg.Zone = &storage.ZoneHuadong
	case "ZoneHuabei":
		cfg.Zone = &storage.ZoneHuabei
	case "ZoneHuanan":
		cfg.Zone = &storage.ZoneHuanan
	case "ZoneBeimei":
		cfg.Zone = &storage.ZoneBeimei
	case "ZoneXinjiapo":
		cfg.Zone = &storage.ZoneXinjiapo
	}
	return &cfg
}
