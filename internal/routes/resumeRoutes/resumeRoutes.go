package resumeRoutes

import (
	"github.com/ehsan-ashik/go-job-tracker-api/internal/handlers/resumeHandler"
	"github.com/gofiber/fiber/v2"
)

func SetupResumeRoutes(router fiber.Router) {
	resume := router.Group("/resume")
	// Create job
	resume.Post("/", resumeHandler.CreateResume)

	// Get all jobs
	resume.Get("/", resumeHandler.GetResumes)

	// Get job by ID
	resume.Get("/:id", resumeHandler.GetResumeByID)

	// Update job
	resume.Put("/:id", resumeHandler.UpdateResume)

	// Delete job
	resume.Delete("/:id", resumeHandler.DeleteResume)

}
