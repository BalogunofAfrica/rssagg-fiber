package main

import (
	"github.com/gofiber/fiber/v2"
)

func handlerError(c *fiber.Ctx) error {
	type ErrorResponse struct {
		Error string `json:"error"`
	}

	return c.Status(400).JSON(ErrorResponse{Error: "Something went wrong"})
}
