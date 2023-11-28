package repository

import (
	"context"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/edufriendchen/applet-platform/model"
	"github.com/jmoiron/sqlx"
	"strconv"
)

const (
	createFileRecordQuery = `
		INSERT INTO uploaded_file
		    (path, bucket, type, status) 
		VALUES (?, ?, ?, 1);
	`

	deleteFileRecordQuery = `
		UPDATE uploaded_file
		SET status = 0  
		WHERE id = ?
	`

	getFileRecordQuery = `
		SELECT id, path, bucket, type, status,  created_at, updated_at 
		FROM uploaded_file
		WHERE TRUE 
	`
)

type FileRepository struct {
	conn *sqlx.DB
}

func (mdb *FileRepository) CreateFileRecord(ctx context.Context, req *model.File) error {
	res, err := mdb.conn.ExecContext(
		ctx,
		createFileRecordQuery,
		req.Path,
		req.Bucket,
		req.Type,
	)

	if err != nil {
		hlog.CtxErrorf(ctx, "[CreateActivityRecord] failed create product setup", err)

		return err
	}

	// get last insert ID
	lastID, err := res.LastInsertId()
	if err != nil {
		hlog.CtxErrorf(ctx, "[CreateActivityRecord] failed get last insert ID", err)

		return err
	}
	req.ID = uint64(lastID)

	return nil
}

func (mdb *FileRepository) DeleteFileRecord(ctx context.Context, req model.File) error {
	res, err := mdb.conn.ExecContext(
		ctx,
		deleteFileRecordQuery,
		req.ID,
	)

	if err != nil {
		hlog.CtxErrorf(ctx, "[CreateActivityRecord] failed create product setup", err)

		return err
	}

	// get last insert ID
	lastID, err := res.LastInsertId()
	if err != nil {
		hlog.CtxErrorf(ctx, "[CreateActivityRecord] failed get last insert ID", err)

		return err
	}
	req.ID = uint64(lastID)

	return nil
}

func (mdb *FileRepository) GetFileRecordList(ctx context.Context, req model.File) ([]model.File, error) {
	var data []model.File
	var item model.File

	var sqlStatement = getFileRecordQuery
	var filterQuery string
	queryArgs := make([]interface{}, 0)

	if req.ID > 0 {
		filterQuery += " AND id = ? "
		queryArgs = append(queryArgs, strconv.FormatUint(req.ID, 10))
	}

	if req.Type != "" {
		filterQuery += " AND type = ? "
		queryArgs = append(queryArgs, req.Type)
	}

	if req.Status > 0 {
		filterQuery += " AND status = ? "
		queryArgs = append(queryArgs, req.Status)
	}

	sqlStatement += filterQuery
	qry, err := mdb.conn.Query(sqlStatement, queryArgs...)
	if err != nil {
		hlog.CtxErrorf(ctx, "[GetFileRecordList] db execute", err)

		return data, err
	}
	defer qry.Close()

	for qry.Next() {
		err = qry.Scan(
			&item.ID,
			&item.Path,
			&item.Bucket,
			&item.Type,
			&item.Status,
			&item.CreatedAt,
			&item.UpdatedAt,
		)
		if err != nil {
			hlog.CtxErrorf(ctx, "[GetFileRecordList] db scan", err)

			return data, err
		}
		data = append(data, item)
	}

	return data, nil
}
