package jobRoutes

import (
	"github.com/ehsan-ashik/go-job-tracker-api/internal/handlers/jobHandler"
	"github.com/gofiber/fiber/v2"
)

func SetupJobRoutes(router fiber.Router) {
	job := router.Group("/job")
	// Create job
	job.Post("/", jobHandler.CreateJob)

	// Get all jobs
	job.Get("/", jobHandler.GetJobs)

	// Get job by ID
	job.Get("/:id", jobHandler.GetJobByID)

	// Update job
	job.Put("/:id", jobHandler.UpdateJob)

	// Delete job
	job.Delete("/:id", jobHandler.DeleteJob)
}
