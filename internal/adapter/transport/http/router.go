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
	route.Post("/create-user", s.CreateUser, s.RateLimiter(5, time.Minute), s.CheckPermission("HR", "add"))
	route.Post("/create-department", s.CreateDepartment, s.RateLimiter(5, time.Minute), s.CheckPermission("HR", "add"))
}

func (s *server) carManagementRouter() {
	route := s.app.Group("/car-management", s.GetUserDetail)
	route.Post("/create-car", s.CreateCar, s.RateLimiter(5, time.Minute), s.CheckPermission("Car", "add"))
	route.Post("/create-brand", s.CreateBrand, s.RateLimiter(5, time.Minute), s.CheckPermission("Car", "add"))
	route.Post("/create-model", s.CreateModel, s.RateLimiter(5, time.Minute), s.CheckPermission("Car", "add"))
	route.Post("/create-color", s.CreateColor, s.RateLimiter(5, time.Minute), s.CheckPermission("Car", "add"))
	route.Post("/create-fuel", s.CreateFuel, s.RateLimiter(5, time.Minute), s.CheckPermission("Car", "add"))
	route.Post("/create-transmission", s.CreateTransmission, s.RateLimiter(5, time.Minute), s.CheckPermission("Car", "add"))
}
