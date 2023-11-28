package banner

import (
	"github.com/edufriendchen/applet-platform/infrastructure/cache"
)

type Service struct {
	cache cache.CacheStore
}

type IService interface {
}

func NewBannerService(
	cache cache.CacheStore,
) IService {
	return &Service{
		cache: cache,
	}
}
