package handlers

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/n0o01lh/ml-url-shortener/internals/core/ports"
)

type ResolverHandlers struct {
	service ports.ServiceOrchestrator
}

func NewResolverHandlers(service ports.ServiceOrchestrator) *ResolverHandlers {
	return &ResolverHandlers{service: service}
}

var _ ports.ResolverHandlers = (*ResolverHandlers)(nil)

func (h *ResolverHandlers) Get(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	shortedUrlUpdated, err := h.service.GetShortUrl(id)

	if err != nil {
		log.Error(err)
		ctx.SendStatus(http.StatusInternalServerError)
		return err
	}

	return ctx.Redirect(shortedUrlUpdated.OriginalUrl, fiber.StatusFound)
}
