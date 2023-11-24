package activity

import (
	"context"
	"time"

	"github.com/edufriendchen/applet-platform/constant"
	"github.com/edufriendchen/applet-platform/infrastructure/cache"
	"github.com/edufriendchen/applet-platform/infrastructure/repository"
	"github.com/edufriendchen/applet-platform/model"
)

type Management struct {
	cache              cache.CacheStore
	activityRepository repository.ActivityRepository
}

type IActivityManagement interface {
	GetActivityList(ctx context.Context, req ListActivityRequest) ([]ListActivityResponse, error)
	GetActivityDetail(ctx context.Context, req model.Activity) ([]ListActivityResponse, error)
	ParticipateActivity(ctx context.Context, req model.Activity) error
	AbandonActivity(ctx context.Context, req model.Activity) error
}

func NewActivityManagement(
	cache cache.CacheStore,
	activityRepository repository.ActivityRepository,
) IActivityManagement {
	return &Management{
		cache:              cache,
		activityRepository: activityRepository,
	}
}

type ListActivityRequest struct {
	ID        uint64                `json:"id"`
	PerPage   int                   `json:"per_page"`
	Page      int                   `json:"page"`
	Type      constant.ActivityType `json:"type"`
	StartTime time.Time             `json:"start_time"`
	EndTime   time.Time             `json:"end_time"`
	Able      bool                  `json:"able"`
}

type ListActivityResponse struct {
	Id        uint64    `json:"id"`
	PosterUrl string    `json:"poster_url"`
	Title     string    `json:"title"`
	Type      int       `json:"type"`
	Welfare   string    `json:"welfare"`
	StartTime time.Time `json:"start_time"`
	EndTime   time.Time `json:"end_time"`
	Able      bool      `json:"able"`
	VisitNum  int64     `json:"visit_num"`
}

type DetailResponse struct {
	Id        uint      `json:"id"`
	PosterUrl string    `json:"poster_url"`
	Type      int       `json:"type"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	Welfare   string    `json:"welfare"`
	StartTime time.Time `json:"start_time"`
	EndTime   time.Time `json:"end_time"`
	Able      bool      `json:"able"`
}
