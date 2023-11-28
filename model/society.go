package model

import "github.com/edufriendchen/applet-platform/constant"

type Society struct {
	ID                uint64          `json:"id"`
	Avatar            string          `json:"avatar"`
	Name              string          `json:"name"`
	Type              int             `json:"type"`
	Principal         string          `json:"principal"`
	QualificationFile string          `json:"qualification_file"`
	Status            constant.Status `json:"status"`
	BaseModel
}

type SimpleSociety struct {
	ID     uint64 `json:"id"`
	Avatar string `json:"avatar"`
	Name   string `json:"name"`
	Type   int    `json:"type"`
}
