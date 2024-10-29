package ports

import (
	"github.com/gofiber/fiber/v2"
	"github.com/n0o01lh/ml-url-shortener/internals/core/domain"
)

type ResolverService interface {
	Get(id string) (*domain.ShortedUrl, error)
}

type ResolverRepository interface {
	Get(id string) (*domain.ShortedUrl, error)
}

type ResolverHandlers interface {
	Get(context *fiber.Ctx) error
}
