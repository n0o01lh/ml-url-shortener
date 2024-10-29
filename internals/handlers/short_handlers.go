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
	service ports.ServiceOrchestrator
}

func NewShortHandlers(service ports.ServiceOrchestrator) *ShortHandlers {
	return &ShortHandlers{service: service}
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

	shortRequest := domain.NewShortRequest(url, &available)
	shortedUrl, error := h.service.CreateShortUrl(shortRequest)

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
	id := ctx.Params("id")
	shortedRequest := new(domain.ShortRequest)

	if err := ctx.BodyParser(&shortedRequest); err != nil {
		log.Error(err)
		ctx.SendStatus(http.StatusBadRequest)
		return err
	}

	shortedUrlUpdated, err := h.service.UpdateShortUrl(id, shortedRequest)

	if err != nil {
		log.Error(err)
		ctx.SendStatus(http.StatusInternalServerError)
		return err
	}

	ctx.JSON(shortedUrlUpdated)

	return nil
}
