package article

import (
	// "ex01/routes/auth"
	"ex01/services/article"
	"log"

	"github.com/gofiber/fiber/v2"

	"ex01/models"
)

type ApiArticleRoute struct {
	service article.IService
	logger  *log.Logger
}

func New(service article.IService) *fiber.App {
	// ar := ApiArticleRoute{
	// 	service: service,
	// 	logger:  log.Default(),
	// }

	router := fiber.New()
	// router.Get("/", ar.getAll)
	// router.Get("/:id", ar.get)
	// router.Post("/",
		// auth.CheckAuthorization,
		// ar.create,
	// )

	return router
}

func (ar *ApiArticleRoute) getAll(c *fiber.Ctx) error {
	articles, err := ar.service.GetAll()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.JSON(articles)
}

// GetPostById godoc
// @Summary Show a single post
// @Description Get a single post by ID
// @Tags articles
// @Accept  json
// @Produce  json
// @Param id path int true "Post ID"
// @Success 200 {json} string "Post HTML"
// @Failure 404 {json} string "Not Found"
// @Router /api/articles/{id} [get]
func (ar *ApiArticleRoute) get(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil || id < 1 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid ID",
		})
	}

	article, err := ar.service.Get(uint(id))
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.JSON(article)
}

// CreatePost godoc
// @Summary Post a new article
// @Description Post a new article from the admin panel
// @Tags articles
// @Accept  application/x-www-form-urlencoded
// @Produce application/json
// @Param title formData string true "Article Title"
// @Param content formData string true "Article Content"
// @Success 201 {nil} nil "Created"
// @Router /api/articles [post]
func (ar *ApiArticleRoute) create(c *fiber.Ctx) error {
	var article struct {
		Title   string `form:"title"`
		Content string `form:"content"`
	}

	if err := c.BodyParser(&article); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	if err := ar.service.Create(&models.Article{
		Title:   article.Title,
		Content: article.Content,
	}); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "HERE " + err.Error(),
		})
	}

	c.Status(fiber.StatusCreated)
	return nil
}
