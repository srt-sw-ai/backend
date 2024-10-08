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

// CreateReport godoc
// @Summary 신고 생성
// @Description 새로운 신고를 생성합니다.
// @Tags Reports
// @Accept json
// @Produce json
// @Param createReportDto body dto.CreateReportDto true "신고 정보"
// @Success 201 {object} map[string]interface{} "신고가 성공적으로 생성되었습니다"
// @Failure 400 {object} map[string]interface{} "잘못된 요청입니다"
// @Failure 500 {object} map[string]interface{} "신고 생성에 실패했습니다"
// @Router /reports [post]
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

// FindAllReports godoc
// @Summary 모든 신고 조회
// @Description 모든 신고를 조회합니다.
// @Tags Reports
// @Produce json
// @Success 200 {object} map[string]interface{} "신고 목록"
// @Failure 500 {object} map[string]interface{} "신고 조회에 실패했습니다"
// @Router /reports [get]
func (c *ReportsController) FindAllReports(ctx *fiber.Ctx) error {
	result := c.reportsService.FindAllReports()
	return ctx.Status(result.Status).JSON(result.Data)
}

// FindReport godoc
// @Summary 신고 조회
// @Description 신고 ID로 신고를 조회합니다.
// @Tags Reports
// @Produce json
// @Param reportId path uint true "신고 ID"
// @Success 200 {object} map[string]interface{} "신고 정보"
// @Failure 400 {object} map[string]interface{} "잘못된 신고 ID입니다"
// @Failure 404 {object} map[string]interface{} "신고를 찾을 수 없습니다"
// @Router /reports/{reportId} [get]
func (c *ReportsController) FindReport(ctx *fiber.Ctx) error {
	reportId, err := strconv.ParseUint(ctx.Params("reportId"), 10, 32)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "잘못된 신고 ID입니다"})
	}

	result := c.reportsService.FindReport(uint(reportId))
	return ctx.Status(result.Status).JSON(result.Data)
}

// FindReportByUserId godoc
// @Summary 사용자 신고 조회
// @Description 사용자 ID로 신고를 조회합니다.
// @Tags Reports
// @Produce json
// @Success 200 {object} map[string]interface{} "사용자 신고 목록"
// @Failure 400 {object} map[string]interface{} "유효하지 않은 사용자 ID입니다"
// @Failure 500 {object} map[string]interface{} "신고 조회에 실패했습니다"
// @Router /reports/by-user [get]
func (c *ReportsController) FindReportByUserId(ctx *fiber.Ctx) error {
	userId, ok := ctx.Locals("userId").(uint)
	if !ok {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "유효하지 않은 사용자 ID입니다"})
	}

	result := c.reportsService.FindReportByUserId(userId)
	return ctx.Status(result.Status).JSON(result.Data)
}