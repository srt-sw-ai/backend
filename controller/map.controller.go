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

func (c *MapController) CreateMarker(ctx *fiber.Ctx) error {
	var createMarkerDto dto.CreateMarkerDto
	if err := ctx.BodyParser(&createMarkerDto); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "잘못된 요청 입니다"})
	}

	userId := ctx.Locals("userId").(uint)
	createMarkerDto.UserID = userId

	result := c.mapService.CreateMarker(&createMarkerDto)
	return ctx.Status(result.Status).JSON(result.Data)
}

func (c *MapController) FindAllMarker(ctx *fiber.Ctx) error {
	result := c.mapService.FindAllMarker()
	return ctx.Status(result.Status).JSON(result.Data)
}

func (c *MapController) FindMarker(ctx *fiber.Ctx) error {
	markerID, err := strconv.Atoi(ctx.Params("markerId"))
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "잘못된 위치 ID입니다"})
	}

	result := c.mapService.FindMarker(markerID)
	return ctx.Status(result.Status).JSON(result.Data)
}