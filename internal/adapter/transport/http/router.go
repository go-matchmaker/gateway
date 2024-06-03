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
	route := s.app.Group("/hr-management", s.GetToken, s.GetUserDetail, s.RateLimiter(5, time.Minute))
	route.Post("/create-user", s.CreateUser, s.CheckPermission("HR", "add"))
	route.Post("/create-department", s.CreateDepartment, s.CheckPermission("HR", "add"))
}

func (s *server) carManagementRouter() {
	route := s.app.Group("/car-management", s.GetToken, s.GetUserDetail, s.RateLimiter(5, time.Minute))
	route.Post("/create-car", s.CreateCar, s.CheckPermission("Car", "add"))
	route.Post("/create-brand", s.CreateBrand, s.CheckPermission("Car", "add"))
	route.Post("/create-model", s.CreateModel, s.CheckPermission("Car", "add"))
	route.Post("/create-color", s.CreateColor, s.CheckPermission("Car", "add"))
	route.Post("/create-fuel", s.CreateFuel, s.CheckPermission("Car", "add"))
	route.Post("/create-transmission", s.CreateTransmission, s.CheckPermission("Car", "add"))
}
