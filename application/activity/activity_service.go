package activity

import (
	"context"
	"fmt"
	"github.com/edufriendchen/applet-platform/common"
	"github.com/edufriendchen/applet-platform/constant"

	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/edufriendchen/applet-platform/model"
)

// GetActivityList 获取活动列表
func (s *Service) GetActivityList(ctx context.Context, req Request) ([]Response, error) {
	pagination := model.Pagination{
		PerPage: req.PerPage,
		Page:    req.Page,
	}

	data, err := s.activityRepository.GetActivityList(pagination, &model.Activity{
		ID:        req.ID,
		Type:      req.Type,
		StartTime: req.StartTime,
		EndTime:   req.EndTime,
	})
	if err != nil {
		hlog.CtxErrorf(ctx, "[GetActivityList] - GetActivityList err", err)

		return nil, err
	}

	var res []Response
	for _, item := range data {
		key := fmt.Sprintf(constant.ActivityVisitNumPrefix, item.ID)
		num, err := s.cache.GetInt(key)
		if err != nil {
			hlog.CtxErrorf(ctx, "[GetActivityList] - GetInt err", err)

			return nil, err
		}
		res = append(res, Response{
			ID:        item.ID,
			PosterUrl: item.PosterURL,
			Title:     item.Title,
			Type:      item.Type,
			StartTime: item.StartTime,
			EndTime:   item.EndTime,
			Status:    item.Status,
			VisitNum:  int64(num),
		})
	}

	return res, nil
}

// GetActivityDetail 获取活动详情
func (s *Service) GetActivityDetail(ctx context.Context, id uint64) (DetailResponse, error) {
	return DetailResponse{}, nil
}

// ParticipateActivity 参与活动
func (s *Service) ParticipateActivity(ctx context.Context, id uint64) error {

	userID := ctx.Value(constant.CtxUserIDInfo).(uint64)
	if userID == 0 {
		hlog.CtxErrorf(ctx, "[AbandonActivity] - Get Context with user id is 0")

		return common.ErrUserNotFound
	}

	err := s.activityRepository.CreateActivityRecord(ctx, &model.ActivityRecord{
		ActivityID:    id,
		ParticipantID: userID,
		Status:        constant.Pending,
	})
	if err != nil {
		hlog.CtxErrorf(ctx, "[AbandonActivity] - UpdateActivityRecord", err)

		return err
	}

	return nil
}

// AbandonActivity 放弃活动
func (s *Service) AbandonActivity(ctx context.Context, req AbandonRequest) error {

	err := s.activityRepository.UpdateActivityRecord(ctx, &model.ActivityRecord{
		Note:   req.Reason,
		Status: constant.Deactivate,
	})
	if err != nil {
		hlog.CtxErrorf(ctx, "[AbandonActivity] - UpdateActivityRecord", err)

		return err
	}

	return nil
}

// SubmitActivity 提交活动
func (s *Service) SubmitActivity(ctx context.Context, req model.Activity) error {

	err := s.activityRepository.UpdateActivityRecord(ctx, &model.ActivityRecord{
		Link:   "",
		Note:   "",
		Status: constant.Active,
	})
	if err != nil {
		hlog.CtxErrorf(ctx, "[SubmitActivity] - UpdateActivityRecord", err)

		return err
	}

	return nil
}
