package handlers

import (
	"user_service/helpers"
	"user_service/models"
	"user_service/services"
	"user_service/utils"

	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	service *services.UserService
}

func NewUserHandler(handler *services.UserService) *UserHandler {
	return &UserHandler{handler}
}

func (h *UserHandler) RegisterUserHandler(c *fiber.Ctx) error {
	user := new(models.RegisterUser)

	if err := c.BodyParser(user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	if err := utils.ValidatorStruct(user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err,
		})
	}

	if err := h.service.RegisterUser(*user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "successfuly register user",
	})
}

func (h *UserHandler) LoginUserHandler(c *fiber.Ctx) error {
	user := new(models.LoginUser)

	if err := c.BodyParser(user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	if err := utils.ValidatorStruct(user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err,
		})
	}

	token, err := h.service.LoginUser(*user)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	res := helpers.ResponseJson("login success", token)

	return c.Status(fiber.StatusOK).JSON(res)
}
