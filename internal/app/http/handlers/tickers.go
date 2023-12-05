package handlers

import (
	"errors"
	"github.com/Kokkibegushidoktor/test1/internal/models"
	"github.com/Kokkibegushidoktor/test1/internal/service"
	"github.com/gofiber/fiber/v2"
	"strings"
	"time"
)

type addTickerInput struct {
	Ticker string `json:"ticker"`
}

func (h *Handler) AddTicker(c *fiber.Ctx) error {
	var inp addTickerInput
	if err := c.BodyParser(&inp); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	inp.Ticker = strings.ToUpper(inp.Ticker)

	if err := h.services.Tickers.Create(c.Context(), service.CreateTickerInput{Symbol: inp.Ticker}); err != nil {
		if errors.Is(err, models.ErrTickerAlreadyExists) {
			return c.Status(fiber.StatusBadRequest).SendString(err.Error())
		}
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	return c.Status(fiber.StatusOK).JSON(map[string]string{"ticker": inp.Ticker})
}

type fetchTickerInput struct {
	Ticker string `json:"ticker"`
	From   string `json:"from"`
	To     string `json:"to"`
}

func (h *Handler) Fetch(c *fiber.Ctx) error {
	inp := fetchTickerInput{
		Ticker: c.Query("ticker"),
		From:   c.Query("from"),
		To:     c.Query("to"),
	}

	var err error
	var serviceInput service.FetchTickerInput

	serviceInput.DateFrom, err = time.Parse(time.RFC3339, inp.From)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	serviceInput.DateTo, err = time.Parse(time.RFC3339, inp.To)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	serviceInput.Symbol = inp.Ticker

	res, err := h.services.Tickers.Fetch(c.Context(), serviceInput)
	if err != nil {
		if errors.Is(err, models.ErrRatesNotFound) {
			return c.Status(fiber.StatusBadRequest).SendString(err.Error())
		}
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	return c.Status(fiber.StatusOK).JSON(res)
}
