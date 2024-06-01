package http

import (
	"errors"
	"gateway/internal/core/util"
	"gateway/internal/dto"
	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v3"
)

const (
	AuthHeader = "Authorization"
	AuthType   = "Bearer"
	AuthPublic = "AuthPublic"
	UserDetail = "UserDetail"
)

func (s *server) GetUserDetail(c fiber.Ctx) error {
	sess, err := s.session.Get(c)
	if err != nil {
		return s.errorResponse(c, "error getting session", err, nil, fiber.StatusInternalServerError)
	}

	userDetail := sess.Get("userDetail")
	if userDetail == nil {
		return s.errorResponse(c, "user not authenticated", errors.New("user not authenticated"), nil, fiber.StatusUnauthorized)
	}

	var user dto.UserLoginResponse
	err = json.Unmarshal([]byte(userDetail.(string)), &user)
	if err != nil {
		return s.errorResponse(c, "error unmarshalling user detail", err, nil, fiber.StatusInternalServerError)
	}

	cacheKey := util.GenerateCacheKey("permission", user.User.ID)
	cachedPermissions, err := s.cache.Get(c.Context(), cacheKey)
	if err != nil {
		return s.errorResponse(c, "error getting cache", err, nil, fiber.StatusInternalServerError)
	}

	var permissions map[string]dto.Permission
	err = json.Unmarshal(cachedPermissions, &permissions)
	if err != nil {
		return s.errorResponse(c, "error unmarshalling permissions", err, nil, fiber.StatusInternalServerError)
	}

	user.User.UserPermissions = permissions
	c.Locals(UserDetail, user)
	return c.Next()
}

func (s *server) HRAddPermission(c fiber.Ctx) error {
	userDetail, ok := c.Locals(UserDetail).(dto.AuthMiddlewareResponse)
	if !ok {
		return s.errorResponse(c, "user detail not found in context", errors.New("user detail not found in context"), nil, fiber.StatusUnauthorized)
	}

	hrAttributes, ok := userDetail.Attributes["HR"]
	if !ok {
		return s.errorResponse(c, "user does not have HR permission", errors.New("user does not have HR permission"), nil, fiber.StatusUnauthorized)
	}

	if !hrAttributes.Add {
		return s.errorResponse(c, "user does not have create permission ", errors.New("user does not have create permission "), nil, fiber.StatusUnauthorized)
	}

	return c.Next()
}
