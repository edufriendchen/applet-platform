package activity

import (
	"context"
	"net/http"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/cloudwego/hertz/pkg/route"
	"github.com/edufriendchen/applet-platform/application/activity"
)

type RestHandler struct {
	activity activity.IActivityService
}

func NewRestHandler(
	activity activity.IActivityService,
) *RestHandler {
	return &RestHandler{
		activity: activity,
	}
}

func (rh *RestHandler) Mount(router *route.RouterGroup) {
	{
		router.Handle(http.MethodGet, "/activity", rh.GetActivityList)
		router.Handle(http.MethodGet, "/activity/detail", rh.GetActivityList)
		router.Handle(http.MethodPost, "/activity", rh.GetActivityList)
		router.Handle(http.MethodDelete, "/activity", rh.GetActivityList)
	}
}

func (rh *RestHandler) GetActivityList(ctx context.Context, c *app.RequestContext) {
	var req activity.Request
	err := c.BindAndValidate(&req)

	res, err := rh.activity.GetActivityList(ctx, req)
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.H{
			"message": "internal err",
		})

		return
	}

	c.JSON(http.StatusOK, utils.H{
		"message": "success",
		"data":    res,
	})
}
