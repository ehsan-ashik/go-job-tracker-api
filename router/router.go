package router

import (
	"github.com/ehsan-ashik/go-job-tracker-api/internal/routes/companyRoutes"
	"github.com/ehsan-ashik/go-job-tracker-api/internal/routes/jobCategoryRoutes"
	"github.com/ehsan-ashik/go-job-tracker-api/internal/routes/jobDescriptionRoutes"
	"github.com/ehsan-ashik/go-job-tracker-api/internal/routes/jobRoutes"
	"github.com/ehsan-ashik/go-job-tracker-api/internal/routes/resumeRoutes"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api")

	// Job Routes
	jobRoutes.SetupJobRoutes(api)

	// Company Routes
	companyRoutes.SetupCompanyRoutes(api)

	// Resume Routes
	resumeRoutes.SetupResumeRoutes(api)

	// Job Description Routes
	jobDescriptionRoutes.SetupJobDescriptionRoutes(api)

	// Job Category Routes
	jobCategoryRoutes.SetupJobCategoryRoutes(api)
}
