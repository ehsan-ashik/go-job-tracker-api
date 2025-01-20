package main

import (
	"os"

	"github.com/ehsan-ashik/go-job-tracker-api/database"
	"github.com/ehsan-ashik/go-job-tracker-api/internal/filestorageservice"
	"github.com/ehsan-ashik/go-job-tracker-api/router"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	app := fiber.New()

	app.Use(cors.New())

	filestorageservice.ServiceClientSharedKey(os.Getenv("AZURE_ACCOUNT_NAME"), os.Getenv("AZURE_ACCESS_KEY"))
	database.ConnectDB()

	router.SetupRoutes(app)

	app.Listen(":3000")
}
