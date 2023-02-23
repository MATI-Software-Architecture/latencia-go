package api

import (
	"github.com/gofiber/fiber/v2"
)

func Router() {
	api := fiber.New(fiber.Config{
		Prefork:           true,
		ReduceMemoryUsage: true,
		RequestMethods:    []string{"POST"},
		Network:           "tcp4",
		ServerHeader:      "GoLang API",
		AppName:           "GoLang API",
	})
	api.Post("/center", post_center)
	api.Listen(":3000")
}
