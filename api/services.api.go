package api

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/mati/latencia/cmd"
	"github.com/mati/latencia/schema"
)

func post_center(c *fiber.Ctx) error {
	start := time.Now()
	body := schema.Center{}
	if err := c.BodyParser(&body); err != nil {
		message := schema.Response{Status: "error", Message: err.Error()}
		return c.Status(fiber.StatusBadRequest).JSON(message)
	}

	cmd.PriorityCh <- schema.Task{Body: body, Start: start}

	message := schema.Response{Status: "ok", Message: "Center received"}
	return c.Status(fiber.StatusOK).JSON(message)
}
