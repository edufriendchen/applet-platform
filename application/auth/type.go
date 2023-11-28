package auth

import (
	"context"
	"errors"
	"github.com/edufriendchen/applet-platform/infrastructure/repository"

	"github.com/edufriendchen/applet-platform/common/jwt"
	"github.com/edufriendchen/applet-platform/infrastructure/cache"
)

var (
	ErrAuthorization = errors.New("third party authorization error")
)

const (
	BearerValueTotalCount = 2
	WxAuthUrl             = "https://api.weixin.qq.com/sns/jscode2session?appid=%s&secret=%s&js_code=%s&grant_type=authorization_code"
)

type AuthService struct {
	cache          cache.CacheStore
	userRepository repository.UserRepository
	tokenManager   jwt.TokenManagerService
}

type IAuthService interface {
	ExternalAuthorize(ctx context.Context, authorization string) (context.Context, error)
	InternalAuthorize(ctx context.Context, authorization string) (context.Context, error)
}

type ThirdPartyAuthService interface {
	Validate(ctx context.Context, req ValidateRequest) string
	LoginCredentialsVerification(ctx context.Context, code string) (*LoginResponse, error)
	Binding(ctx context.Context, code string) (*LoginResponse, error)
}

type ValidateRequest struct {
	nonce     string
	timestamp string
	signature string
	echoStr   string
}

type LoginResponse struct {
	Email        string `json:"email"`
	FullName     string `json:"full_name"`
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}
