package society

import (
	"context"

	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/edufriendchen/applet-platform/model"
)

func (s *Management) GetSocieties(ctx context.Context, req model.Society) ([]model.Society, error) {
	return nil, nil
}

func (s *Management) ApplySociety(ctx context.Context, req ApplySocietyRequest) error {
	var (
		err  error
		keys = make([]string, 2)
	)

	defer func() {
		if err != nil {
			for _, key := range keys {
				if err := s.storage.DeleteFile(ctx, key); err != nil {
					hlog.Error("ApplySociety defer DeleteFile err", err)
				}
			}
		}
	}()

	avatarFileRes, err := s.storage.UploadFile(ctx, model.UploadFileRequest{
		File: req.AvatarFile,
	})
	if err != nil {
		hlog.Error("ApplySociety storage.UploadFile err", err)

		return err
	}
	keys[1] = avatarFileRes.Key

	qualificationFileRes, err := s.storage.UploadFile(ctx, model.UploadFileRequest{
		File: req.QualificationFile,
	})
	if err != nil {
		hlog.Error("ApplySociety storage.UploadFile err", err)

		return err
	}
	keys[1] = qualificationFileRes.Key

	return nil
}
