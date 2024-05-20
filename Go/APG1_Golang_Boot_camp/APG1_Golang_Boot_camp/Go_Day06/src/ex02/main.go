package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	_ "github.com/gofiber/fiber/v2/utils"
)

func main() {
	app := fiber.New(fiber.Config{
        Prefork:       true,
        CaseSensitive: true,
        StrictRouting: true,
        ServerHeader:  "Fiber",
        AppName: "Test App v1.0.1",
    })

    app.Get("/", func(c *fiber.Ctx) error {
        return fiber.NewError(782, "Custom error message")
    })
    app.Get("/:name?", func(c *fiber.Ctx) error {
        if c.Params("name") != "" {
            return c.SendString("Hello " + c.Params("name"))
            // => Hello john
        }
        return c.SendString("Where is john?")
    })

	log.Println("Server started on port 3000")
	log.Fatal(app.Listen(":3000"))

}
