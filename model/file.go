package model

import "mime/multipart"

type FileType string

type File struct {
	ID     uint64   `json:"id"`
	Path   string   `json:"path"`
	Bucket string   `json:"bucket"`
	Type   FileType `json:"type"`
	Status int      `json:"status"`
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
