package repository

import (
	"context"
	"github.com/edufriendchen/applet-platform/model"
	"github.com/jmoiron/sqlx"
)

type UserRepository struct {
	conn *sqlx.DB
}

// GetUserList 获取用户列表
func (mdb *UserRepository) GetUserList(ctx context.Context, req model.User) ([]model.User, error) {
	return nil, nil
}

func (mdb *UserRepository) CreateUser(ctx context.Context, req *model.User) error {
	return nil
}
