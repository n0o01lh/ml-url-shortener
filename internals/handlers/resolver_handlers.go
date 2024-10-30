package handlers

import (
	"context"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/n0o01lh/ml-url-shortener/internals/core/ports"
	"github.com/n0o01lh/ml-url-shortener/internals/core/services"
	"golang.org/x/time/rate"
)

type ResolverHandlers struct {
	service *services.ServiceOrchestrator
	limiter *rate.Limiter
}

func NewResolverHandlers(service *services.ServiceOrchestrator, limiter *rate.Limiter) *ResolverHandlers {
	return &ResolverHandlers{service: service, limiter: limiter}
}

var _ ports.ResolverHandlers = (*ResolverHandlers)(nil)

func (h *ResolverHandlers) Get(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	err := h.limiter.Wait(context.TODO())

	if err != nil {
		return err
	}

	h.service.GetShortUrl(id)
	result := <-h.service.GetJobQueue().ResultChannel

	if result.Error != nil {
		log.Error(result.Error)
		ctx.SendStatus(http.StatusInternalServerError)
		return result.Error
	}

	return ctx.Redirect(result.ShortedUrl.OriginalUrl, fiber.StatusFound)
}
