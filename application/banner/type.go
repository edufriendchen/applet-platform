package banner

import (
	"context"
	"github.com/edufriendchen/applet-platform/infrastructure/cache"
	"github.com/edufriendchen/applet-platform/infrastructure/repository/activity"
	"github.com/edufriendchen/applet-platform/model"
)

type Management struct {
	cache              cache.CacheStore
	activityRepository activity.ActivityRepository
}

type IActivityManagement interface {
	GetActivityList(ctx context.Context, req model.Activity) ([]model.Activity, error)
	GetActivityDetail(ctx context.Context, req model.Activity) ([]model.Activity, error)
	ParticipateActivity(ctx context.Context, req model.Activity) error
	AbandonActivity(ctx context.Context, req model.Activity) error
}

func NewBannerManagement(
	cache cache.CacheStore,
	activityRepository activity.ActivityRepository,
) IActivityManagement {
	return &Management{
		cache:              cache,
		activityRepository: activityRepository,
	}
}
