package main

import (
	"log"
	"user_service/db"
	"user_service/handlers"
	"user_service/repository"
	"user_service/routes"
	"user_service/services"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	app := fiber.New()
	app.Use(logger.New())
	app.Use(cors.New())

	dataBase, _ := db.ConnectDb()

	defer db.CloseMongo()

	userRepo := repository.NewUserRepository(dataBase)

	userService := services.NewUserService(userRepo)

	userHandler := handlers.NewUserHandler(userService)

	userRoute := routes.NewUserRoute(userHandler)

	userRoute.UserGroupApi(app)
	app.Get("*", func(c *fiber.Ctx) error {
		return c.SendString("no route")
	})

	log.Fatal(app.Listen(":8001"))
}
