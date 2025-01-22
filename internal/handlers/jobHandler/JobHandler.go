package jobHandler

import (
	"fmt"
	"strconv"

	"github.com/ehsan-ashik/go-job-tracker-api/database"
	"github.com/ehsan-ashik/go-job-tracker-api/internal/handlers/common"
	"github.com/ehsan-ashik/go-job-tracker-api/internal/model"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func CreateJob(ctx *fiber.Ctx) error {
	job := new(model.Job)
	err := ctx.BodyParser(job)
	if err != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"status":  "error",
			"message": fmt.Sprintf("Invalid Body. Error: %v", err.Error()),
			"data":    nil,
		})
	}

	//insert company
	if job.Company.Name != "" {
		err = database.DB.FirstOrCreate(&job.Company, model.Company{Name: job.Company.Name}).Error

		if err != nil {
			return ctx.Status(500).JSON(fiber.Map{
				"status":  "error",
				"message": fmt.Sprintf("Could not create Company. Error: %v", err.Error()),
				"data":    nil,
			})
		}
		job.CompanyID = job.Company.ID
	} else {
		return ctx.Status(500).JSON(fiber.Map{
			"status":  "error",
			"message": "Company is required.",
			"data":    nil,
		})
	}

	//insert job category
	if job.JobCategory.Name != "" {
		err = database.DB.FirstOrCreate(&job.JobCategory, model.JobCategory{Name: job.JobCategory.Name}).Error

		if err != nil {
			return ctx.Status(500).JSON(fiber.Map{
				"status":  "error",
				"message": fmt.Sprintf("Could not create Job Category. Error: %v", err.Error()),
				"data":    nil,
			})
		}
		job.JobCategoryID = job.JobCategory.ID
	} else {
		return ctx.Status(500).JSON(fiber.Map{
			"status":  "error",
			"message": "Job Category is required.",
			"data":    nil,
		})
	}

	//create job
	job.ID = uuid.New()

	if job.JobDescription.Description != "" {
		err = database.DB.Create(&job).Error
	} else {
		err = database.DB.Omit("JobDescription").Create(&job).Error
	}
	if err != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"status":  "error",
			"message": fmt.Sprintf("Could not create job. Error: %v", err.Error()),
			"data":    nil,
		})
	}

	return ctx.Status(200).JSON(fiber.Map{
		"status":  "success",
		"message": "Job Created Successfully",
		"data":    job,
	})
}

func GetJobs(ctx *fiber.Ctx) error {
	var jobs []model.Job
	err := database.DB.Scopes(common.Paginate(ctx), common.Sort(ctx), common.Filter(ctx)).Preload("Company").Preload("JobCategory").Find(&jobs).Error

	if err != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"status":  "error",
			"message": fmt.Sprintf("Could not fetch jobs. Error: %v", err.Error()),
			"data":    nil,
		})
	}

	//get row count
	var rowCount int64
	database.DB.Model(model.Job{}).Scopes(common.Filter(ctx)).Count(&rowCount)
	ctx.Response().Header.Add("X-Total-Rows", strconv.FormatInt(rowCount, 10))

	return ctx.Status(200).JSON(fiber.Map{
		"status":  "success",
		"message": "",
		"data":    jobs,
	})
}

func GetJobByID(ctx *fiber.Ctx) error {
	var job model.Job
	id := ctx.Params("id")
	err := database.DB.Preload("Company").Preload("JobCategory").Preload("JobDescription").Preload("Resume").First(&job, "id = ?", id).Error

	if err != nil {
		return ctx.Status(404).JSON(fiber.Map{
			"status":  "error",
			"message": fmt.Sprintf("Job not found. Error: %v", err.Error()),
			"data":    nil,
		})
	}

	return ctx.Status(200).JSON(fiber.Map{
		"status":  "success",
		"message": "",
		"data":    job,
	})
}

func UpdateJob(ctx *fiber.Ctx) error {
	var job model.Job
	id := ctx.Params("id")
	err := database.DB.First(&job, "id = ?", id).Error

	if err != nil {
		return ctx.Status(404).JSON(fiber.Map{
			"status":  "error",
			"message": fmt.Sprintf("Job not found. Error: %v", err.Error()),
			"data":    nil,
		})
	}

	var updatedJob model.Job
	err = ctx.BodyParser(&updatedJob)

	if err != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"status":  "error",
			"message": fmt.Sprintf("Invalid Body. Error: %v", err.Error()),
			"data":    nil,
		})
	}

	if updatedJob.Status != "" {
		job.Status = updatedJob.Status
	}
	if updatedJob.Excitement != nil && *updatedJob.Excitement != 0 {
		job.Excitement = updatedJob.Excitement
	}
	if updatedJob.IsReferred != job.IsReferred {
		job.IsReferred = updatedJob.IsReferred
	}
	if updatedJob.ReferredBy != nil && *updatedJob.ReferredBy != "" {
		job.ReferredBy = updatedJob.ReferredBy
	}
	if updatedJob.Location != nil && *updatedJob.Location != "" {
		job.Location = updatedJob.Location
	}
	if updatedJob.Position != "" {
		job.Position = updatedJob.Position
	}
	if updatedJob.ResonseDate != nil {
		job.ResonseDate = updatedJob.ResonseDate
	}
	if updatedJob.Remark != nil && *updatedJob.Remark != "" {
		job.Remark = updatedJob.Remark
	}

	err = database.DB.Save(&job).Error

	if err != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"status":  "error",
			"message": fmt.Sprintf("Could not update job. Error: %v", err.Error()),
			"data":    nil,
		})
	}

	return ctx.Status(200).JSON(fiber.Map{
		"status":  "success",
		"message": "Job Updated Successfully",
		"data":    job,
	})
}

func DeleteJob(ctx *fiber.Ctx) error {
	var job model.Job
	id := ctx.Params("id")
	err := database.DB.Preload("JobDescription").First(&job, "id = ?", id).Error

	if err != nil {
		return ctx.Status(404).JSON(fiber.Map{
			"status":  "error",
			"message": fmt.Sprintf("Job not found. Error: %v", err.Error()),
			"data":    nil,
		})
	}

	err = database.DB.Select("JobDescription").Delete(&job).Error

	if err != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"status":  "error",
			"message": fmt.Sprintf("Could not delete job. Error: %v", err.Error()),
			"data":    nil,
		})
	}

	return ctx.Status(200).JSON(fiber.Map{
		"status":  "success",
		"message": "Job Deleted Successfully",
		"data":    nil,
	})
}
