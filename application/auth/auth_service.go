package auth

import (
	"context"
	"fmt"
	"strings"

	"github.com/edufriendchen/applet-platform/common"
	"github.com/edufriendchen/applet-platform/constant"
	"github.com/edufriendchen/applet-platform/infrastructure/cache"
	"github.com/edufriendchen/applet-platform/infrastructure/repository"
	"github.com/edufriendchen/applet-platform/model"
)

const (
	RedisBlockedUserTknKey = "block_user_token:%d"
)

func NewAuthService(
	cache cache.CacheStore,
	userRepository repository.UserRepository,
) IAuthService {
	return &AuthService{
		cache:          cache,
		userRepository: userRepository,
	}
}

func (c *AuthService) ExternalAuthorize(ctx context.Context, authorization string) (context.Context, error) {

	if authorization == "" {
		return nil, common.ErrUnauthorized
	}

	fields := strings.Fields(authorization)
	if len(fields) < BearerValueTotalCount {
		return nil, common.ErrUnauthorized
	}

	if fields[0] != "Bearer" {
		return nil, common.ErrUnauthorized
	}

	accessToken := fields[1]

	// verify authorization token
	claims, err := c.tokenManager.VerifyAccessToken(accessToken)
	if err != nil {
		return nil, common.ErrForbidden
	}

	// check blocked user token
	key := fmt.Sprintf(RedisBlockedUserTknKey, claims.UserID)
	token, _ := c.cache.GetBytes(key)

	if token != nil && string(token) == accessToken {
		return nil, common.ErrUnauthorized
	}

	//get status user
	externalMembers, err := c.userRepository.GetUserList(ctx, model.User{
		ID: claims.UserID,
	})
	if err != nil {
		return nil, common.ErrInternal
	}

	if len(externalMembers) == 0 {
		return nil, common.ErrUserNotFound
	}

	switch externalMembers[0].Status {
	case constant.ExternalMemberDeactivate:
		return nil, common.ErrUserDeactivate
	case constant.ExternalMemberBlocked:
		return nil, common.ErrUserBlocked
	case constant.ExternalMemberInvited:
		return nil, common.ErrUserInvited
	}

	// set verified custom claims to context if needed here
	ctx = context.WithValue(ctx, constant.ExternalMemberID, claims.UserID)
	ctx = context.WithValue(ctx, constant.OpenID, claims.OpenID)

	return ctx, nil
}

func (c *AuthService) InternalAuthorize(ctx context.Context, authorization string) (context.Context, error) {

	if authorization == "" {
		return nil, common.ErrUnauthorized
	}

	fields := strings.Fields(authorization)
	if len(fields) < BearerValueTotalCount {
		return nil, common.ErrUnauthorized
	}

	if fields[0] != "Bearer" {
		return nil, common.ErrUnauthorized
	}

	accessToken := fields[1]

	// verify authorization token
	claims, err := c.tokenManager.VerifyAccessToken(accessToken)
	if err != nil {
		return nil, common.ErrForbidden
	}

	// check blocked user token
	key := fmt.Sprintf(RedisBlockedUserTknKey, claims.UserID)
	token, _ := c.cache.GetBytes(key)

	if token != nil && string(token) == accessToken {
		return nil, common.ErrUnauthorized
	}

	//get status user
	externalMembers, err := c.userRepository.GetUserList(ctx, model.User{
		ID: claims.UserID,
	})
	if err != nil {
		return nil, common.ErrInternal
	}

	if len(externalMembers) == 0 {
		return nil, common.ErrUserNotFound
	}

	switch externalMembers[0].Status {
	case constant.ExternalMemberDeactivate:
		return nil, common.ErrUserDeactivate
	case constant.ExternalMemberBlocked:
		return nil, common.ErrUserBlocked
	case constant.ExternalMemberInvited:
		return nil, common.ErrUserInvited
	}

	// set verified custom claims to context if needed here
	ctx = context.WithValue(ctx, constant.ExternalMemberID, claims.UserID)
	ctx = context.WithValue(ctx, constant.OpenID, claims.OpenID)

	return ctx, nil
}

func (c *AuthService) Login(ctx context.Context, authorization string) (context.Context, error) {

	if authorization == "" {
		return nil, common.ErrUnauthorized
	}

	fields := strings.Fields(authorization)
	if len(fields) < BearerValueTotalCount {
		return nil, common.ErrUnauthorized
	}

	if fields[0] != "Bearer" {
		return nil, common.ErrUnauthorized
	}

	accessToken := fields[1]

	// verify authorization token
	claims, err := c.tokenManager.VerifyAccessToken(accessToken)
	if err != nil {
		return nil, common.ErrForbidden
	}

	// check blocked user token
	key := fmt.Sprintf(RedisBlockedUserTknKey, claims.UserID)
	token, _ := c.cache.GetBytes(key)

	if token != nil && string(token) == accessToken {
		return nil, common.ErrUnauthorized
	}

	//get status user
	externalMembers, err := c.userRepository.GetUserList(ctx, model.User{
		ID: claims.UserID,
	})
	if err != nil {
		return nil, common.ErrInternal
	}

	if len(externalMembers) == 0 {
		return nil, common.ErrUserNotFound
	}

	switch externalMembers[0].Status {
	case constant.ExternalMemberDeactivate:
		return nil, common.ErrUserDeactivate
	case constant.ExternalMemberBlocked:
		return nil, common.ErrUserBlocked
	case constant.ExternalMemberInvited:
		return nil, common.ErrUserInvited
	}

	// set verified custom claims to context if needed here
	ctx = context.WithValue(ctx, constant.ExternalMemberID, claims.UserID)
	ctx = context.WithValue(ctx, constant.OpenID, claims.OpenID)

	return ctx, nil
}
