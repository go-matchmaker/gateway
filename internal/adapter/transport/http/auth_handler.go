package http

import (
	"bytes"
	"errors"
	"fmt"
	"gateway/internal/dto"
	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v3"
	"net/http"
)

func (s *server) Login(c fiber.Ctx) error {
	loginBody := c.Body()
	if loginBody == nil {
		return s.errorResponse(c, "login body is empty", errors.New("login body is empty"), nil, fiber.StatusBadRequest)
	}

	req, err := http.NewRequest("POST", "http://localhost:8001/auth/login", bytes.NewBuffer(loginBody))
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
	fmt.Println("login success", resp.Body)
	var loginResponse dto.UserLoginResponse
	err = json.NewDecoder(resp.Body).Decode(&loginResponse)
	if err != nil {
		return s.errorResponse(c, "error decoding response", err, nil, fiber.StatusInternalServerError)
	}

	return s.successResponse(c, loginResponse, "login success", fiber.StatusOK)
}
