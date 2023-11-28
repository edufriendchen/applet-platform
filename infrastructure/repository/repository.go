package repository

import (
	"context"
	"github.com/edufriendchen/applet-platform/model"
	"github.com/jmoiron/sqlx"
)

type Repository struct {
	ActivityRepository ActivityRepository
	FileRepository     FileRepository
	UserRepository     UserRepository
}

func NewRepository(conn *sqlx.DB) Repository {
	return Repository{
		ActivityRepository: ActivityRepository{conn},
		FileRepository:     FileRepository{conn},
		UserRepository:     UserRepository{conn},
	}
}

type IActivityRepository interface {
	GetActivityTotal(req *model.Activity) error
	GetActivityList(pagination model.Pagination, req *model.Activity) ([]model.Activity, error)
	CreateActivityRecord(ctx context.Context, req *model.ActivityRecord) error
	UpdateActivityRecord(ctx context.Context, req *model.ActivityRecord) error
}

type IFileRepository interface {
	CreateFileRecord(ctx context.Context, req *model.File) (int64, error)
	DeleteFileRecord(ctx context.Context, req model.File) error
	GetFileRecordList(ctx context.Context, req model.File) ([]model.File, error)
}

type IUserRepository interface {
	GetUserList(ctx context.Context, req model.User) ([]model.User, error)
	CreateUser(ctx context.Context, req *model.User) error
}
