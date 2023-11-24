package society

import (
	"context"
	"mime/multipart"

	"github.com/edufriendchen/applet-platform/infrastructure/cache"
	"github.com/edufriendchen/applet-platform/infrastructure/storage"
	"github.com/edufriendchen/applet-platform/model"
)

type Management struct {
	cache   cache.CacheStore
	storage *storage.QiNiuStorage
}

type ISocietyManagement interface {
	GetSocieties(ctx context.Context, req model.Society) ([]model.Society, error)
	ApplySociety(ctx context.Context, req ApplySocietyRequest) error
}

func NewSocietyManagement(
	cache cache.CacheStore,
	storage *storage.QiNiuStorage,
) ISocietyManagement {
	return &Management{
		cache:   cache,
		storage: storage,
	}
}

type ApplySocietyRequest struct {
	AvatarFile        *multipart.FileHeader `form:"avatar_file"`
	Name              string                `form:"name"`
	Type              int                   `form:"type"`
	Principal         string                `form:"principal"`
	QualificationFile *multipart.FileHeader `form:"qualification_file"`
}
