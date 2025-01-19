package companyHandler

import (
	"fmt"

	"github.com/ehsan-ashik/go-job-tracker-api/database"
	"github.com/ehsan-ashik/go-job-tracker-api/internal/model"
	"github.com/gofiber/fiber/v2"
)

func CreateCompany(ctx *fiber.Ctx) error {
	company := new(model.Company)
	err := ctx.BodyParser(company)
	if err != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"status":  "error",
			"message": fmt.Sprintf("Invalid Body. Error: %v", err.Error()),
			"data":    nil,
		})
	}

	//create company
	err = database.DB.Create(&company).Error

	if err != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"status":  "error",
			"message": fmt.Sprintf("Could not create Company. Error: %v", err.Error()),
			"data":    nil,
		})
	}

	return ctx.Status(200).JSON(fiber.Map{
		"status":  "success",
		"message": "Company Created Successfully",
		"data":    company,
	})
}

func GetCompanys(ctx *fiber.Ctx) error {
	var companies []model.Company
	database.DB.Find(&companies)

	return ctx.Status(200).JSON(fiber.Map{
		"status":  "success",
		"message": "",
		"data":    companies,
	})
}

func GetCompanyByID(ctx *fiber.Ctx) error {
	var company model.Company
	id := ctx.Params("id")
	err := database.DB.Preload("Jobs").First(&company, "id = ?", id).Error

	if err != nil {
		return ctx.Status(404).JSON(fiber.Map{
			"status":  "error",
			"message": fmt.Sprintf("Company not found. Error: %v", err.Error()),
			"data":    nil,
		})
	}

	return ctx.Status(200).JSON(fiber.Map{
		"status":  "success",
		"message": "",
		"data":    company,
	})
}

func UpdateCompany(ctx *fiber.Ctx) error {
	var company model.Company
	id := ctx.Params("id")
	err := database.DB.First(&company, "id = ?", id).Error

	if err != nil {
		return ctx.Status(404).JSON(fiber.Map{
			"status":  "error",
			"message": fmt.Sprintf("Company not found. Error: %v", err.Error()),
			"data":    nil,
		})
	}

	var updatedCompany model.Company
	err = ctx.BodyParser(&updatedCompany)

	if err != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"status":  "error",
			"message": fmt.Sprintf("Invalid Body. Error: %v", err.Error()),
			"data":    nil,
		})
	}

	if updatedCompany.Name != "" {
		company.Name = updatedCompany.Name
	}
	if updatedCompany.Excitement != nil {
		company.Excitement = updatedCompany.Excitement
	}

	if updatedCompany.CareerCiteLink != nil && *updatedCompany.CareerCiteLink != "" {
		company.CareerCiteLink = updatedCompany.CareerCiteLink
	}

	err = database.DB.Save(&company).Error

	if err != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"status":  "error",
			"message": fmt.Sprintf("Could not update company. Error: %v", err.Error()),
			"data":    nil,
		})
	}

	return ctx.Status(200).JSON(fiber.Map{
		"status":  "success",
		"message": "Company Updated Successfully",
		"data":    company,
	})
}

func DeleteCompany(ctx *fiber.Ctx) error {
	var company model.Company
	id := ctx.Params("id")
	err := database.DB.Preload("Jobs").First(&company, "id = ?", id).Error

	if err != nil {
		return ctx.Status(404).JSON(fiber.Map{
			"status":  "error",
			"message": fmt.Sprintf("Company not found. Error: %v", err.Error()),
			"data":    nil,
		})
	}

	if len(company.Jobs) > 0 {
		return ctx.Status(400).JSON(fiber.Map{
			"status":  "error",
			"message": "Company has jobs. Please delete the jobs first",
			"data":    nil,
		})
	}

	err = database.DB.Delete(&company).Error

	if err != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"status":  "error",
			"message": fmt.Sprintf("Could not delete company. Error: %v", err.Error()),
			"data":    nil,
		})
	}

	return ctx.Status(200).JSON(fiber.Map{
		"status":  "success",
		"message": "Company Deleted Successfully",
		"data":    nil,
	})
}
