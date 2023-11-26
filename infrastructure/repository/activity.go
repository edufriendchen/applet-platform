package repository

import (
	"context"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/edufriendchen/applet-platform/constant"
	"github.com/edufriendchen/applet-platform/model"
	"github.com/jmoiron/sqlx"
	"strconv"
)

const (
	getActivityListQuery = `
		SELECT id, title, poster_url, content, welfare, type, start_time, end_time, status, created_at, updated_at 
		FROM activity 
		WHERE true
	`

	createActivityRecordQuery = `
		INSERT INTO activity_record
		    (activity_id, participants, type, status) 
		VALUES (?, ?, ?, 1);
	`

	updateActivityRecordQuery = `
		UPDATE activity_record
		SET submit = ?, 
			note = ?, 
			status = ? 
		WHERE id = ?
	`
)

type ActivityRepository struct {
	conn *sqlx.DB
}

// GetActivityTotal 统计活动数量
func (mdb *ActivityRepository) GetActivityTotal(req *model.Activity) error {
	return nil
}

// GetActivityList 获取活动列表
func (mdb *ActivityRepository) GetActivityList(pagination model.Pagination, req *model.Activity) ([]model.Activity, error) {
	var data []model.Activity
	var dataActivity model.Activity

	var sqlStatement = getActivityListQuery
	var filterQuery string
	queryArgs := make([]interface{}, 0)

	if req.ID > 0 {
		filterQuery += " AND id = ? "
		queryArgs = append(queryArgs, strconv.FormatUint(req.ID, 10))
	}
	if req.Type > 0 {
		filterQuery += " AND type = ? "
		queryArgs = append(queryArgs, req.Type)
	}
	if req.Status > 0 {
		filterQuery += " AND status = 1 "
		queryArgs = append(queryArgs, req.Status)
	}

	if req.StartTime != nil {
		filterQuery += " AND start_date >= ? "
	}
	if req.EndTime != nil {
		filterQuery += " AND end_date <= ? "
	}

	if pagination.IsDESC {
		filterQuery += " ORDER BY id DESC"
	} else {
		filterQuery += " ORDER BY id ASC"
	}

	sqlStatement += filterQuery
	qry, err := mdb.conn.Query(sqlStatement, queryArgs...)
	if err != nil {

		return data, err
	}
	defer qry.Close()

	for qry.Next() {
		err = qry.Scan(
			&dataActivity.ID,
			&dataActivity.Title,
			&dataActivity.PosterURL,
			&dataActivity.Content,
			&dataActivity.Welfare,
			&dataActivity.Type,
			&dataActivity.StartTime,
			&dataActivity.EndTime,
			&dataActivity.Status,
			&dataActivity.CreatedAt,
			&dataActivity.UpdatedAt,
		)
		if err != nil {
			return data, err
		}

		data = append(data, dataActivity)
	}

	return data, nil
}

// CreateActivityRecord 创建参与记录
func (mdb *ActivityRepository) CreateActivityRecord(ctx context.Context, req *model.ActivityRecord) error {

	res, err := mdb.conn.ExecContext(
		ctx,
		createActivityRecordQuery,
		req.ActivityID,
		req.ParticipantID,
		req.Type,
		constant.Pending,
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

// UpdateActivityRecord 修改参与记录
func (mdb *ActivityRepository) UpdateActivityRecord(ctx context.Context, req *model.ActivityRecord) error {
	res, err := mdb.conn.ExecContext(
		ctx,
		updateActivityRecordQuery,
		req.Link,
		req.Note,
		req.Status,
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
