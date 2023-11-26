package model

import "github.com/edufriendchen/applet-platform/constant"

type Society struct {
	ID                uint64          `gorm:"primaryKey"`
	Avatar            string          `json:"avatar"`
	Name              string          `json:"name"`
	Type              int             `json:"type"`
	Principal         string          `json:"principal"`
	QualificationFile string          `json:"qualification_file"`
	Status            constant.Status `json:"status"`
	BaseModel
}

func (Society) TableName() string {
	return "society"
}

type SimpleSociety struct {
	ID     uint64 `json:"id"`
	Avatar string `json:"avatar"`
	Name   string `json:"name"`
	Type   int    `json:"type"`
}
