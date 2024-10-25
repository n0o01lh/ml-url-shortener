package server

import (
	"github.com/gofiber/fiber/v2"
	"github.com/n0o01lh/ml-url-shortener/internals/core/ports"
)

type Server struct {
	shortHandlers ports.ShortHandlers
}

func NewServer(shortHandlers ports.ShortHandlers) *Server {
	return &Server{shortHandlers: shortHandlers}
}

func (s *Server) Initialize() {
	app := fiber.New()

	shortRoutes := app.Group("/short")

	shortRoutes.Post("/create", s.shortHandlers.Create)
	shortRoutes.Patch("/update/:id", s.shortHandlers.Update)

	app.Listen(":3000")
}
