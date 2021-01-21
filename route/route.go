package route

import (
	"fmt"
	"strings"
	"time"

	"github.com/arganaphangquestian/user/model"
	"github.com/arganaphangquestian/user/repository"
	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
)

type userRepository struct {
	repo repository.UserRepository
}

func (r *userRepository) register(c *fiber.Ctx) error {
	p := new(model.InputUser)
	if err := c.BodyParser(p); err != nil {
		return c.Status(500).JSON(&fiber.Map{
			"success": false,
			"message": err,
		})
	}
	response, err := r.repo.Register(*p)
	if err != nil {
		return c.Status(500).JSON(&fiber.Map{
			"success": false,
			"message": err,
		})
	}
	return c.Status(201).JSON(&fiber.Map{
		"success": true,
		"message": "Register endpoint reached",
		"data":    response,
	})
}

func (r *userRepository) users(c *fiber.Ctx) error {
	response, err := r.repo.Users()
	if err != nil {
		return c.Status(500).JSON(&fiber.Map{
			"success": false,
			"message": err,
		})
	}
	return c.Status(201).JSON(&fiber.Map{
		"success": true,
		"message": "Get All Users endpoint reached",
		"data":    response,
	})
}

func (r *userRepository) login(c *fiber.Ctx) error {
	p := new(model.Login)
	if err := c.BodyParser(p); err != nil {
		return c.Status(500).JSON(&fiber.Map{
			"success": false,
			"message": err,
		})
	}
	user, err := r.repo.Login(*p)
	token, err := createToken(*user)
	if err != nil {
		return c.Status(500).JSON(&fiber.Map{
			"success": false,
			"message": err,
		})
	}
	return c.Status(201).JSON(&fiber.Map{
		"success": true,
		"message": "Login Successfully",
		"data":    token,
	})
}

func createToken(user model.User) (string, error) {
	claims := jwt.MapClaims{}
	claims["user"] = user
	claims["exp"] = time.Now().Add(time.Minute * 60 * 24 * 30).Unix() // 1 month
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := at.SignedString([]byte("MY_SUPER_SECRET_KEY"))
	if err != nil {
		return "", err
	}
	return token, nil
}

func (r *userRepository) dashboard(c *fiber.Ctx) error {
	authorizationHeader := c.Get("Authorization")
	if !strings.Contains(authorizationHeader, "Bearer") {
		return c.Status(403).JSON(&fiber.Map{
			"success": false,
			"message": "Authorization must be valid",
		})
	}
	user, err := extractToken(authorizationHeader)
	if err != nil {
		return c.Status(500).JSON(&fiber.Map{
			"success": false,
			"message": err,
		})
	}
	return c.Status(201).JSON(&fiber.Map{
		"success": true,
		"message": "DASHBOARD",
		"data":    user,
	})
}

func extractToken(authorizationHeader string) (interface{}, error) {
	claims := jwt.MapClaims{}
	tokenString := strings.Replace(authorizationHeader, "Bearer ", "", -1)
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		if method, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Signing method invalid")
		} else if method != jwt.SigningMethodHS256 {
			return nil, fmt.Errorf("Signing method invalid")
		}
		return []byte("MY_SUPER_SECRET_KEY"), nil
	})
	if err != nil {
		return nil, fmt.Errorf("Error : %s", err)
	}
	if !token.Valid {
		return nil, fmt.Errorf("Token not valid")
	}
	return claims["user"], nil
}

// New Route
func New(repository repository.UserRepository) *fiber.App {
	app := fiber.New()
	repo := &userRepository{repository}
	app.Get("/user", repo.users)
	app.Post("/register", repo.register)
	app.Post("/login", repo.login)
	app.Get("/dashboard", repo.dashboard)
	return app
}
