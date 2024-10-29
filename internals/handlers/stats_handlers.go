package handlers

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/n0o01lh/ml-url-shortener/internals/core/ports"
)

type StatsHandlers struct {
	service ports.ServiceOrchestrator
}

func NewStatsHandlers(service ports.ServiceOrchestrator) *StatsHandlers {
	return &StatsHandlers{service: service}
}

var _ ports.StatsHandlers = (*StatsHandlers)(nil)

func (h *StatsHandlers) Get(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	stats, err := h.service.GetStats(id)

	if err != nil {
		log.Error(err)
		ctx.SendStatus(http.StatusInternalServerError)
		return err
	}

	ctx.JSON(stats)
	ctx.SendStatus(http.StatusOK)

	return nil
}
