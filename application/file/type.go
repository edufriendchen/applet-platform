package file

import (
	"context"

	"github.com/edufriendchen/applet-platform/infrastructure/repository"
	"github.com/edufriendchen/applet-platform/infrastructure/storage"
	"github.com/edufriendchen/applet-platform/model"
)

type Management struct {
	fileRepository repository.FileRepository
	storage        *storage.QiNiuStorage
}

type IFileManagement interface {
	UploadFile(ctx context.Context, req model.UploadFileRequest) (*model.UploadFileResponse, error)
}

func NewFileManagement(
	storage *storage.QiNiuStorage,
	fileRepository repository.FileRepository,
) IFileManagement {
	return &Management{
		storage:        storage,
		fileRepository: fileRepository,
	}
}
