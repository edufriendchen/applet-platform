package auth

import (
	"bytes"
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

type TiktokAuthProvider struct {
	clientKey        string
	clientSecret     string
	cache            cache.CacheStore
	memberRepository repository.UserRepository
	tokenManager     jwt.TokenManagerService
}

func NewTiktokAuthService(
	clientKey string,
	clientSecret string,
	memberRepository repository.UserRepository,
	tokenManager jwt.TokenManagerService,
) ThirdPartyAuthService {
	return &TiktokAuthProvider{
		clientKey:        clientKey,
		clientSecret:     clientSecret,
		memberRepository: memberRepository,
		tokenManager:     tokenManager,
	}
}

func (s *TiktokAuthProvider) Validate(ctx context.Context, req ValidateRequest) string {
	arr := []string{req.nonce, req.signature}
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

func (s *TiktokAuthProvider) LoginCredentialsVerification(ctx context.Context, code string) (*LoginResponse, error) {
	url := "https://open.douyin.com/oauth/access_token/"

	data := map[string]string{
		"grant_type":    "authorization_code",
		"client_key":    s.clientKey,
		"client_secret": s.clientSecret,
		"code":          code,
	}

	jsonData, _ := json.Marshal(data)

	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		hlog.Error("[LoginCredentialsVerification] http post err", err)

		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		hlog.Error("[LoginCredentialsVerification] http status err", err)

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

	members, err := s.memberRepository.GetUserList(ctx, model.User{
		WXOpenID: openID,
	})
	if err != nil {
		hlog.Error("[LoginCredentialsVerification] memberRepository.Get:", err)

		return nil, err
	}

	var membersID int64 = 0
	if len(members) == 0 {
		req := model.User{
			WXOpenID: openID,
			Status:   constant.ExternalMemberActive,
		}

		err := s.memberRepository.CreateUser(ctx, &req)
		if err != nil {
			hlog.Error("[LoginCredentialsVerification] memberRepository.Create", err)

			return nil, err
		}
		membersID = int64(req.ID)
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

func (s *TiktokAuthProvider) Binding(ctx context.Context, code string) (*LoginResponse, error) {
	return &LoginResponse{}, nil
}
