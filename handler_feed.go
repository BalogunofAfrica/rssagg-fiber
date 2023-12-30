package main

import (
	"fmt"
	"time"

	"github.com/balogunofafrica/rssagg/internal/database"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func (apiCfg *apiConfig) handlerCreateFeed(c *fiber.Ctx, user database.User) error {
	type parameters struct {
		Name string `json:"name"`
		Url  string `json:"url"`
	}

	params := parameters{}
	err := c.BodyParser(&params)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": fmt.Sprintf("Bad request %v", err),
		})
	}

	feed, err := apiCfg.DB.CreateFeed(c.Context(), database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      params.Name,
		Url:       params.Url,
		UserID:    user.ID,
	})
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": fmt.Sprintf("Could not create feed %v", err),
		})
	}

	return c.Status(201).JSON(databaseFeedToFeed(feed))
}

func (apiCfg *apiConfig) handlerGetFeeds(c *fiber.Ctx) error {
	feeds, err := apiCfg.DB.GetFeeds(c.Context())
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": fmt.Sprintf("Could not get feeds %v", err),
		})
	}

	return c.Status(200).JSON(databaseFeedsToFeeds(feeds))
}
