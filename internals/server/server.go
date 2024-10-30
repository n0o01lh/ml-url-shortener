package server

import (
	"github.com/gofiber/fiber/v2"
	"github.com/n0o01lh/ml-url-shortener/internals/core/ports"
)

type Server struct {
	shortHandlers   ports.ShortHandlers
	resolveHandlers ports.ResolverHandlers
	statsHandlers   ports.StatsHandlers
}

func NewServer(shortHandlers ports.ShortHandlers, resolveHandlers ports.ResolverHandlers, statsHandlers ports.StatsHandlers) *Server {
	return &Server{
		shortHandlers:   shortHandlers,
		resolveHandlers: resolveHandlers,
		statsHandlers:   statsHandlers,
	}
}

func (s *Server) Initialize() {
	app := fiber.New()

	shortRoutes := app.Group("/short")

	shortRoutes.Post("/create", s.shortHandlers.Create)
	shortRoutes.Patch("/update/:id", s.shortHandlers.Update)

	statsRoutes := app.Group("/stats")

	statsRoutes.Get("/:id", s.statsHandlers.Get)

	resolveRoutes := app.Group("/")

	resolveRoutes.Get("/:id", s.resolveHandlers.Get)

	app.Listen(":3000")
}
