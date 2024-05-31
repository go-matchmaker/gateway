package http

func (s *server) SetupRouter() {
	s.authRouter()
}

func (s *server) authRouter() {
	route := s.app.Group("/auth")
	route.Post("/login", s.Login)
}
