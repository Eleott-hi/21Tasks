package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"time"

	swagger "github.com/arsmn/fiber-swagger/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/utils"
	"github.com/golang-jwt/jwt/v4"
	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	// "github.com/russross/blackfriday/v2"

	_ "ex01/docs"
)

var (
	db            *gorm.DB
	adminUsername string
	adminPassword string
	serviceSecret string
)

type (
	Article struct {
		gorm.Model
		Title   string `json:"title"`
		Content string `json:"content"`
	}

	User struct {
		gorm.Model
		Username string `json:"username"`
		Password string `json:"password"`
	}
)

func init() {
	file, err := os.Open("admin_credentials.txt")
	if err != nil {
		log.Fatal("Error opening admin credentials file:", err)
	}
	defer file.Close()

	var dbUser, dbPassword, dbName string
	_, err = fmt.Fscanf(file, "admin_username: %s\nadmin_password: %s\nservice_secret: %s\ndb_user: %s\ndb_password: %s\ndb_name: %s\n",
		&adminUsername, &adminPassword, &serviceSecret, &dbUser, &dbPassword, &dbName)
	if err != nil {
		log.Fatal("Error reading admin credentials:", err)
	}

	dsn := fmt.Sprintf("host=localhost user=%s password=%s dbname=%s port=5432 sslmode=disable TimeZone=Asia/Novosibirsk", dbUser, dbPassword, dbName)

	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	db.AutoMigrate(&Article{})
}

// @title My Blog API
// @version 1.0
// @description This is a simple blog API with admin panel, markdown support, and pagination.
// @host localhost:8888
// @BasePath /
func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})
	app.Get("/post/:id", GetPostById)
	app.Post("/post", CreatePost)
	app.Post("/login", Login)
	// app.GET("/admin", GetUser(Admin))
	// app.POST("/admin", GetUser(AdminPost))
	// app.GET("/login", Login)
	// app.POST("/login", LoginPost)
	app.Get("/swagger/*", swagger.HandlerDefault)
	// app.ServeFiles("/css/*filepath", http.Dir("css"))
	// app.ServeFiles("/images/*filepath", http.Dir("images"))
	// app.ServeFiles("/js/*filepath", http.Dir("js"))

	// Wrap the app with middleware
	// wrappedapp := applyMiddleware(app, GetUser)

	log.Println("Server started on port 8888")
	log.Fatal(app.Listen("127.0.0.1:8888"))
}

// GetPostById godoc
// @Summary Show a single post
// @Description Get a single post by ID
// @Tags public
// @Accept  json
// @Produce  html
// @Param id path int true "Post ID"
// @Success 200 {json} string "Post HTML"
// @Failure 404 {string} string "Not Found"
// @Router /post/{id} [get]
func GetPostById(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid ID")
	}

	if id < 1 {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid ID")
	}

	var article Article

	db.First(&article, 1)
	if article.ID == 0 {
		return c.Status(fiber.StatusNotFound).SendString("Not Found")
	}

	return c.Status(fiber.StatusOK).JSON(article)
}

// CreatePost godoc
// @Summary Post a new article
// @Description Post a new article from the admin panel
// @Tags admin
// @Accept  application/x-www-form-urlencoded
// @Produce application/json
// @Param title formData string true "Article Title"
// @Param content formData string true "Article Content"
// @Success 201 {} nil "Created"
// @Router /post [post]
func CreatePost(c *fiber.Ctx) error {

	var article struct {
		Title   string `form:"title"`
		Content string `form:"content"`
	}

	if err := c.BodyParser(&article); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	log.Default().Println(article)

	db.Create(
		&Article{
			Title:   utils.CopyString(article.Title),
			Content: utils.CopyString(article.Content),
		},
	)

	c.Status(fiber.StatusCreated)
	return nil
}

// Login godoc
// @Summary Login to admin panel
// @Description Login to the admin panel
// @Tags admin
// @Accept application/x-www-form-urlencoded
// @Produce text/plain
// @Param username formData string true "Admin Username"
// @Param password formData string true "Admin Password"
// @Success 200 {string} string "Logged in"
// @Router /login [post]
func Login(c *fiber.Ctx) error {
	var user struct {
		Username string `form:"username"`
		Password string `form:"password"`
	}

	if err := c.BodyParser(&user); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	is_admin := user.Username == adminUsername && user.Password == adminPassword
	token, err := generateJWT(jwt.MapClaims{
		"username": user.Username,
		"is_admin": is_admin,
	})

	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	c.Cookie(
		&fiber.Cookie{
			Name:     "token",
			Value:    token,
			HTTPOnly: true,
		},
	)

	return c.Status(fiber.StatusOK).SendString(token)
}

func generateJWT(payload jwt.MapClaims) (string, error) {
	if payload == nil {
		return "", errors.New("payload is nil")
	}

	payload["exp"] = time.Now().Add(time.Minute).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)

	tokenString, err := token.SignedString(serviceSecret)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
