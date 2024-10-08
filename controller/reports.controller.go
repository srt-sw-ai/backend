package controller

import (
	"strconv"
	"swai/dto"
	"swai/service"

	"github.com/gofiber/fiber/v2"
)

type ReportsController struct {
	reportsService *service.ReportsService
	mapService     *service.MapService
}

func NewReportsController(reportsService *service.ReportsService, mapService *service.MapService) *ReportsController {
	return &ReportsController{
		reportsService: reportsService,
		mapService:     mapService,
	}
}

func (c *ReportsController) CreateReport(ctx *fiber.Ctx) error {
	var createReportDto dto.CreateReportDto
	if err := ctx.BodyParser(&createReportDto); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "잘못된 요청 입니다"})
	}

	userId := ctx.Locals("userId").(uint)
	result := c.reportsService.CreateReport(userId, createReportDto)

	if result.Status == fiber.StatusCreated {
		reportId, ok := result.Data.(fiber.Map)["reportId"].(uint)
		if !ok {
			return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "신고 ID를 가져오는데 실패했습니다"})
		}

		createMarkerDto := dto.CreateMarkerDto{
			Type:      createReportDto.Type,
			Latitude:  createReportDto.Latitude,
			Longitude: createReportDto.Longitude,
			ReportID:  int(reportId),
			UserID:    userId,
		}

		markerResult := c.mapService.CreateMarker(userId, createMarkerDto)
		if markerResult.Status != fiber.StatusCreated {
			return ctx.Status(markerResult.Status).JSON(markerResult.Data)
		}
	}

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