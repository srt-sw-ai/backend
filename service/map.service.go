package service

import (
	"swai/common"
	"swai/dto"
	"swai/entity"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type MapService struct {
	DB *gorm.DB
}

func NewMapService(db *gorm.DB) *MapService {
	return &MapService{DB: db}
}

func (s *MapService) CreateMarker(userId uint, createMarkerDto dto.CreateMarkerDto) common.ServiceResult {
	marker := entity.Map{
		Type:      createMarkerDto.Type,
		Latitude:  createMarkerDto.Latitude,
		Longitude: createMarkerDto.Longitude,
		ReportID:  createMarkerDto.ReportID,
		UserID:    userId,
	}

	if err := s.DB.Create(&marker).Error; err != nil {
		return common.ServiceResult{
			Status: fiber.StatusInternalServerError,
			Data:   fiber.Map{"success": false, "message": "위치 생성에 실패했습니다"},
		}
	}

	return common.ServiceResult{
		Status: fiber.StatusCreated,
		Data:   fiber.Map{"success": true},
	}
}

func (s *MapService) FindAllMarker() common.ServiceResult {
	var markers []entity.Map
	if err := s.DB.Find(&markers).Error; err != nil {
		return common.ServiceResult{
			Status: fiber.StatusInternalServerError,
			Data:   fiber.Map{"success": false, "message": "위치 조회에 실패했습니다"},
		}
	}

	return common.ServiceResult{
		Status: fiber.StatusOK,
		Data:   fiber.Map{"success": true, "body": markers},
	}
}

func (s *MapService) FindMarker(id uint) common.ServiceResult {
	var marker entity.Map
	if err := s.DB.First(&marker, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return common.ServiceResult{
				Status: fiber.StatusNotFound,
				Data:   fiber.Map{"success": false, "message": "위치를 찾을 수 없습니다"},
			}
		}
		return common.ServiceResult{
			Status: fiber.StatusInternalServerError,
			Data:   fiber.Map{"success": false, "message": "위치 조회 중 오류가 발생했습니다"},
		}
	}

	return common.ServiceResult{
		Status: fiber.StatusOK,
		Data:   fiber.Map{"success": true, "body": marker},
	}
}
