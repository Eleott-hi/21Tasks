package auth

import (
	"ex01/config"
	"ex01/utils"
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

// AuthRoute represents the authentication route handler.
type AuthRoute struct {
	logger *log.Logger
}

func New() *fiber.App {
	ar := AuthRoute{
		logger: log.Default(),
	}

	router := fiber.New()
	router.Post("/login", ar.login)
	router.Post("/logout", ar.logout)

	return router
}

func GetUserMiddleware(c *fiber.Ctx) error {
	var token struct {
		Token string `cookie:"token"`
	}

	if err := c.CookieParser(&token); err != nil || token.Token == "" {
		c.ClearCookie("token")
	} else {
		payload, err := utils.DecodeJWT(token.Token)
		if err != nil {
			log.Println(err)
			c.ClearCookie("token")
		} else {
			c.Locals("user", payload["username"])
		}
	}

	return c.Next()
}

// @Summary Login
// @Description Authenticate user and generate a JWT token.
// @Tags auth
// @Accept application/x-www-form-urlencoded
// @Produce application/json
// @Param username formData string true "Username"
// @Param password formData string true "Password"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Router /auth/login [post]
func (ar *AuthRoute) login(c *fiber.Ctx) error {
	var user struct {
		Username string `form:"username"`
		Password string `form:"password"`
	}

	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	is_admin := false
	for _, admin := range config.Config.AdminInfo.Admins {
		if admin.Username == user.Username && admin.Password == user.Password {
			is_admin = true
			break
		}
	}

	if !is_admin {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Invalid username or password",
		})
	}

	token, err := utils.CreateToken(jwt.MapClaims{
		"username": user.Username,
	})
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	c.Cookie(&fiber.Cookie{
		Name:     "token",
		Value:    token,
		HTTPOnly: true,
		Expires:  time.Now().Add(time.Minute),
	})

	return nil
}

// logout handles user logout and token clearing.
// @Summary Logout
// @Description Clear the authentication token.
// @Tags auth
// @Success 200 {string} string "ok"
// @Router /auth/logout [post]
func (ar *AuthRoute) logout(c *fiber.Ctx) error {
	c.ClearCookie("token")
	c.Cookie(&fiber.Cookie{
		Name:     "token",
		Value:    "",
		HTTPOnly: true,
		Expires:  time.Unix(0, 0),
	})
	log.Println("User logged out")
	return nil
}
