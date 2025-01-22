package jobDescriptionRoutes

import (
	"github.com/ehsan-ashik/go-job-tracker-api/internal/handlers/jobDescriptionHandler"
	"github.com/gofiber/fiber/v2"
)

func SetupJobDescriptionRoutes(router fiber.Router) {
	jobDescription := router.Group("/job_description")
	// Create jobDescription
	jobDescription.Post("/", jobDescriptionHandler.CreateJobDescription)

	// Get all companies
	jobDescription.Get("/", jobDescriptionHandler.GetJobDescriptions)

	// Get jobDescription by ID
	jobDescription.Get("/:id", jobDescriptionHandler.GetJobDescriptionByID)

	// Update jobDescription
	jobDescription.Put("/:id", jobDescriptionHandler.UpdateJobDescription)

	// Delete jobDescription
	jobDescription.Delete("/:id", jobDescriptionHandler.DeleteJobDescription)
}
