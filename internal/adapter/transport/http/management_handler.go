package http

import (
	"bytes"
	"errors"
	"github.com/gofiber/fiber/v3"
	"net/http"
)

func (s *server) CreateUser(c fiber.Ctx) error {
	body := c.Body()
	if body == nil {
		return s.errorResponse(c, "Invalid request", nil, nil, 400)
	}

	req, err := http.NewRequest("POST", "http://localhost:8001/management/create", bytes.NewBuffer(body))
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

	return s.successResponse(c, nil, "User created successfully", fiber.StatusOK)
}
