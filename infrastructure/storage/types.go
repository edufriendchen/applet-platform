package storage

import (
	"context"

	"github.com/edufriendchen/applet-platform/model"
	"github.com/qiniu/api.v7/v7/auth/qbox"
)

type QiNiuStorage struct {
	mac *qbox.Mac

	UseHTTPS      bool
	UseCdnDomains bool
	Zone          string
	BaseURI       string
	bucket        string
	AccessKey     string
	SecretKey     string
}

type Storage interface {
	UploadFile(ctx context.Context, req model.UploadFileRequest) (*model.UploadFileResponse, error)
	DeleteFile(ctx context.Context, key string) error
}

// NewStorage initialize connection for redis and store it into redis pool
func NewStorage(accessKey string, secretKey string) (*QiNiuStorage, error) {
	mac := qbox.NewMac(accessKey, secretKey)
	return &QiNiuStorage{
		mac:    mac,
		bucket: "society-platform",
	}, nil
}
