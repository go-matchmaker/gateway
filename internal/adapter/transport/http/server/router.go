package server

import "fmt"

func (s *server) SetupRouter() {
	s.authRouter()
}

func (s *server) authRouter() {
	route := s.app.Group("/auth")
	fmt.Println(route)
}
