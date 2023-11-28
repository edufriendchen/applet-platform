package activity

import (
	"context"
	"github.com/edufriendchen/applet-platform/model"
	"time"

	"github.com/edufriendchen/applet-platform/constant"
	"github.com/edufriendchen/applet-platform/infrastructure/cache"
	"github.com/edufriendchen/applet-platform/infrastructure/repository"
)

type Service struct {
	cache              cache.CacheStore
	activityRepository repository.ActivityRepository
}

type IActivityService interface {
	GetActivityList(ctx context.Context, req Request) ([]Response, error)
	GetActivityDetail(ctx context.Context, id uint64) (DetailResponse, error)
	ParticipateActivity(ctx context.Context, id uint64) error
	AbandonActivity(ctx context.Context, req AbandonRequest) error
}

func NewActivityService(
	cache cache.CacheStore,
	activityRepository repository.ActivityRepository,
) IActivityService {
	return &Service{
		cache:              cache,
		activityRepository: activityRepository,
	}
}

type Request struct {
	ID        uint64                `json:"id"`
	PerPage   int                   `json:"per_page"`
	Page      int                   `json:"page"`
	Type      constant.ActivityType `json:"type"`
	StartTime *time.Time            `json:"start_time"`
	EndTime   *time.Time            `json:"end_time"`
	Able      bool                  `json:"able"`
}

type Response struct {
	ID        uint64                `json:"id"`
	PosterUrl string                `json:"poster_url"`
	Title     string                `json:"title"`
	Type      constant.ActivityType `json:"type"`
	StartTime *time.Time            `json:"start_time"`
	EndTime   *time.Time            `json:"end_time"`
	Status    constant.Status       `json:"able"`
	VisitNum  int64                 `json:"visit_num"`
}

type DetailResponse struct {
	ID   uint64 `json:"id"`
	Able bool   `json:"able"`
	List []model.SimpleUser
}

type AbandonRequest struct {
	ID     uint64 `json:"id"`
	Reason string `json:"reason"`
}

type SubmitRequest struct {
	ID     uint64 `json:"id"`
	Submit string `json:"submit"`
}
