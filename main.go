// Code generated by hertz generator.

package main

import (
	"github.com/edufriendchen/applet-platform/common"
	"github.com/edufriendchen/applet-platform/constant"
	"log"
	"time"

	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/edufriendchen/applet-platform/application"
	"github.com/edufriendchen/applet-platform/common/jwt"
	"github.com/edufriendchen/applet-platform/infrastructure/cache"
	"github.com/edufriendchen/applet-platform/infrastructure/repository"
	"github.com/edufriendchen/applet-platform/infrastructure/storage"
	"github.com/edufriendchen/applet-platform/transport"
	hertzlogrus "github.com/hertz-contrib/obs-opentelemetry/logging/logrus"
	"github.com/spf13/viper"
)

func main() {
	// init config
	common.InitConfig()

	// init log components
	hlog.SetLogger(hertzlogrus.NewLogger())
	hlog.SetLevel(hlog.LevelDebug)

	// init hertz server
	hertz := server.Default(
		server.WithStreamBody(true),
		server.WithHostPorts(viper.GetString(constant.ServerAddress)),
	)

	// init mysql drive
	db, err := common.GetDatabase("database.master")
	if err != nil {
		log.Println("error getting database master")
		panic(err)
	}

	// init redis
	cacheImpl, err := cache.NewCacheStore3(viper.GetString("redis.master.host"), viper.GetString("redis.master.port"))
	if err != nil {
		hlog.Fatal("[Main] redis conn err:", err)
	}

	// init storage
	storageImpl, err := storage.NewStorage(viper.GetString("storage.qi-niu.access-key"), viper.GetString("storage.qi-niu.secret-key"))
	if err != nil {
		hlog.Fatal("[Main] storage  conn err:", err)
	}

	// new repository
	repositoryImpl := repository.NewRepository(db)

	// new jwtManage
	accessTokenDuration, err := time.ParseDuration(viper.GetString("server.auth.jwt.access.duration"))
	if err != nil {
		hlog.Fatal("[Main] ParseDuration accessTokenDuration err:", err)
	}
	refreshTokenDuration, err := time.ParseDuration(viper.GetString("server.auth.jwt.refresh.duration"))
	if err != nil {
		hlog.Fatal("[Main] ParseDuration refreshTokenDuration err:", err)
	}
	tokenManager := jwt.NewJWTManager(
		viper.GetString("server.auth.jwt.access.secretkey"),
		viper.GetString("server.auth.jwt.refresh.secretkey"),
		accessTokenDuration,
		refreshTokenDuration,
	)

	// new application
	restServer := application.NewApplication(cacheImpl, repositoryImpl, storageImpl, tokenManager)
	externalGroup := hertz.Group("/external")

	// new transport
	transport.NewHttpServer(restServer, externalGroup)

	hertz.Spin()
}