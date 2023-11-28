package model

import "github.com/edufriendchen/applet-platform/constant"

type User struct {
	ID        uint64                `json:"id"`
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

type SimpleUser struct {
	Avatar   string `json:"avatar_url"`
	NickName string `json:"nick_name"`
}
