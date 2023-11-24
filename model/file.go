package model

import "mime/multipart"

type File struct {
	ID     uint64 `gorm:"primaryKey"`
	Path   string `json:"path"`
	Bucket string `json:"bucket"`
	Type   string `json:"type"`
	Status int    `json:"status"`
	BaseModel
}

type UploadFileRequest struct {
	File  *multipart.FileHeader `form:"file"`
	Extra map[string]string
}

type UploadFileResponse struct {
	Key  string
	Path string
	Size int64
}
