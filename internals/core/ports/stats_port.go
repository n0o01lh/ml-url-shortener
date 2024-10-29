package ports

import (
	"github.com/gofiber/fiber/v2"
	"github.com/n0o01lh/ml-url-shortener/internals/core/domain"
)

type StatsService interface {
	Create(id string) error
	Update(id string) error
	Get(id string) (*domain.Stats, error)
}

type StatsRepository interface {
	Create(stats *domain.Stats) error
	Update(id string) error
	Get(id string) (*domain.Stats, error)
}

type StatsHandlers interface {
	Get(context *fiber.Ctx) error
}
