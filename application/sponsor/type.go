package sponsor

import (
	"context"

	"github.com/edufriendchen/applet-platform/model"
)

type Management struct {
}

type ISponsorManagement interface {
	CreateActivity(ctx context.Context, req *model.Activity) error
	UpdateActivity(ctx context.Context, req *model.Activity) error
	GetActivityList(ctx context.Context, req model.Activity) ([]model.BaseModel, error)
	GetActivityDetail(ctx context.Context, req model.Activity) ([]model.Activity, error)
	ApplyActivity(ctx context.Context, req model.Activity) error
}
