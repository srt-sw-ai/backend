package service

import (
	"swai/dto"
	"swai/entity"

	"gorm.io/gorm"
)

type ReportsService struct {
	DB *gorm.DB
}

func NewReportsService(db *gorm.DB) *ReportsService {
	return &ReportsService{DB: db}
}

func (s *ReportsService) CreateReport(userId uint, createReportDto dto.CreateReportDto) error {
	report := entity.Report{
		Type:     createReportDto.Type,
		Title:    createReportDto.Title,
		Content:  createReportDto.Content,
		Location: createReportDto.Location,
		Date:     createReportDto.Date,
		UserID:   userId,
	}
	return s.DB.Create(&report).Error
}

func (s *ReportsService) FindAllReports() ([]entity.Report, error) {
	var reports []entity.Report
	err := s.DB.Find(&reports).Error
	return reports, err
}

func (s *ReportsService) FindReport(id uint) (*entity.Report, error) {
	var report entity.Report
	err := s.DB.First(&report, id).Error
	if err != nil {
		return nil, err
	}
	return &report, nil
}

func (s *ReportsService) FindReportByUserId(userId uint) ([]entity.Report, error) {
	var reports []entity.Report

	err := s.DB.Where("user_id = ?", userId).Find(&reports).Error
	if err != nil {
		return nil, err
	}

	return reports, nil
}
