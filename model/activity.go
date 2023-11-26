package model

import (
	"github.com/edufriendchen/applet-platform/constant"
	"time"
)

type Activity struct {
	ID        uint64                `json:"id" gorm:"primaryKey"`
	Title     string                `json:"title"`
	PosterURL string                `json:"poster_url"`
	Content   string                `json:"content"`
	Welfare   string                `json:"welfare"`
	Type      constant.ActivityType `json:"type"`
	StartTime *time.Time            `json:"start_time"`
	EndTime   *time.Time            `json:"end_time"`
	Status    constant.Status       `json:"status"`
	BaseModel
}

type ActivityRecord struct {
	ID            uint64          `json:"id"`
	ActivityID    uint64          `json:"activity_id"`
	ParticipantID uint64          `json:"participant_id"`
	Type          int             `json:"type"`
	Link          string          `json:"link"`
	Note          string          `json:"note"`
	Status        constant.Status `json:"status"`
	BaseModel
}
