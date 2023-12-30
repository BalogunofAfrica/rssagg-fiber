package main

import (
	"fmt"

	"github.com/balogunofafrica/rssagg/internal/auth"
	"github.com/balogunofafrica/rssagg/internal/database"
	"github.com/gofiber/fiber/v2"
)

type authHandler func(*fiber.Ctx, database.User) error

func (cfg *apiConfig) middlewareAuth(handler authHandler) func(*fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		apiKey, err := auth.GetAPIKey(c)
		if err != nil {
			return c.Status(403).JSON(fiber.Map{
				"error": fmt.Sprintf("Auth error %v", err),
			})
		}

		user, err := cfg.DB.GetUserByAPIKey(c.Context(), apiKey)
		if err != nil {
			return c.Status(400).JSON(fiber.Map{
				"error": fmt.Sprintf("Could not get user %v", err),
			})
		}

		return handler(c, user)
	}
}
