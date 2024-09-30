package controller

import (
	"fmt"
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
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "잘못된 요청 입니다",
		})
	}

	userId := ctx.Locals("userId").(uint)
	fmt.Println(userId)
	err := c.reportsService.CreateReport(userId, createReportDto)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": "신고 생성에 실패했습니다",
		})
	}

	return ctx.JSON(fiber.Map{
		"success": true,
	})
}

func (c *ReportsController) FindAllReports(ctx *fiber.Ctx) error {
	reports, err := c.reportsService.FindAllReports()
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": "신고 조회에 실패했습니다",
		})
	}

	return ctx.JSON(fiber.Map{
		"success": true,
		"body":    reports,
	})
}

func (c *ReportsController) FindReport(ctx *fiber.Ctx) error {
	reportId, err := strconv.ParseUint(ctx.Params("reportId"), 10, 32)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "잘못된 신고 ID입니다",
		})
	}

	report, err := c.reportsService.FindReport(uint(reportId))
	if err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"success": false,
			"message": "신고를 찾을 수 없습니다",
		})
	}

	return ctx.JSON(fiber.Map{
		"success": true,
		"body":    report,
	})
}

func (c *ReportsController) FindReportByUserId(ctx *fiber.Ctx) error {
	userId, ok := ctx.Locals("userId").(uint)
	if !ok {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "유효하지 않은 사용자 ID입니다",
		})
	}

	reports, err := c.reportsService.FindReportByUserId(userId)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": "신고 조회에 실패했습니다",
		})
	}

	return ctx.JSON(fiber.Map{
		"success": true,
		"body":    reports,
	})
}
