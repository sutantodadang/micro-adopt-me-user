package routes

import (
	"user_service/handlers"

	"github.com/gofiber/fiber/v2"
)

type userRoute struct {
	handler *handlers.UserHandler
}

func NewUserRoute(handler *handlers.UserHandler) *userRoute {
	return &userRoute{handler}
}

func (ur *userRoute) UserGroupApi(route *fiber.App) {

	user := route.Group("/api/v2/user")

	user.Post("/", ur.handler.RegisterUserHandler)

	user.Post("/login", ur.handler.LoginUserHandler)

	route.Get("/health", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"message": "i'am fine",
		})
	})
}
