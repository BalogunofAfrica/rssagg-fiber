package main

import (
	"github.com/gofiber/fiber/v2"
)

func handlerReadiness(c *fiber.Ctx) error {
	type Response struct {
		Message string `json:"message"`
	}

	return c.JSON(Response{Message: "Hello there"})
}
