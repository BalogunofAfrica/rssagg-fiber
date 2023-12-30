package main

import (
	"fmt"
	"time"

	"github.com/balogunofafrica/rssagg/internal/database"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func (apiCfg *apiConfig) handlerCreateFeedFollow(c *fiber.Ctx, user database.User) error {
	type parameters struct {
		FeedID uuid.UUID `json:"feedId"`
	}

	params := parameters{}
	err := c.BodyParser(&params)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": fmt.Sprintf("Bad request %v", err),
		})
	}

	feedFollow, err := apiCfg.DB.CreateFeedFollow(c.Context(), database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		UserID:    user.ID,
		FeedID:    params.FeedID,
	})
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": fmt.Sprintf("Could not create feed follow %v", err),
		})
	}

	return c.Status(201).JSON(databaseFeedFollowToFeedFollow(feedFollow))
}

func (apiCfg *apiConfig) handlerGetFeedFollows(c *fiber.Ctx, user database.User) error {
	feedFollow, err := apiCfg.DB.GetFeedFollows(c.Context(), user.ID)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": fmt.Sprintf("Could not get feed follow %v", err),
		})
	}

	return c.Status(200).JSON(databaseFeedFollowsToFeedFollows(feedFollow))
}

func (apiCfg *apiConfig) handlerDeleteFeedFollow(c *fiber.Ctx, user database.User) error {
	feedFollowStr := c.Params("feedFollowId")
	feedFollowId, err := uuid.Parse(feedFollowStr)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": fmt.Sprintf("Bad request %v", err),
		})
	}

	_, err = apiCfg.DB.DeleteFeedFollow(c.Context(), database.DeleteFeedFollowParams{
		ID:     feedFollowId,
		UserID: user.ID,
	})
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": fmt.Sprintf("Could not delete feed follow %v", err),
		})
	}

	return c.Status(202).JSON(fiber.Map{
		"message": "Deleted successfully",
	})
}
