package jobCategoryRoutes

import (
	"github.com/ehsan-ashik/go-job-tracker-api/internal/handlers/jobCategoryHandler"
	"github.com/gofiber/fiber/v2"
)

func SetupJobCategoryRoutes(router fiber.Router) {
	jobCategory := router.Group("/job_category")
	// Create jobCategory
	jobCategory.Post("/", jobCategoryHandler.CreateJobCategory)

	// Get all companies
	jobCategory.Get("/", jobCategoryHandler.GetJobCategories)

	// Get jobCategory by ID
	jobCategory.Get("/:id", jobCategoryHandler.GetJobCategoryByID)

	// Update jobCategory
	jobCategory.Put("/:id", jobCategoryHandler.UpdateJobCategory)

	// Delete jobCategory
	jobCategory.Delete("/:id", jobCategoryHandler.DeleteJobCategory)
}
