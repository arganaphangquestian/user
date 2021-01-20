package route

import (
	"github.com/arganaphangquestian/user/model"
	"github.com/arganaphangquestian/user/repository"
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

// New Route
func New(repository repository.UserRepository) *fiber.App {
	app := fiber.New()
	repo := &userRepository{repository}
	app.Get("/user", repo.users)
	app.Post("/register", repo.register)
	return app
}
