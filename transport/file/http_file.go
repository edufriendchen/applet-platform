package file

import (
	"context"
	"net/http"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/cloudwego/hertz/pkg/route"
	"github.com/edufriendchen/applet-platform/application/file"
	"github.com/edufriendchen/applet-platform/model"
)

// RestHandler handler
type RestHandler struct {
	fileManage file.IFileManagement
}

// NewRestHandler create new rest handler
func NewRestHandler(
	fileManage file.IFileManagement,
) *RestHandler {
	return &RestHandler{
		fileManage: fileManage,
	}
}

func (rh *RestHandler) Mount(router *route.RouterGroup) {
	{
		router.Handle(http.MethodPost, "/file", rh.UploadFile)
	}
}

func (rh *RestHandler) UploadFile(ctx context.Context, c *app.RequestContext) {

	var req model.UploadFileRequest
	err := c.BindAndValidate(&req)

	res, err := rh.fileManage.UploadFile(ctx, req)
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
