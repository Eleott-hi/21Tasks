package html

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
	// "log"

	"ex01/services/article"
	// "ex01/models"
	// "ex01/routes/auth"
)

type ArticleHTMLRoute struct {
	service article.IService
}

func New(service article.IService) *fiber.App {
	ar := ArticleHTMLRoute{
		service: service,
	}

	view_engine := html.New("./templates", ".html")
	view_engine.AddFunc("add", func(a, b int) int {
		return a + b
	})
	view_engine.AddFunc("sub", func(a, b int) int {
		return a - b
	})

	router := fiber.New(fiber.Config{
		Views: view_engine,
	})

	router.Get("/", ar.indexHTML)

	return router
}

func (ar *ArticleHTMLRoute) indexHTML(c *fiber.Ctx) error {
	var page struct {
		Page int `query:"page"`
	}

	if err := c.QueryParser(&page); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}
	if page.Page == 0 {
		return c.Status(fiber.StatusBadRequest).SendString("page must be greater than 0")
	}

	limit := 3
	offset := (page.Page - 1) * 3

	articles, err := ar.service.GetAll(offset, limit)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	return c.Render("index", fiber.Map{
		"Title":    "Blog",
		"Message":  "My blog",
		"Articles": articles,
		"Page":     page.Page,
	})
}
