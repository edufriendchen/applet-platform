package auth

import (
	"context"
	"crypto/sha1"
	"fmt"
	"net/http"
	"sort"
	"strings"

	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/hertz/pkg/common/json"
	"github.com/edufriendchen/applet-platform/common/jwt"
	"github.com/edufriendchen/applet-platform/constant"
	"github.com/edufriendchen/applet-platform/infrastructure/cache"
	"github.com/edufriendchen/applet-platform/infrastructure/repository"
	"github.com/edufriendchen/applet-platform/model"
)

type WechatAuthProvider struct {
	appid          string
	appSecret      string
	token          string
	cache          cache.CacheStore
	userRepository repository.UserRepository
	tokenManager   jwt.TokenManagerService
}

func NewWechatAuthService(
	appid string,
	appSecret string,
	token string,
	userRepository repository.UserRepository,
	tokenManager jwt.TokenManagerService,
) ThirdPartyAuthService {
	return &WechatAuthProvider{
		appid:          appid,
		appSecret:      appSecret,
		token:          token,
		userRepository: userRepository,
		tokenManager:   tokenManager,
	}
}

func (s *WechatAuthProvider) Validate(ctx context.Context, req ValidateRequest) string {
	arr := []string{req.nonce, req.signature, s.token}
	sort.Strings(arr)

	str := strings.Join(arr, "")
	h := sha1.New()
	h.Write([]byte(str))
	sha1Sum := fmt.Sprintf("%x", h.Sum(nil))

	if sha1Sum != req.signature {

		return ""
	}

	return req.echoStr
}

func (s *WechatAuthProvider) LoginCredentialsVerification(ctx context.Context, code string) (*LoginResponse, error) {
	openidURL := fmt.Sprintf(WxAuthUrl, s.appid, s.appSecret, code)
	resp, err := http.Get(openidURL)
	if err != nil {
		hlog.Error("[LoginCredentialsVerification] Error http get request:", err)

		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		hlog.Error("[LoginCredentialsVerification] http status != 200:", err)

		return nil, err
	}

	var result map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {

		return nil, err
	}

	if _, exist := result["openid"]; !exist {
		return nil, ErrAuthorization
	}

	openID := result["openid"].(string)

	members, err := s.userRepository.GetUserList(ctx, model.User{
		WXOpenID: openID,
	})
	if err != nil {
		hlog.Error("[LoginCredentialsVerification] memberRepository.Get:", err)

		return nil, err
	}

	var membersID int64 = 0
	if len(members) == 0 {
		newMember := model.User{
			ID:       1,
			WXOpenID: openID,
			Status:   constant.ExternalMemberActive,
		}
		err := s.userRepository.CreateUser(ctx, &newMember)
		if err != nil {
			hlog.Error("[LoginCredentialsVerification] memberRepository.Create", err)

			return nil, err
		}
		membersID = int64(newMember.ID)
	} else {
		membersID = int64(members[0].ID)
	}

	accessToken, err := s.tokenManager.GenerateAccessToken(&jwt.UserForToken{
		UserID: uint64(membersID),
		OpenID: openID,
	})
	if err != nil {
		hlog.Error("[LoginCredentialsVerification] tokenManager.GenerateAccessToken err", err)

		return nil, err
	}

	refreshToken, err := s.tokenManager.GenerateRefreshToken(&jwt.UserForToken{
		UserID: 0,
		OpenID: openID,
	})
	if err != nil {
		hlog.Error("[LoginCredentialsVerification] tokenManager. GenerateRefreshToken err", err)

		return nil, err
	}

	return &LoginResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}

func (s *WechatAuthProvider) Binding(ctx context.Context, code string) (*LoginResponse, error) {

	openidURL := fmt.Sprintf(WxAuthUrl, s.appid, s.appSecret, code)
	resp, err := http.Get(openidURL)
	if err != nil {
		hlog.Error("[LoginCredentialsVerification] Error http get request:", err)

		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		hlog.Error("[LoginCredentialsVerification] http status != 200:", err)

		return nil, err
	}

	var result map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {

		return nil, err
	}

	if _, exist := result["openid"]; !exist {
		return nil, ErrAuthorization
	}

	openID := result["openid"].(string)

	members, err := s.userRepository.GetUserList(ctx, model.User{
		WXOpenID: openID,
	})
	if err != nil {
		hlog.Error("[LoginCredentialsVerification] memberRepository.Get:", err)

		return nil, err
	}

	var membersID int64 = 0
	if len(members) == 0 {
		newMember := model.User{
			ID:       1,
			WXOpenID: openID,
			Status:   constant.ExternalMemberActive,
		}
		err := s.userRepository.CreateUser(ctx, &newMember)
		if err != nil {
			hlog.Error("[LoginCredentialsVerification] memberRepository.Create", err)

			return nil, err
		}
		membersID = int64(newMember.ID)
	} else {
		membersID = int64(members[0].ID)
	}

	accessToken, err := s.tokenManager.GenerateAccessToken(&jwt.UserForToken{
		UserID: uint64(membersID),
		OpenID: openID,
	})
	if err != nil {
		hlog.Error("[LoginCredentialsVerification] tokenManager.GenerateAccessToken err", err)

		return nil, err
	}

	refreshToken, err := s.tokenManager.GenerateRefreshToken(&jwt.UserForToken{
		UserID: 0,
		OpenID: openID,
	})
	if err != nil {
		hlog.Error("[LoginCredentialsVerification] tokenManager. GenerateRefreshToken err", err)

		return nil, err
	}

	return &LoginResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}
