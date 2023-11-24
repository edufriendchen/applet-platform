package model

import "github.com/edufriendchen/applet-platform/constant"

type Member struct {
	ID        uint64                `json:"id" gorm:"primarykey"`
	AvatarURL string                `json:"avatar_url"`
	Phone     string                `json:"phone"`
	WXOpenID  string                `json:"wx_open_id"`
	DYOpenID  string                `json:"dy_open_id"`
	NickName  string                `json:"nick_name"`
	Name      string                `json:"name"`
	Type      int                   `json:"type"`
	Status    constant.MemberStatus `json:"status"`
	BaseModel
}

type SimpleMember struct {
	Avatar   string `json:"avatar"`
	NickName string `json:"nick_name"`
}
