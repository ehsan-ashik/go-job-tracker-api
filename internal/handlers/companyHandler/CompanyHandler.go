package companyHandler

import "github.com/gofiber/fiber/v2"

func CreateCompany(ctx *fiber.Ctx) error {
	return ctx.SendString("Create Company")
}

func GetCompanys(ctx *fiber.Ctx) error {
	return ctx.SendString("Get Companys")
}

func GetCompanyByID(ctx *fiber.Ctx) error {
	return ctx.SendString("Get Company By ID")
}

func UpdateCompany(ctx *fiber.Ctx) error {
	return ctx.SendString("Update Company By ID")
}

func DeleteCompany(ctx *fiber.Ctx) error {
	return ctx.SendString("Delete Company By ID")
}
