package jobDescriptionHandler

import (
	"fmt"
	"strconv"

	"github.com/ehsan-ashik/go-job-tracker-api/database"
	"github.com/ehsan-ashik/go-job-tracker-api/internal/handlers/common"
	"github.com/ehsan-ashik/go-job-tracker-api/internal/model"
	"github.com/gofiber/fiber/v2"
)

func CreateJobDescription(ctx *fiber.Ctx) error {
	jobDescription := new(model.JobDescription)
	err := ctx.BodyParser(jobDescription)
	if err != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"status":  "error",
			"message": fmt.Sprintf("Invalid Body. Error: %v", err.Error()),
			"data":    nil,
		})
	}

	//create job description
	err = database.DB.Create(&jobDescription).Error

	if err != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"status":  "error",
			"message": fmt.Sprintf("Could not create Job Description. Error: %v", err.Error()),
			"data":    nil,
		})
	}

	return ctx.Status(200).JSON(fiber.Map{
		"status":  "success",
		"message": "Job Description Created Successfully",
		"data":    jobDescription,
	})
}

func GetJobDescriptions(ctx *fiber.Ctx) error {
	var jobDescriptions []model.JobDescription

	err := database.DB.Scopes(common.Paginate(ctx), common.Sort(ctx), common.Filter(ctx)).Find(&jobDescriptions).Error

	if err != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"status":  "error",
			"message": fmt.Sprintf("Could not fetch job descriptions. Error: %v", err.Error()),
			"data":    nil,
		})
	}

	//get row count
	var rowCount int64
	database.DB.Model(model.JobDescription{}).Scopes(common.Filter(ctx)).Count(&rowCount)
	ctx.Response().Header.Add("X-Total-Rows", strconv.FormatInt(rowCount, 10))

	return ctx.Status(200).JSON(fiber.Map{
		"status":  "success",
		"message": "",
		"data":    jobDescriptions,
	})
}

func GetJobDescriptionByID(ctx *fiber.Ctx) error {
	var jobDescription model.JobDescription
	id := ctx.Params("id")
	err := database.DB.Preload("Job").First(&jobDescription, "id = ?", id).Error

	if err != nil {
		return ctx.Status(404).JSON(fiber.Map{
			"status":  "error",
			"message": fmt.Sprintf("Job Description not found. Error: %v", err.Error()),
			"data":    nil,
		})
	}

	return ctx.Status(200).JSON(fiber.Map{
		"status":  "success",
		"message": "",
		"data":    jobDescription,
	})
}

func UpdateJobDescription(ctx *fiber.Ctx) error {
	var jobDescription model.JobDescription
	id := ctx.Params("id")
	err := database.DB.First(&jobDescription, "id = ?", id).Error

	if err != nil {
		return ctx.Status(404).JSON(fiber.Map{
			"status":  "error",
			"message": fmt.Sprintf("Job Description not found. Error: %v", err.Error()),
			"data":    nil,
		})
	}

	var updatedJobDescription model.JobDescription
	err = ctx.BodyParser(&updatedJobDescription)

	if err != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"status":  "error",
			"message": fmt.Sprintf("Invalid Body. Error: %v", err.Error()),
			"data":    nil,
		})
	}

	if updatedJobDescription.Description != "" {
		jobDescription.Description = updatedJobDescription.Description
	}

	err = database.DB.Save(&jobDescription).Error

	if err != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"status":  "error",
			"message": fmt.Sprintf("Could not update Job Description. Error: %v", err.Error()),
			"data":    nil,
		})
	}

	return ctx.Status(200).JSON(fiber.Map{
		"status":  "success",
		"message": "Job Description Updated Successfully",
		"data":    jobDescription,
	})
}

func DeleteJobDescription(ctx *fiber.Ctx) error {
	var jobDescription model.JobDescription
	id := ctx.Params("id")
	err := database.DB.First(&jobDescription, "id = ?", id).Error

	if err != nil {
		return ctx.Status(404).JSON(fiber.Map{
			"status":  "error",
			"message": fmt.Sprintf("Job Description not found. Error: %v", err.Error()),
			"data":    nil,
		})
	}

	err = database.DB.Delete(&jobDescription).Error

	if err != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"status":  "error",
			"message": fmt.Sprintf("Could not delete Job Description. Error: %v", err.Error()),
			"data":    nil,
		})
	}

	return ctx.Status(200).JSON(fiber.Map{
		"status":  "success",
		"message": "Job Description Deleted Successfully",
		"data":    nil,
	})
}
