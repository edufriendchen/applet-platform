package file

import (
	"context"

	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/edufriendchen/applet-platform/model"
)

func (s *Management) UploadFile(ctx context.Context, req model.UploadFileRequest) (*model.UploadFileResponse, error) {
	res, err := s.storage.UploadFile(ctx, req)
	if err != nil {
		hlog.Error("[UploadFile] storage.UploadFile", err)

		return nil, err
	}

	defer func(s *Management, ctx context.Context, key string) {
		err := s.DeleteFile(ctx, key)
		if err != nil {
			hlog.CtxErrorf(ctx, "[UploadFile] defer DeleteFile", err)
		}
	}(s, ctx, res.Key)

	err = s.fileRepository.CreateFileRecord(ctx, &model.File{
		Path: res.Path,
	})
	if err != nil {
		hlog.CtxErrorf(ctx, "[UploadFile] fileRepository.Create", err)

		return nil, err
	}

	return res, err
}

func (s *Management) DeleteFile(ctx context.Context, key string) error {
	err := s.storage.DeleteFile(ctx, key)
	if err != nil {
		hlog.CtxErrorf(ctx, "[DeleteFile] storage.DeleteFile", err)

		return err
	}

	return nil
}
