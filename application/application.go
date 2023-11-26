package application

import (
	"github.com/edufriendchen/applet-platform/application/activity"
	"github.com/edufriendchen/applet-platform/application/auth"
	"github.com/edufriendchen/applet-platform/application/file"
	"github.com/edufriendchen/applet-platform/application/society"
	"github.com/edufriendchen/applet-platform/common/jwt"
	"github.com/edufriendchen/applet-platform/infrastructure/cache"
	"github.com/edufriendchen/applet-platform/infrastructure/repository"
	"github.com/edufriendchen/applet-platform/infrastructure/storage"
	"github.com/spf13/viper"
)

type Application struct {
	ActivityService      activity.IActivityService
	SocietyManagement    society.ISocietyManagement
	FileManagement       file.IFileManagement
	AuthManagement       auth.IAuthService
	WechatAuthManagement auth.ThirdPartyAuthService
	TiktokAuthManagement auth.ThirdPartyAuthService
}

func NewApplication(
	cache cache.CacheStore,
	repository repository.Repository,
	storage *storage.QiNiuStorage,
	tokenManager jwt.TokenManagerService,
) Application {
	return Application{
		ActivityService: activity.NewActivityService(cache, repository.ActivityRepository),
		FileManagement:  file.NewFileManagement(storage, repository.FileRepository),
		AuthManagement:  auth.NewAuthService(cache, repository.MemberRepository),
		WechatAuthManagement: auth.NewWechatAuthService(
			viper.GetString("third.party.auth.wechat.app-id"),
			viper.GetString("third.party.auth.wechat.app-secret"),
			viper.GetString("third.party.auth.wechat.app-secret"),
			repository.MemberRepository,
			tokenManager),
		TiktokAuthManagement: auth.NewTiktokAuthService(
			viper.GetString("third.party.auth.tiktok.app-id"),
			viper.GetString("third.party.auth.tiktok.app-secret"),
			repository.MemberRepository,
			tokenManager),
	}
}
