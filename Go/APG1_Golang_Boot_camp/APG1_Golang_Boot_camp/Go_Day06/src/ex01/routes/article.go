package routes

import (
	// "log"

	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
	"github.com/golang-jwt/jwt/v5"

	// "log"

	"ex01/config"
	"ex01/models"
	"ex01/services/article"
	"ex01/utils"
)

type ArticleHTMLRoute struct {
	service article.IService
}

func New(service article.IService) *fiber.App {
	ar := ArticleHTMLRoute{
		service: service,
	}

	view_engine := html.New("./templates", ".html")
	// log.Println(view_engine.Templates)

	view_engine.AddFunc("add", func(a, b int) int {
		return a + b
	})
	view_engine.AddFunc("sub", func(a, b int) int {
		return a - b
	})

	router := fiber.New(fiber.Config{
		Views: view_engine,
	})

	router.Get("/", func(c *fiber.Ctx) error {
		return c.Redirect("/articles")
	})
	router.Get("/articles", ar.indexHTML)
	router.Get("/articles/:id", ar.show_article)

	router.Get("/login", ar.checkIfLoggedIn, ar.loginHTML)
	router.Post("/login", ar.checkIfLoggedIn, ar.login)
	router.Post("/logout", ar.logout)

	router.Get("/post",  CheckAuthorization, ar.postHTML)
	router.Post("/post", CheckAuthorization, ar.post)

	return router
}

func (ar *ArticleHTMLRoute) indexHTML(c *fiber.Ctx) error {
	if q := c.Queries(); len(q) == 0 {
		return c.Redirect("/articles?page=1")
	}

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

	isLoggedIn := c.Locals("user") != nil

	return c.Render("index", fiber.Map{
		"Title":    "Blog",
		"Message":  "My blog",
		"LoggedIn": isLoggedIn,
		"Articles": articles,
		"Page":     page.Page,
		"HasPrev":  page.Page > 1,
		"HasNext":  len(articles) == limit,
	})
}

func (ar *ArticleHTMLRoute) show_article(c *fiber.Ctx) error {

	id, err := c.ParamsInt("id", 0)
	if err != nil || id < 1 {
		return c.Status(fiber.StatusBadRequest).SendString("invalid id")
	}

	article, err := ar.service.Get(uint(id))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	return c.Render("article", fiber.Map{
		"Title":   "Article: " + article.Title,
		"Article": article,
	})
}

func (ar *ArticleHTMLRoute) checkIfLoggedIn(c *fiber.Ctx) error {
	if user := c.Locals("user"); user != nil {
		return c.Redirect("/articles")
	}

	return c.Next()
}

func (ar *ArticleHTMLRoute) loginHTML(c *fiber.Ctx) error {
	return c.Render("login", fiber.Map{
		"Message": "",
	})
}

func (ar *ArticleHTMLRoute) login(c *fiber.Ctx) error {
	var user struct {
		Username string `form:"username"`
		Password string `form:"password"`
	}

	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).Render(
			"login",
			fiber.Map{"Message": err.Error()},
		)
	}

	is_admin := false
	for _, admin := range config.Config.AdminInfo.Admins {
		if admin.Username == user.Username && admin.Password == user.Password {
			is_admin = true
			break
		}
	}

	if !is_admin {
		return c.Status(fiber.StatusUnauthorized).Render(
			"login",
			fiber.Map{
				"Message": "Invalid username or password",
			},
		)
	}

	token, err := utils.CreateToken(jwt.MapClaims{"username": user.Username})

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).Render(
			"login",
			fiber.Map{"Message": err.Error()},
		)
	}

	c.Cookie(&fiber.Cookie{
		Name:     "token",
		Value:    token,
		HTTPOnly: true,
		Expires:  time.Now().Add(time.Minute),
	})

	return c.Status(fiber.StatusOK).Redirect("/articles")
}

func (ar *ArticleHTMLRoute) postHTML(c *fiber.Ctx) error {
	return c.Render("add_article", fiber.Map{
		"Message": "",
	})
}

func (ar *ArticleHTMLRoute) post(c *fiber.Ctx) error {
	var article struct {
		Title   string `form:"title"`
		Content string `form:"content"`
	}

	if err := c.BodyParser(&article); err != nil {
		return c.Status(fiber.StatusBadRequest).Render(
			"add_article",
			fiber.Map{"Message": err.Error()},
		)
	}

	if article.Title == "" || article.Content == "" {
		return c.Status(fiber.StatusBadRequest).Render(
			"add_article",
			fiber.Map{"Message": "Title and content are required"},
		)
	}

	if err := ar.service.Create(&models.Article{
		Title:   article.Title,
		Content: article.Content,
	}); err != nil {
		return c.Status(fiber.StatusInternalServerError).Render(
			"add_article",
			fiber.Map{"Message": err.Error()},
		)
	}

	return c.Status(fiber.StatusCreated).Render("add_article", fiber.Map{
		"Message": "Post created",
	})
}

func (ar *ArticleHTMLRoute) logout(c *fiber.Ctx) error {
	utils.ClearCookie(c, "token")
	return c.Redirect("/articles")
}

func CheckAuthorization(c *fiber.Ctx) error {
	if user := c.Locals("user"); user == nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Unauthorized",
		})
	}

	return c.Next()
}
