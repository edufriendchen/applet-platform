package transport

import (
	"net/http"

	"github.com/cloudwego/hertz/pkg/route"
	"github.com/edufriendchen/applet-platform/application"
	activityRh "github.com/edufriendchen/applet-platform/transport/activity"
	"github.com/edufriendchen/applet-platform/transport/auth"
	fileRh "github.com/edufriendchen/applet-platform/transport/file"
	societyRh "github.com/edufriendchen/applet-platform/transport/society"
)

func NewHttpServer(application application.Application, router *route.RouterGroup) {

	router.Use()
	societyRh.NewRestHandler(application.SocietyManagement).Mount(router)
	activityRh.NewRestHandler(application.ActivityManagement).Mount(router)
	fileRh.NewRestHandler(application.FileManagement).Mount(router)
	auth.NewRestHandler(application.AuthManagement, application.WechatAuthManagement, application.TiktokAuthManagement).Mount(router)
}

func httpResponseWrite(rw http.ResponseWriter, response interface{}, statusCode int) {
}
