package main

import (
	"github.com/ehsan-ashik/go-job-tracker-api/database"
	"github.com/ehsan-ashik/go-job-tracker-api/router"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	database.ConnectDB()

	router.SetupRoutes(app)

	app.Listen(":3000")
}
