package activity

import (
	"context"
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
	Welfare   string                `json:"welfare"`
	StartTime *time.Time            `json:"start_time"`
	EndTime   *time.Time            `json:"end_time"`
	Able      bool                  `json:"able"`
	VisitNum  int64                 `json:"visit_num"`
}

type DetailResponse struct {
	ID        uint64    `json:"id"`
	PosterUrl string    `json:"poster_url"`
	Type      int       `json:"type"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	Welfare   string    `json:"welfare"`
	StartTime time.Time `json:"start_time"`
	EndTime   time.Time `json:"end_time"`
	Able      bool      `json:"able"`
}

type AbandonRequest struct {
	ID     uint64 `json:"id"`
	Reason string `json:"reason"`
}
