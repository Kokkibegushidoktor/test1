package handlers

import "github.com/gofiber/fiber/v2"

type addTickerInput struct {
	Ticker string `json:"ticker" validate:"required"`
}

func (h *Handler) AddTicker(c *fiber.Ctx) error {
	var inp addTickerInput
	if err := c.BodyParser(&inp); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	return c.Status(fiber.StatusOK).JSON(map[string]string{"ticker": inp.Ticker})
}
