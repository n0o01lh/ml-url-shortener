package ports

import (
	"github.com/gofiber/fiber/v2"
	domain "github.com/n0o01lh/ml-url-shortener/internals/core/domain"
)

type ShortService interface {
	Create(shortRequest *domain.ShortRequest) (*domain.ShortedUrl, error)
	Update(id string, shortRequest *domain.ShortRequest) (*domain.ShortedUrl, error)
}

type ShortRepository interface {
	Create(shortedUrl *domain.ShortedUrl) (*domain.ShortedUrl, error)
	Update(id string, shortedUrl *domain.ShortedUrl) (*domain.ShortedUrl, error)
}

type ShortHandlers interface {
	Create(context *fiber.Ctx) error
	Update(context *fiber.Ctx) error
}
