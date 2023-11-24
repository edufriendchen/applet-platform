package repository

import (
	"github.com/edufriendchen/applet-platform/model"
	"github.com/jmoiron/sqlx"
)

type MemberRepository struct {
	conn *sqlx.DB
}

// GetMemberList 获取会员列表
func (mdb *MemberRepository) GetMemberList(req model.Member) ([]model.Member, error) {
	return nil, nil
}

func (mdb *MemberRepository) Create(req *model.Member) error {
	return nil
}
