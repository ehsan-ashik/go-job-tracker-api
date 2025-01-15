package companyRoutes

import (
	"github.com/ehsan-ashik/go-job-tracker-api/internal/handlers/companyHandler"
	"github.com/gofiber/fiber/v2"
)

func SetupCompanyRoutes(router fiber.Router) {
	company := router.Group("/company")
	// Create company
	company.Post("/", companyHandler.CreateCompany)

	// Get all companies
	company.Get("/", companyHandler.GetCompanys)

	// Get company by ID
	company.Get("/:id", companyHandler.GetCompanyByID)

	// Update company
	company.Put("/:id", companyHandler.UpdateCompany)

	// Delete company
	company.Delete("/:id", companyHandler.DeleteCompany)
}
