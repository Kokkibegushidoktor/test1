package handlers

import (
	"errors"
	"github.com/Kokkibegushidoktor/test1/internal/models"
	"github.com/Kokkibegushidoktor/test1/internal/service"
	"github.com/gofiber/fiber/v2"
)

type addTickerInput struct {
	Ticker string `json:"ticker" validate:"required"`
}

func (h *Handler) AddTicker(c *fiber.Ctx) error {
	var inp addTickerInput
	if err := c.BodyParser(&inp); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	if err := h.services.Tickers.Create(c.Context(), service.CreateTickerInput{Symbol: inp.Ticker}); err != nil {
		if errors.Is(err, models.ErrTickerAlreadyExists) {
			return c.Status(fiber.StatusBadRequest).SendString(err.Error())
		}
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	return c.Status(fiber.StatusOK).JSON(map[string]string{"ticker": inp.Ticker})
}
