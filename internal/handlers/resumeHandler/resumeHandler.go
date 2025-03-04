package resumeHandler

import (
	"fmt"
	"io"
	"strconv"
	"strings"

	"github.com/ehsan-ashik/go-job-tracker-api/database"
	"github.com/ehsan-ashik/go-job-tracker-api/internal/filestorageservice"
	"github.com/ehsan-ashik/go-job-tracker-api/internal/handlers/common"
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
	err := database.DB.Scopes(common.Paginate(ctx), common.Sort(ctx), common.Filter(ctx)).Find(&resumes).Error

	if err != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"status":  "error",
			"message": fmt.Sprintf("Could not fetch resumes. Error: %v", err.Error()),
			"data":    nil,
		})
	}

	//get row count
	var rowCount int64
	database.DB.Model(model.Resume{}).Scopes(common.Filter(ctx)).Count(&rowCount)
	ctx.Response().Header.Add("X-Total-Rows", strconv.FormatInt(rowCount, 10))

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

	updated_title := ctx.FormValue("title")
	updated_title = strings.TrimSpace(updated_title)
	if updated_title == "" {
		updated_title = resume.Title
	}
	updated_remark := ctx.FormValue("remark")
	var updated_url = resume.URL

	// upload file to blob storage
	file, err := ctx.FormFile("file")
	if err == nil {
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

		url_splits := strings.Split(resume.URL, "/")
		old_title := url_splits[len(url_splits)-1]

		if filestorageservice.CheckIfBlobExists(old_title) {
			err = filestorageservice.DeleteBlobFile(old_title)
			if err != nil {
				return ctx.Status(500).JSON(fiber.Map{
					"status":  "error",
					"message": "Could not delete old file. Error: " + err.Error(),
					"data":    nil,
				})
			}
		}

		new_title := fmt.Sprintf("%v.%v", strings.ReplaceAll(strings.ToLower(updated_title), " ", "_"), extention)
		updated_url, err = filestorageservice.UploadBlobFile(new_title, content)

		if err != nil {
			return ctx.Status(500).JSON(fiber.Map{
				"status":  "error",
				"message": "Could not upload new file. Error: " + err.Error(),
				"data":    nil,
			})
		}
	}

	if updated_title != resume.Title {
		resume.Title = updated_title
	}
	if updated_remark != *resume.Remark {
		*resume.Remark = updated_remark
	}
	if updated_url != resume.URL {
		resume.URL = updated_url
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

	if filestorageservice.CheckIfBlobExists(fileName) {
		err := filestorageservice.DeleteBlobFile(fileName)

		if err != nil {
			return ctx.Status(500).JSON(fiber.Map{
				"status":  "error",
				"message": "Could not delete file. Error: " + err.Error(),
				"data":    nil,
			})
		}
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
