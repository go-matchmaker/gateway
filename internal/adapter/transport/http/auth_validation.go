package http

import (
	"errors"
	"gateway/internal/dto"
	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v3"
)

func (s *server) LoginValidation(c fiber.Ctx) error {
	data := new(dto.UserLoginRequest)
	body := c.Body()
	err := json.Unmarshal(body, &data)
	if err != nil {
		return s.errorResponse(c, "invalid request body", err, nil, fiber.StatusBadRequest)
	}

	validationErrors := ValidateRequestByStruct(body)
	if len(validationErrors) > 0 {
		return s.errorResponse(c, "validation failed", errors.New("validaiton error"), validationErrors, fiber.StatusUnprocessableEntity)
	}

	return c.Next()
}
