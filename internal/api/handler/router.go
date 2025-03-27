package handler

import (
	"counter/internal/service"
	"github.com/gofiber/fiber/v2"
)

type handler struct {
	service service.BannerService
}

func Register(
	r fiber.Router,
	service service.BannerService,
) {
	h := &handler{service: service}

	r.Get("/counter/:bannerID", h.incrementClick)
	r.Post("/stats/:bannerID", h.getStats)
}
