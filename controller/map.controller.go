package controller

import (
	"strconv"

	"swai/dto"
	"swai/service"

	"github.com/gofiber/fiber/v2"
)

type MapController struct {
	mapService *service.MapService
}

func NewMapController(mapService *service.MapService) *MapController {
	return &MapController{mapService: mapService}
}

// CreateMarker godoc
// @Summary 마커 생성
// @Description 새로운 마커를 생성합니다.
// @Tags Map
// @Accept json
// @Produce json
// @Param createMarkerDto body dto.CreateMarkerDto true "마커 정보"
// @Success 201 {object} map[string]interface{} "마커가 성공적으로 생성되었습니다"
// @Failure 400 {object} map[string]interface{} "잘못된 요청입니다"
// @Failure 500 {object} map[string]interface{} "마커 생성에 실패했습니다"
// @Router /map [post]
func (c *MapController) CreateMarker(ctx *fiber.Ctx) error {
	var createMarkerDto dto.CreateMarkerDto
	if err := ctx.BodyParser(&createMarkerDto); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "잘못된 요청 입니다"})
	}

	userId := ctx.Locals("userId").(uint)
	result := c.mapService.CreateMarker(userId, createMarkerDto)
	return ctx.Status(result.Status).JSON(result.Data)
}

// FindAllMarker godoc
// @Summary 모든 마커 조회
// @Description 모든 마커를 조회합니다.
// @Tags Map
// @Produce json
// @Success 200 {object} map[string]interface{} "마커 목록"
// @Failure 500 {object} map[string]interface{} "마커 조회에 실패했습니다"
// @Router /map [get]
func (c *MapController) FindAllMarker(ctx *fiber.Ctx) error {
	result := c.mapService.FindAllMarker()
	return ctx.Status(result.Status).JSON(result.Data)
}

// FindMarker godoc
// @Summary 마커 조회
// @Description 마커 ID로 마커를 조회합니다.
// @Tags Map
// @Produce json
// @Param markerId path uint true "마커 ID"
// @Success 200 {object} map[string]interface{} "마커 정보"
// @Failure 400 {object} map[string]interface{} "잘못된 마커 ID입니다"
// @Failure 404 {object} map[string]interface{} "마커를 찾을 수 없습니다"
// @Router /map/{markerId} [get]
func (c *MapController) FindMarker(ctx *fiber.Ctx) error {
	markerID, err := strconv.ParseUint(ctx.Params("markerId"), 10, 32)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"success": false, "message": "잘못된 위치 ID입니다"})
	}

	result := c.mapService.FindMarker(uint(markerID))
	return ctx.Status(result.Status).JSON(result.Data)
}