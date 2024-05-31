package http

import (
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/limiter"
	"time"
)

func (s *server) RateLimiter(max int, expiration time.Duration) func(c fiber.Ctx) error {
	return func(c fiber.Ctx) error {
		limiterMiddleware := limiter.New(limiter.Config{
			Max:          max,
			Expiration:   expiration,
			LimitReached: limitReachedFunc,
			KeyGenerator: func(c fiber.Ctx) string {
				remoteIp := c.IP()
				if c.Get("X-NginX-Proxy") == "true" {
					remoteIp = c.Get("X-Real-IP")
				}
				return remoteIp
			},
		})

		err := limiterMiddleware(c)
		if err != nil {
			return err
		}

		return c.Next()
	}
}
