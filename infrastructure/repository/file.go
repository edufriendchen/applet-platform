package repository

import (
	"github.com/edufriendchen/applet-platform/model"
	"github.com/jmoiron/sqlx"
)

type FileRepository struct {
	conn *sqlx.DB
}

func (db *FileRepository) Create(req model.File) (int64, error) {
	return 0, nil
}

func (db *FileRepository) Delete(req model.File) error {
	return nil
}

func (db *FileRepository) Update(req model.File) (int64, error) {
	return 0, nil
}

func (db *FileRepository) Get(req model.File) ([]model.File, error) {
	return nil, nil
}
