package main

import (
	// "fmt"
	"fmt"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

var jwtKey = []byte("secret-key")

var users = map[string]string{
	"user1": "password1",
	"user2": "password2",
}

type IError struct {
	Field string      `json:"field"`
	Tag   string      `json:"tag"`
	Value interface{} `json:"value"`
}

type Claims struct {
	Password string `json:"password"`
	Username string `json:"username" validate:"required,min=3,max=12"`

	Email string `json:"email"`
	jwt.StandardClaims
}

var validate = validator.New()

func ValidateClaims(c *fiber.Ctx) error {
	var errors []*IError
	body := new(Claims)
	if err := c.BodyParser(body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	err := validate.Struct(body)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var el IError
			el.Field = err.Field()
			el.Tag = err.Tag()
			el.Value = err.Param()
			errors = append(errors, &el)
		}
		return c.Status(fiber.StatusBadRequest).JSON(errors)
	}

	return c.Next()
}

func Signin(c *fiber.Ctx) error {
	// c.Route().Name
	creds := new(Claims)
	if err := c.BodyParser(creds); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	expectedPassword, ok := users[creds.Username]
	if !ok || expectedPassword != creds.Password {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Invalid credentials",
		})
	}

	expirationTime := time.Now().Add(5 * time.Minute)
	claims := &Claims{
		Username: creds.Username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to generate token",
		})
	}

	return c.JSON(fiber.Map{
		"message": "Signin successful",
		"token":   tokenString,
	})
}

func TokenAuth(errorHandler fiber.Handler) fiber.Handler {
	return func(c *fiber.Ctx) error {
		authHeader := c.Get("Authorization")
		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			// fmt.Print(authHeader, "in if condition")
			return errorHandler(c)
		}
		token := strings.TrimPrefix(authHeader, "Bearer ")
		c.Locals("token", token)

		return c.Next()
	}
}

func main() {
	app := fiber.New()
	app.Hooks().OnName(func(r fiber.Route) error {
		fmt.Print("Method: " + r.Method + "\n")

		return nil
	})

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Thank god")
	})

	app.Post("/signin", ValidateClaims, Signin)

	// app.Use(TokenAuth())
	app.Use(TokenAuth(func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Unauthorized",
		})
	}))

	app.Get("/protected", func(c *fiber.Ctx) error {
		token := c.Locals("token").(string)
		return c.JSON(fiber.Map{
			"message": "Protected route",
			"token":   token,
		})
	})

	app.Listen(":3000")
}
