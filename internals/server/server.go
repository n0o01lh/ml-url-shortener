package server

import "github.com/gofiber/fiber/v2"

type Server struct {
}

func NewServer() *Server {
	return &Server{}
}

func (s *Server) Initialize() {
	app := fiber.New()

	app.Listen(":3000")
}
