package main

import (
	"log"
	"time"

	swagger "github.com/arsmn/fiber-swagger/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/logger"

	"ex01/config"
	"ex01/database"
	_ "ex01/docs"
	article_repository "ex01/repositories/article"
	html_route "ex01/routes"
	article_service "ex01/services/article"
	"ex01/utils"
)

// @title My Blog API
// @version 1.0
// @description This is a simple blog API with admin panel, markdown support, and pagination.
// @host localhost:8888
// @BasePath /
func main() {
	Init()

	app := fiber.New()

	database := database.New()
	article_repository := article_repository.New(database)
	article_service := article_service.New(article_repository)

	limiterConfig := limiter.Config{
		Max:        100,
		Expiration: 1 * time.Second,
		KeyGenerator: func(c *fiber.Ctx) string {
			return c.IP()
		},
		LimitReached: func(c *fiber.Ctx) error {
			return c.Status(fiber.StatusTooManyRequests).SendString("429 Too Many Requests")
		},
	}

	// Apply the rate limiter middleware
	app.Use("/*", limiter.New(limiterConfig))
	app.Use("/*", logger.New())
	app.Use("/*", utils.GetUserMiddleware)

	// Serve static files
	app.Static("/", "./public")

	// Routes
	app.Mount("/", html_route.New(article_service))
	app.Get("/swagger/*", swagger.HandlerDefault)

	log.Println("Server started on port 8888")
	log.Fatal(app.Listen("127.0.0.1:8888"))
}

func Init() {
	config.Init("admin_credentials.txt")
}
