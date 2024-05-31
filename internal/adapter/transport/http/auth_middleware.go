package http

import (
	"bytes"
	"errors"
	"fmt"
	"gateway/internal/dto"
	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v3"
	"net/http"
	"strings"
)

const (
	AuthHeader = "Authorization"
	AuthType   = "Bearer"
	AuthPublic = "AuthPublic"
	UserDetail = "UserDetail"
)

func (s *server) GetUserDetail(c fiber.Ctx) error {
	token := c.Get(AuthHeader)
	if token == "" {
		return s.errorResponse(c, "authorization header is not provided", errors.New("authorization header is not provided"), nil, fiber.StatusUnauthorized)
	}

	fields := strings.Fields(token)
	if len(fields) != 2 {
		return s.errorResponse(c, "invalid authorization header format", errors.New("invalid authorization header format"), nil, fiber.StatusUnauthorized)
	}

	authorizationType := strings.ToLower(fields[0])
	if authorizationType != AuthType {
		return s.errorResponse(c, fmt.Sprintf("unsupported authorization type %s", authorizationType), errors.New(fmt.Sprintf("unsupported authorization type %s", authorizationType)), nil, fiber.StatusUnauthorized)
	}

	accessToken := fields[1]
	accessPublic := c.Get(AuthPublic)
	if accessPublic == "" {
		return s.errorResponse(c, "public key is not provided", errors.New("public key is not provided"), nil, fiber.StatusUnauthorized)
	}

	loginBody := dto.UserAuthRequest{
		Token:     accessToken,
		PublicKey: accessPublic,
	}

	jsonData, err := json.Marshal(loginBody)
	if err != nil {
		return s.errorResponse(c, "error marshalling loginBody", err, nil, fiber.StatusInternalServerError)
	}

	req, err := http.NewRequest("POST", "http://localhost:8001/auth/get-detail", bytes.NewBuffer(jsonData))
	if err != nil {
		return s.errorResponse(c, "error creating new request", err, nil, fiber.StatusInternalServerError)
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return s.errorResponse(c, "error sending request", err, nil, fiber.StatusInternalServerError)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return s.errorResponse(c, "received non-OK status code", errors.New("received non-OK status code"), nil, fiber.StatusUnauthorized)
	}

	var authResponse dto.AuthMiddlewareResponse
	err = json.NewDecoder(resp.Body).Decode(&authResponse)
	if err != nil {
		return s.errorResponse(c, "error decoding response", err, nil, fiber.StatusInternalServerError)
	}

	c.Locals(UserDetail, authResponse)

	return c.Next()
}

func (s *server) HRPermission(c fiber.Ctx) error {
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
