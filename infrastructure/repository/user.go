package repository

import (
	"github.com/edufriendchen/applet-platform/model"
	"github.com/jmoiron/sqlx"
)

type UserRepository struct {
	conn *sqlx.DB
}

// GetMemberList 获取会员列表
func (mdb *UserRepository) GetMemberList(req model.User) ([]model.User, error) {
	return nil, nil
}

func (mdb *UserRepository) Create(req *model.User) error {
	return nil
}
