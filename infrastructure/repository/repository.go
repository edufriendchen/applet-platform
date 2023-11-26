package repository

import (
	"context"
	"github.com/edufriendchen/applet-platform/model"
	"github.com/jmoiron/sqlx"
)

type Repository struct {
	ActivityRepository ActivityRepository
	FileRepository     FileRepository
	MemberRepository   MemberRepository
}

func NewRepository(conn *sqlx.DB) Repository {
	return Repository{
		ActivityRepository: ActivityRepository{conn},
		FileRepository:     FileRepository{conn},
		MemberRepository:   MemberRepository{conn},
	}
}

type IActivityRepository interface {
	GetActivityTotal(req *model.Activity) error
	GetActivityList(pagination model.Pagination, req *model.Activity) ([]model.Activity, error)
	CreateActivityRecord(ctx context.Context, req *model.ActivityRecord) error
	UpdateActivityRecord(ctx context.Context, req *model.ActivityRecord) error
}

type IFileRepository interface {
	Create(req model.File) (int64, error)
	Delete(req model.File) error
	Update(req model.File) (int64, error)
	Get(req model.File) ([]model.File, error)
}

type IMemberRepository interface {
	GetMemberList(req model.Member) ([]model.Member, error)
	Create(req *model.Member) error
}
