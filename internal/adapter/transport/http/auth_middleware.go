package http

import (
	"errors"
	"gateway/internal/dto"
	"github.com/goccy/go-json"
	"net/http"
	"reflect"
	"strings"

	"github.com/gofiber/fiber/v3"
)

const (
	AuthHeader = "Authorization"
	AuthType   = "Bearer"
	AuthToken  = "AuthToken"
	AuthPublic = "AuthPublic"
	UserDetail = "UserDetail"
)

func (s *server) GetToken(c fiber.Ctx) error {
	token := c.Get(AuthHeader)
	if token == "" {
		return s.errorResponse(c, "authorization header is not provided", errors.New("authorization header is not provided"), nil, fiber.StatusUnauthorized)
	}

	fields := strings.Fields(token)
	if len(fields) != 2 {
		return s.errorResponse(c, "invalid authorization header format", errors.New("invalid authorization header format"), nil, fiber.StatusUnauthorized)
	}

	if fields[0] != AuthType {
		return s.errorResponse(c, "unsupported authorization type", errors.New("unsupported authorization type"), nil, fiber.StatusUnauthorized)
	}

	accessToken := fields[1]
	accessPublic := c.Get(AuthPublic)
	if accessPublic == "" {
		return s.errorResponse(c, "public key is not provided", errors.New("public key is not provided"), nil, fiber.StatusUnauthorized)
	}

	c.Locals(AuthToken, accessToken)
	c.Locals(AuthPublic, accessPublic)
	return c.Next()
}

func (s *server) GetUserDetail(c fiber.Ctx) error {
	accessToken, ok := c.Locals(AuthToken).(string)
	if !ok {
		return s.errorResponse(c, "access token not found in context", errors.New("access token not found"), nil, fiber.StatusUnauthorized)
	}

	accessPublic, ok := c.Locals(AuthPublic).(string)
	if !ok {
		return s.errorResponse(c, "public key not found in context", errors.New("public key not found"), nil, fiber.StatusUnauthorized)
	}

	req, err := http.NewRequest("POST", "http://localhost:8001/auth/get-me", nil)
	if err != nil {
		return s.errorResponse(c, "error creating new request", err, nil, fiber.StatusInternalServerError)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set(AuthHeader, accessToken)
	req.Header.Set(AuthPublic, accessPublic)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return s.errorResponse(c, "error sending request", err, nil, fiber.StatusInternalServerError)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return s.errorResponse(c, "received non-OK status code", errors.New("received non-OK status code"), nil, fiber.StatusUnauthorized)
	}

	var userDetail dto.GetUserResponse
	if err := json.NewDecoder(resp.Body).Decode(&userDetail); err != nil {
		return s.errorResponse(c, "error decoding response body", err, nil, fiber.StatusInternalServerError)
	}

	c.Locals(UserDetail, userDetail)
	return c.Next()
}

func (s *server) CheckPermission(requiredModule string, requiredAction string) fiber.Handler {
	return func(c fiber.Ctx) error {
		user, ok := c.Locals(UserDetail).(dto.GetUserResponse)
		if !ok {
			return s.errorResponse(c, "user detail not found in context", errors.New("user detail not found"), nil, fiber.StatusForbidden)
		}

		permissions, ok := user.Attributes[requiredModule]
		if !ok {
			return s.errorResponse(c, "module permission not found", errors.New("module permission not found"), nil, fiber.StatusForbidden)
		}

		// Use reflection to check the specific action permission
		val := reflect.ValueOf(permissions)
		field := val.FieldByName(requiredAction)
		if !field.IsValid() {
			return s.errorResponse(c, "invalid action", errors.New("invalid action"), nil, fiber.StatusBadRequest)
		}

		if !field.Bool() {
			return s.errorResponse(c, "action not allowed", errors.New("action not allowed"), nil, fiber.StatusForbidden)
		}

		return c.Next()
	}
}
