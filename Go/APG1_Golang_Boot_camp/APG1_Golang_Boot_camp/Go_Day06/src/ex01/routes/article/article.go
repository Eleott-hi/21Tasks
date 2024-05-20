package article

import (
	"ex01/services/article"
	"github.com/gofiber/fiber/v2"
	"log"

	"ex01/models"
	"ex01/routes/auth"
)

type ArticleRoute struct {
	service article.IService
	logger  *log.Logger
}

func New(service article.IService) *fiber.App {
	ar := ArticleRoute{
		service: service,
		logger:  log.Default(),
	}

	router := fiber.New()
	router.Get("/", ar.getAll)
	router.Get("/:id", ar.get)
	router.Post("/",
		auth.GetUserMiddleware,
		ar.checkAuthorization,
		ar.create,
	)

	return router
}

func (ar *ArticleRoute) getAll(c *fiber.Ctx) error {
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
// @Router /articles/{id} [get]
func (ar *ArticleRoute) get(c *fiber.Ctx) error {
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
// @Router /articles [post]
func (ar *ArticleRoute) create(c *fiber.Ctx) error {
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

func (ar *ArticleRoute) checkAuthorization(c *fiber.Ctx) error {
	if user := c.Locals("user"); user == nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Unauthorized",
		})
	}

	return c.Next()
}
