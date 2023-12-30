package main

import (
	"fmt"
	"time"

	"github.com/balogunofafrica/rssagg/internal/database"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func (apiCfg *apiConfig) handlerCreateUser(c *fiber.Ctx) error {
	type parameters struct {
		Name string `json:"name"`
	}
	params := &parameters{}

	err := c.BodyParser(params)

	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": fmt.Sprintf("Bad request %v", err),
		})
	}

	user, err := apiCfg.DB.CreateUser(c.Context(), database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      params.Name,
	})
	if err != nil {
		return c.Status(400).SendString(fmt.Sprintf("Could not create user %v", err))
	}

	return c.Status(201).JSON(databaseUserToUser(user))
}

func (apiCfg *apiConfig) handlerGetUser(c *fiber.Ctx, user database.User) error {
	return c.Status(200).JSON(databaseUserToUser(user))
}

func (apiCfg *apiConfig) handlerGetPostsForUser(c *fiber.Ctx, user database.User) error {
	posts, err := apiCfg.DB.GetPostsForUser(c.Context(), database.GetPostsForUserParams{
		UserID: user.ID,
		Limit:  10,
	})
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": fmt.Sprintf("Could not get posts %v", err),
		})
	}

	return c.Status(200).JSON(databasePostsToPosts(posts))
}
