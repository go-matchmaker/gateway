package http

import (
	"time"
)

func (s *server) SetupRouter() {
	s.authRouter()
}

func (s *server) authRouter() {
	route := s.app.Group("/auth")
	route.Post("/login", s.Login, s.RateLimiter(5, time.Minute), s.LoginValidation)
}

func (s *server) hrManagementRouter() {
	route := s.app.Group("/hr-management", s.GetUserDetail)
	route.Post("/create-user", s.CreateUser, s.RateLimiter(5, time.Minute), s.HRAddPermission)
}
