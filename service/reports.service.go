package service

import (
	"swai/common"
	"swai/dto"
	"swai/entity"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type ReportsService struct {
	DB *gorm.DB
}

func NewReportsService(db *gorm.DB) *ReportsService {
	return &ReportsService{DB: db}
}

func (s *ReportsService) CreateReport(userId uint, createReportDto dto.CreateReportDto) common.ServiceResult {
	report := entity.Report{
		Type:      createReportDto.Type,
		Title:     createReportDto.Title,
		Content:   createReportDto.Content,
		UserID:    userId,
		Latitude:  createReportDto.Latitude,
		Longitude: createReportDto.Longitude,
	}

	if err := s.DB.Create(&report).Error; err != nil {
		return common.ServiceResult{
			Status: fiber.StatusInternalServerError,
			Data:   fiber.Map{"success": false, "message": "신고 생성에 실패했습니다"},
		}
	}

	return common.ServiceResult{
		Status: fiber.StatusCreated,
		Data:   fiber.Map{"success": true, "reportId": report.ID},
	}
}

func (s *ReportsService) FindAllReports() common.ServiceResult {
	var reports []entity.Report
	if err := s.DB.Find(&reports).Error; err != nil {
		return common.ServiceResult{
			Status: fiber.StatusInternalServerError,
			Data:   fiber.Map{"success": false, "message": "신고 조회에 실패했습니다"},
		}
	}

	return common.ServiceResult{
		Status: fiber.StatusOK,
		Data:   fiber.Map{"success": true, "body": reports},
	}
}

func (s *ReportsService) FindReport(id uint) common.ServiceResult {
	var report entity.Report
	if err := s.DB.First(&report, id).Error; err != nil {
		return common.ServiceResult{
			Status: fiber.StatusNotFound,
			Data:   fiber.Map{"success": false, "message": "신고를 찾을 수 없습니다"},
		}
	}

	return common.ServiceResult{
		Status: fiber.StatusOK,
		Data:   fiber.Map{"success": true, "body": report},
	}
}

func (s *ReportsService) FindReportByUserId(userId uint) common.ServiceResult {
	var reports []entity.Report
	if err := s.DB.Where("user_id = ?", userId).Find(&reports).Error; err != nil {
		return common.ServiceResult{
			Status: fiber.StatusInternalServerError,
			Data:   fiber.Map{"success": false, "message": "신고 조회에 실패했습니다"},
		}
	}

	return common.ServiceResult{
		Status: fiber.StatusOK,
		Data:   fiber.Map{"success": true, "body": reports},
	}
}