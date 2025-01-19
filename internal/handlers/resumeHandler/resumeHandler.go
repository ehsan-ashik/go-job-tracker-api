package resumeHandler

import (
	"fmt"
	"io"
	"strings"

	"github.com/ehsan-ashik/go-job-tracker-api/database"
	"github.com/ehsan-ashik/go-job-tracker-api/internal/filestorageservice"
	"github.com/ehsan-ashik/go-job-tracker-api/internal/model"
	"github.com/gofiber/fiber/v2"
)

func CreateResume(ctx *fiber.Ctx) error {
	file, err := ctx.FormFile("file")
	if err != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"status":  "error",
			"message": "File not found",
			"data":    nil,
		})
	}
	title := ctx.FormValue("title")
	title = strings.TrimSpace(title)
	remark := ctx.FormValue("remark")

	if title == "" {
		return ctx.Status(400).JSON(fiber.Map{
			"status":  "error",
			"message": "Title is required",
			"data":    nil,
		})
	}

	// upload file to blob storage
	buffer, err := file.Open()
	if err != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"status":  "error",
			"message": "Could not open file",
			"data":    nil,
		})
	}
	defer buffer.Close()

	content, _ := io.ReadAll(buffer)

	fileNameParts := strings.Split(file.Filename, ".")
	extention := fileNameParts[len(fileNameParts)-1]
	fileName := fmt.Sprintf("%v.%v", strings.ReplaceAll(strings.ToLower(title), " ", "_"), extention)

	var url string
	if filestorageservice.CheckIfBlobExists(fileName) {
		return ctx.Status(500).JSON(fiber.Map{
			"status":  "error",
			"message": "File with the same name already exists",
			"data":    nil,
		})
	}

	url, err = filestorageservice.UploadBlobFile(fileName, content)

	if err != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"status":  "error",
			"message": "Could not upload file. Error: " + err.Error(),
			"data":    nil,
		})
	}

	// create resume in the db
	resume := new(model.Resume)
	resume.Title = title
	resume.URL = url
	resume.Remark = &remark

	//create job
	err = database.DB.Create(&resume).Error

	if err != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"status":  "error",
			"message": fmt.Sprintf("Could not create resume. Error: %v", err.Error()),
			"data":    nil,
		})
	}

	return ctx.Status(200).JSON(fiber.Map{
		"status":  "success",
		"message": "Resume Created Successfully",
		"data":    resume,
	})
}

func GetResumes(ctx *fiber.Ctx) error {
	var resumes []model.Resume
	database.DB.Find(&resumes)

	return ctx.Status(200).JSON(fiber.Map{
		"status":  "success",
		"message": "",
		"data":    resumes,
	})
}

func GetResumeByID(ctx *fiber.Ctx) error {
	var resume model.Resume
	id := ctx.Params("id")
	err := database.DB.Preload("Jobs").First(&resume, "id = ?", id).Error

	if err != nil {
		return ctx.Status(404).JSON(fiber.Map{
			"status":  "error",
			"message": fmt.Sprintf("Resume not found. Error: %v", err.Error()),
			"data":    nil,
		})
	}

	return ctx.Status(200).JSON(fiber.Map{
		"status":  "success",
		"message": "",
		"data":    resume,
	})
}

func UpdateResume(ctx *fiber.Ctx) error {
	var resume model.Resume
	id := ctx.Params("id")
	err := database.DB.First(&resume, "id = ?", id).Error

	if err != nil {
		return ctx.Status(404).JSON(fiber.Map{
			"status":  "error",
			"message": fmt.Sprintf("Resume not found. Error: %v", err.Error()),
			"data":    nil,
		})
	}

	var updatedResume model.Resume
	err = ctx.BodyParser(&updatedResume)

	if err != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"status":  "error",
			"message": fmt.Sprintf("Invalid Body. Error: %v", err.Error()),
			"data":    nil,
		})
	}

	if updatedResume.Title != "" {
		resume.Title = updatedResume.Title
	}
	if updatedResume.Remark != nil && *updatedResume.Remark != "" {
		resume.Remark = updatedResume.Remark
	}

	err = database.DB.Save(&resume).Error

	if err != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"status":  "error",
			"message": fmt.Sprintf("Could not update resume. Error: %v", err.Error()),
			"data":    nil,
		})
	}

	return ctx.Status(200).JSON(fiber.Map{
		"status":  "success",
		"message": "Resume Updated Successfully",
		"data":    resume,
	})
}

func DeleteResume(ctx *fiber.Ctx) error {
	var resume model.Resume
	id := ctx.Params("id")
	err := database.DB.First(&resume, "id = ?", id).Error

	if err != nil {
		return ctx.Status(404).JSON(fiber.Map{
			"status":  "error",
			"message": fmt.Sprintf("Resume not found. Error: %v", err.Error()),
			"data":    nil,
		})
	}

	urlItems := strings.Split(resume.URL, "/")
	fileName := urlItems[len(urlItems)-1]

	if !filestorageservice.CheckIfBlobExists(fileName) {
		return ctx.Status(500).JSON(fiber.Map{
			"status":  "error",
			"message": "File not found in the storage",
			"data":    nil,
		})
	}

	err = filestorageservice.DeleteBlobFile(fileName)

	if err != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"status":  "error",
			"message": "Could not delete file. Error: " + err.Error(),
			"data":    nil,
		})
	}

	err = database.DB.Delete(&resume).Error

	if err != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"status":  "error",
			"message": fmt.Sprintf("Could not delete resume from DB. Error: %v", err.Error()),
			"data":    nil,
		})
	}

	return ctx.Status(200).JSON(fiber.Map{
		"status":  "success",
		"message": "Resume Deleted Successfully",
		"data":    nil,
	})
}
