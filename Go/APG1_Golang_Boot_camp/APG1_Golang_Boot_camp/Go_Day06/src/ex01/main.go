package main

import (
	"log"

	swagger "github.com/arsmn/fiber-swagger/v2"
	"github.com/gofiber/fiber/v2"
	_ "github.com/lib/pq"

	"ex01/config"
	"ex01/database"
	_ "ex01/docs"
	article_repository "ex01/repositories/article"
	article_route "ex01/routes/article"
	auth_route "ex01/routes/auth"
	article_service "ex01/services/article"
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
	article_route := article_route.New(article_service)

	app.Mount("/articles", article_route)
	app.Mount("/auth", auth_route.New())
	app.Get("/swagger/*", swagger.HandlerDefault)

	log.Println("Server started on port 8888")
	log.Fatal(app.Listen("127.0.0.1:8888"))
}

func Init() {
	config.Init("admin_credentials.txt")
}
