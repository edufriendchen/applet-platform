package auth

import (
	"context"
	"net/http"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/cloudwego/hertz/pkg/route"
	"github.com/edufriendchen/applet-platform/application/auth"
	"github.com/edufriendchen/applet-platform/model"
)

// RestHandler handler
type RestHandler struct {
	auth       auth.IAuthService
	wechatAuth auth.ThirdPartyAuthService
	tiktokAuth auth.ThirdPartyAuthService
}

// NewRestHandler create new rest handler
func NewRestHandler(
	auth auth.IAuthService,
	wechatAuth auth.ThirdPartyAuthService,
	tiktokAuth auth.ThirdPartyAuthService,
) *RestHandler {
	return &RestHandler{
		auth:       auth,
		wechatAuth: wechatAuth,
		tiktokAuth: tiktokAuth,
	}
}

func (rh *RestHandler) Mount(router *route.RouterGroup) {
	{
		router.Handle(http.MethodPost, "/login", rh.Login)
		router.Handle(http.MethodPost, "/refresh", rh.Refresh)
		router.Handle(http.MethodPost, "/wechat/auth/:code", rh.WechatAuthorize)
		router.Handle(http.MethodPost, "/tiktok/auth/:code", rh.TiktokAuthorize)
		router.Handle(http.MethodPost, "/wechat/bundle/:code", rh.WechatBundle)
		router.Handle(http.MethodPost, "/tiktok/Bundle/:code", rh.TiktokBundle)
		router.Handle(http.MethodPost, "/logout", rh.Logout)
	}
}

func (rh *RestHandler) Login(ctx context.Context, c *app.RequestContext) {
	var req model.Activity
	err := c.BindAndValidate(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.H{})

		return
	}

	code := c.Param("code")

	c.JSON(http.StatusOK, utils.H{
		"message": "success",
		"data":    code,
	})
}

func (rh *RestHandler) Refresh(ctx context.Context, c *app.RequestContext) {
	var req model.Activity
	err := c.BindAndValidate(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.H{
			"message": err.Error(),
		})

		return
	}

	code := c.Param("code")

	c.JSON(http.StatusOK, utils.H{
		"message": "success",
		"data":    code,
	})
}

func (rh *RestHandler) WechatAuthorize(ctx context.Context, c *app.RequestContext) {
	var req model.Activity
	err := c.BindAndValidate(&req)

	code := c.Param("code")

	res, err := rh.wechatAuth.LoginCredentialsVerification(ctx, code)
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.H{
			"message": err.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, utils.H{
		"message": "success",
		"data":    res,
	})
}

func (rh *RestHandler) TiktokAuthorize(ctx context.Context, c *app.RequestContext) {
	var req model.Activity
	err := c.BindAndValidate(&req)

	code := c.Param("code")

	res, err := rh.tiktokAuth.LoginCredentialsVerification(ctx, code)
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.H{
			"message": err.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, utils.H{
		"message": "success",
		"data":    res,
	})
}

func (rh *RestHandler) WechatBundle(ctx context.Context, c *app.RequestContext) {
	var req model.Activity
	err := c.BindAndValidate(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.H{
			"message": err.Error(),
		})

		return
	}
	code := c.Param("code")

	c.JSON(http.StatusOK, utils.H{
		"message": "success",
		"data":    code,
	})
}

func (rh *RestHandler) TiktokBundle(ctx context.Context, c *app.RequestContext) {
	var req model.Activity
	err := c.BindAndValidate(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.H{
			"message": err.Error(),
		})

		return
	}
	code := c.Param("code")

	c.JSON(http.StatusOK, utils.H{
		"message": "success",
		"data":    code,
	})
}

func (rh *RestHandler) Logout(ctx context.Context, c *app.RequestContext) {
	var req model.Activity
	err := c.BindAndValidate(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.H{
			"message": err.Error(),
		})

		return
	}

	code := c.Param("code")

	c.JSON(http.StatusOK, utils.H{
		"message": "success",
		"data":    code,
	})
}
