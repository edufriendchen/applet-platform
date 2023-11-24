package activity

import (
	"context"
	"fmt"
	"github.com/edufriendchen/applet-platform/constant"

	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/edufriendchen/applet-platform/model"
)

// GetActivityList 获取活动列表
func (s *Management) GetActivityList(ctx context.Context, req ListActivityRequest) ([]ListActivityResponse, error) {
	pagination := model.Pagination{
		PerPage: req.PerPage,
		Page:    req.Page,
	}

	data, err := s.activityRepository.GetActivityList(pagination, &model.Activity{
		ID:        req.ID,
		Type:      req.Type,
		StartTime: &req.StartTime,
		EndTime:   &req.EndTime,
	})
	if err != nil {
		hlog.CtxErrorf(ctx, "[GetActivityList] - GetActivityList err", err)

		return nil, err
	}

	var res []ListActivityResponse
	for _, item := range data {
		key := fmt.Sprintf(constant.ActivityVisitNumPrefix, item.ID)
		num, err := s.cache.GetInt(key)
		if err != nil {
			hlog.CtxErrorf(ctx, "[GetActivityList] - GetInt err", err)

			return nil, err
		}
		res = append(res, ListActivityResponse{
			VisitNum: int64(num),
		})
	}

	return res, nil
}

// GetActivityDetail 获取活动详情
func (s *Management) GetActivityDetail(ctx context.Context, id model.Activity) ([]ListActivityResponse, error) {
	return nil, nil
}

// ParticipateActivity 参与活动
func (s *Management) ParticipateActivity(ctx context.Context, req model.Activity) error {

	return nil
}

// AbandonActivity 放弃活动
func (s *Management) AbandonActivity(ctx context.Context, req model.Activity) error {

	return nil
}
