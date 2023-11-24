package society

import (
	"context"
	"net/http"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/cloudwego/hertz/pkg/route"
	"github.com/edufriendchen/applet-platform/application/society"
	"github.com/edufriendchen/applet-platform/model"
)

// RestHandler handler
type RestHandler struct {
	societyAuth society.ISocietyManagement
}

// NewRestHandler create new rest handler
func NewRestHandler(
	societyAuth society.ISocietyManagement,
) *RestHandler {
	return &RestHandler{
		societyAuth: societyAuth,
	}
}

func (rh *RestHandler) Mount(router *route.RouterGroup) {
	{
		router.Handle(http.MethodGet, "/society", rh.GetSocietyList)
		router.Handle(http.MethodPost, "/society", rh.ApplySociety)
	}
}

func (rh *RestHandler) PrivateMount(router *route.RouterGroup) {
	{
		router.Handle(http.MethodGet, "/society", rh.GetSocietyList)
		router.Handle(http.MethodPost, "/society", rh.ApplySociety)
	}
}

func (rh *RestHandler) GetSocietyList(ctx context.Context, c *app.RequestContext) {
	var req model.Society
	err := c.BindAndValidate(&req)

	res, err := rh.societyAuth.GetSocieties(ctx, req)
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

func (rh *RestHandler) ApplySociety(ctx context.Context, c *app.RequestContext) {
	var req society.ApplySocietyRequest
	err := c.BindAndValidate(&req)

	err = rh.societyAuth.ApplySociety(ctx, req)
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.H{
			"message": "internal err",
		})

		return
	}

	c.JSON(http.StatusOK, utils.H{
		"message": "success",
		"data":    nil,
	})
}
