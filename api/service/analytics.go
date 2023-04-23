package service

import (
	"github.com/cluster05/linktree/api/model"
	"github.com/cluster05/linktree/api/repository"
)

type AnalyticsService interface {
	CreateAnalytics(model.CreateAnalyticsDTO) (bool, error)
	ReadAnalytics(model.JWTPayload) ([]model.Analytics, error)
}

type analyticsService struct {
	analyticsRepository repository.AnalyticsRepository
}

type AnalyticsServiceConfig struct {
	AnalyticsRepository repository.AnalyticsRepository
}

func NewAnalyticsService(config *AnalyticsServiceConfig) AnalyticsService {
	return &analyticsService{
		analyticsRepository: config.AnalyticsRepository,
	}
}

func (as *analyticsService) CreateAnalytics(createAnalyticsDTO model.CreateAnalyticsDTO) (bool, error) {
	analytics := model.Analytics{
		LinkId:        createAnalyticsDTO.LinkId,
		ContinentCode: createAnalyticsDTO.ContinentCode,
		RegionCode:    createAnalyticsDTO.RegionCode,
		City:          createAnalyticsDTO.City,
		Pincode:       createAnalyticsDTO.Pincode,
		Latitude:      createAnalyticsDTO.Latitude,
		Longitude:     createAnalyticsDTO.Longitude,
		UserAgent:     createAnalyticsDTO.UserAgent,
		OS:            createAnalyticsDTO.OS,
	}
	return as.analyticsRepository.CreateAnalytics(analytics)
}

func (as *analyticsService) ReadAnalytics(user model.JWTPayload) ([]model.Analytics, error) {
	return as.analyticsRepository.ReadAnalytics(user)
}
