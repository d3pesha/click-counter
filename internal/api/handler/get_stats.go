package handler

import (
	"counter/internal/model/api"
	"github.com/gofiber/fiber/v2"
	"log"
	"strconv"
)

func (h *handler) getStats(c *fiber.Ctx) error {
	bannerID, err := strconv.Atoi(c.Params("bannerID"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid banner ID"})
	}

	request := api.GetStats

	if err = c.BodyParser(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
	}

	stats, err := h.service.GetStatistics(c.Context(), bannerID, request.From, request.To)
	if err != nil {
		log.Println("failed to get stats:", err)

		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{"stats": stats})
}
