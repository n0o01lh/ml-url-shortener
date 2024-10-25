package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/n0o01lh/ml-url-shortener/internals/core/domain"
	"github.com/n0o01lh/ml-url-shortener/internals/core/ports"
)

type ShortHandlers struct {
	shortService ports.ShortService
}

func NewShortHandlers(shortService ports.ShortService) *ShortHandlers {
	return &ShortHandlers{shortService: shortService}
}

var _ ports.ShortHandlers = (*ShortHandlers)(nil)

func (h *ShortHandlers) Create(ctx *fiber.Ctx) error {
	var requestBody map[string]any
	err := json.Unmarshal(ctx.Body(), &requestBody)

	if err != nil {
		ctx.SendStatus(http.StatusBadRequest)
		return err
	}

	url := requestBody["url"].(string)
	available := true //default value

	if requestBody["available"] != nil {
		available = requestBody["available"].(bool)
	}

	shortRequest := domain.NewShortRequest(url, available)
	shortedUrl, error := h.shortService.Create(shortRequest)

	if error != nil {
		log.Error(error)
		ctx.SendStatus(http.StatusInternalServerError)
		return error
	}

	ctx.Status(http.StatusCreated)
	ctx.JSON(shortedUrl)

	return nil
}

func (h *ShortHandlers) Update(ctx *fiber.Ctx) error {
	return nil
}
