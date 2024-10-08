package controller

import (
	"strconv"
	"swai/dto"
	"swai/service"

	"github.com/gofiber/fiber/v2"
)

type ReportsController struct {
	reportsService *service.ReportsService
}

func NewReportsController(reportsService *service.ReportsService) *ReportsController {
	return &ReportsController{reportsService: reportsService}
}

func (c *ReportsController) CreateReport(ctx *fiber.Ctx) error {
	var createReportDto dto.CreateReportDto
	if err := ctx.BodyParser(&createReportDto); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "잘못된 요청 입니다"})
	}

	userId := ctx.Locals("userId").(uint)
	result := c.reportsService.CreateReport(userId, createReportDto)
	return ctx.Status(result.Status).JSON(result.Data)
}

func (c *ReportsController) FindAllReports(ctx *fiber.Ctx) error {
	result := c.reportsService.FindAllReports()
	return ctx.Status(result.Status).JSON(result.Data)
}

func (c *ReportsController) FindReport(ctx *fiber.Ctx) error {
	reportId, err := strconv.ParseUint(ctx.Params("reportId"), 10, 32)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "잘못된 신고 ID입니다"})
	}

	result := c.reportsService.FindReport(uint(reportId))
	return ctx.Status(result.Status).JSON(result.Data)
}

func (c *ReportsController) FindReportByUserId(ctx *fiber.Ctx) error {
	userId, ok := ctx.Locals("userId").(uint)
	if !ok {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "유효하지 않은 사용자 ID입니다"})
	}

	result := c.reportsService.FindReportByUserId(userId)
	return ctx.Status(result.Status).JSON(result.Data)
}