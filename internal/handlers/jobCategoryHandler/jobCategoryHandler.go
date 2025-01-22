package jobCategoryHandler

import (
	"fmt"
	"strconv"

	"github.com/ehsan-ashik/go-job-tracker-api/database"
	"github.com/ehsan-ashik/go-job-tracker-api/internal/handlers/common"
	"github.com/ehsan-ashik/go-job-tracker-api/internal/model"
	"github.com/gofiber/fiber/v2"
)

func CreateJobCategory(ctx *fiber.Ctx) error {
	jobCategory := new(model.JobCategory)
	err := ctx.BodyParser(jobCategory)
	if err != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"status":  "error",
			"message": fmt.Sprintf("Invalid Body. Error: %v", err.Error()),
			"data":    nil,
		})
	}

	//create job category
	err = database.DB.Create(&jobCategory).Error

	if err != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"status":  "error",
			"message": fmt.Sprintf("Could not create Job Category. Error: %v", err.Error()),
			"data":    nil,
		})
	}

	return ctx.Status(200).JSON(fiber.Map{
		"status":  "success",
		"message": "Job Category Created Successfully",
		"data":    jobCategory,
	})
}

func GetJobCategories(ctx *fiber.Ctx) error {
	var jobCategories []model.JobCategory

	err := database.DB.Scopes(common.Paginate(ctx), common.Sort(ctx), common.Filter(ctx)).Find(&jobCategories).Error

	if err != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"status":  "error",
			"message": fmt.Sprintf("Could not fetch job categories. Error: %v", err.Error()),
			"data":    nil,
		})
	}

	//get row count
	var rowCount int64
	database.DB.Model(model.JobCategory{}).Scopes(common.Filter(ctx)).Count(&rowCount)
	ctx.Response().Header.Add("X-Total-Rows", strconv.FormatInt(rowCount, 10))

	return ctx.Status(200).JSON(fiber.Map{
		"status":  "success",
		"message": "",
		"data":    jobCategories,
	})
}

func GetJobCategoryByID(ctx *fiber.Ctx) error {
	var jobCategory model.JobCategory
	id := ctx.Params("id")
	err := database.DB.Preload("Jobs").First(&jobCategory, "id = ?", id).Error

	if err != nil {
		return ctx.Status(404).JSON(fiber.Map{
			"status":  "error",
			"message": fmt.Sprintf("Job Category not found. Error: %v", err.Error()),
			"data":    nil,
		})
	}

	return ctx.Status(200).JSON(fiber.Map{
		"status":  "success",
		"message": "",
		"data":    jobCategory,
	})
}

func UpdateJobCategory(ctx *fiber.Ctx) error {
	var jobCategory model.JobCategory
	id := ctx.Params("id")
	err := database.DB.First(&jobCategory, "id = ?", id).Error

	if err != nil {
		return ctx.Status(404).JSON(fiber.Map{
			"status":  "error",
			"message": fmt.Sprintf("Job Category not found. Error: %v", err.Error()),
			"data":    nil,
		})
	}

	var updatedJobCategory model.JobCategory
	err = ctx.BodyParser(&updatedJobCategory)

	if err != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"status":  "error",
			"message": fmt.Sprintf("Invalid Body. Error: %v", err.Error()),
			"data":    nil,
		})
	}

	if updatedJobCategory.Name != "" {
		jobCategory.Name = updatedJobCategory.Name
	}
	if updatedJobCategory.Description != nil || *updatedJobCategory.Description != "" {
		jobCategory.Description = updatedJobCategory.Description
	}

	err = database.DB.Save(&jobCategory).Error

	if err != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"status":  "error",
			"message": fmt.Sprintf("Could not update Job Category. Error: %v", err.Error()),
			"data":    nil,
		})
	}

	return ctx.Status(200).JSON(fiber.Map{
		"status":  "success",
		"message": "Job Category Updated Successfully",
		"data":    jobCategory,
	})
}

func DeleteJobCategory(ctx *fiber.Ctx) error {
	var jobCategory model.JobCategory
	id := ctx.Params("id")
	err := database.DB.Preload("Jobs").First(&jobCategory, "id = ?", id).Error

	if err != nil {
		return ctx.Status(404).JSON(fiber.Map{
			"status":  "error",
			"message": fmt.Sprintf("Job Category not found. Error: %v", err.Error()),
			"data":    nil,
		})
	}

	if len(jobCategory.Jobs) > 0 {
		return ctx.Status(400).JSON(fiber.Map{
			"status":  "error",
			"message": "Job Category has jobs. Please delete the jobs first",
			"data":    nil,
		})
	}

	err = database.DB.Delete(&jobCategory).Error

	if err != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"status":  "error",
			"message": fmt.Sprintf("Could not delete Job Category. Error: %v", err.Error()),
			"data":    nil,
		})
	}

	return ctx.Status(200).JSON(fiber.Map{
		"status":  "success",
		"message": "Job Category Deleted Successfully",
		"data":    nil,
	})
}
