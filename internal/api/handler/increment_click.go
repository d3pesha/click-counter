package handler

import (
	"github.com/gofiber/fiber/v2"
	"log"
	"strconv"
)

func (h *handler) incrementClick(c *fiber.Ctx) error {
	bannerID, err := strconv.Atoi(c.Params("bannerID"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid banner ID"})
	}

	err = h.service.RegisterClick(c.Context(), bannerID)
	if err != nil {
		log.Println("Error while register click:", err)

		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{"message": "Click registered"})
}
