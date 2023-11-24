package repository

import (
	"github.com/edufriendchen/applet-platform/model"
	"github.com/jmoiron/sqlx"
	"strconv"
)

const (
	getActivityListQuery = `
		SELECT id, title, poster_url, content, welfare, type, start_time, end_time, status 
		FROM activity 
		WHERE true`
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

	var sqlStatement = getActivityListQuery
	var filterQuery string
	queryArgs := make([]interface{}, 0)

	if req.ID > 0 {
		filterQuery += " and id = ? "
		queryArgs = append(queryArgs, strconv.FormatUint(req.ID, 10))
	}

	if req.Type > 0 {
		filterQuery += " and type = ? "
		queryArgs = append(queryArgs, req.Type)
	}

	if req.Status > 0 {
		filterQuery += " and enabled = 1 "
	}

	if req.StartTime != nil {
		filterQuery += " AND SUBSTRING(platform_flag_bit, 1, 1) = '1' "
	}
	if req.EndTime != nil {
		filterQuery += " AND SUBSTRING(platform_flag_bit, 2, 1) = '1' "
	}

	if pagination.IsDESC {
		filterQuery += " ORDER BY order_i DESC"
	} else {
		filterQuery += " ORDER BY order_i ASC"
	}

	sqlStatement += filterQuery
	qry, err := mdb.conn.Query(sqlStatement, queryArgs...)
	if err != nil {

		return data, err
	}
	defer qry.Close()

	for qry.Next() {
		dataActivity := model.Activity{}
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
