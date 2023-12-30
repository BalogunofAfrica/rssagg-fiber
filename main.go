package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/balogunofafrica/rssagg/internal/database"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"

	_ "github.com/lib/pq"
)

type apiConfig struct {
	DB *database.Queries
}

func main() {
	godotenv.Load(".env")

	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("PORT is not found in the environment")
	}

	dbUrl := os.Getenv("DB_URL")
	if dbUrl == "" {
		log.Fatal("DB_URL is not found in the environment")
	}

	conn, err := sql.Open("postgres", dbUrl)
	if err != nil {
		log.Fatal("Fatal, can't connect to database", err)
	}

	apiConfig := apiConfig{DB: database.New(conn)}

	go startScrapping(apiConfig.DB, 10, time.Minute)

	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins:     "https://*, https://*",
		AllowMethods:     "GET, POST, PUT, DELETE, OPTIONS",
		AllowHeaders:     "Origin, Content-Type, Accept",
		ExposeHeaders:    "Link",
		AllowCredentials: false,
		MaxAge:           300,
	}))

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	v1 := app.Group("/v1")

	v1.Get("/health", handlerReadiness)
	v1.Get("/error", handlerError)

	v1.Post("/users", apiConfig.handlerCreateUser)
	v1.Get("/users", apiConfig.middlewareAuth(apiConfig.handlerGetUser))
	v1.Get("/users/posts", apiConfig.middlewareAuth(apiConfig.handlerGetPostsForUser))

	v1.Post("/feeds", apiConfig.middlewareAuth(apiConfig.handlerCreateFeed))
	v1.Get("/feeds", apiConfig.handlerGetFeeds)

	v1.Post("/feed_follows", apiConfig.middlewareAuth(apiConfig.handlerCreateFeedFollow))
	v1.Get("/feed_follows", apiConfig.middlewareAuth(apiConfig.handlerGetFeedFollows))
	v1.Delete("/feed_follows/:feedFollowId", apiConfig.middlewareAuth(apiConfig.handlerDeleteFeedFollow))

	fmt.Printf("server starting om port %v", port)
	err = app.Listen(":" + port)

	if err != nil {
		log.Fatal(err)
	}
}
