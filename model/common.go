package model

import "time"

type BaseModel struct {
	CreatedAt time.Time  `json:"created_at"`
	CreatedBy uint       `json:"created_by"`
	UpdatedAt *time.Time `json:"updated_at"`
	UpdatedBy *uint      `json:"updated_by"`
}

type Pagination struct {
	PerPage int  `json:"per_page"`
	Page    int  `json:"page"`
	IsDESC  bool `json:"is_desc"`
}

type Status int

const (
	InvalidStatus Status = 1
	ActivityStatus
)
