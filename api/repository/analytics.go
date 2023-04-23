package repository

import (
	"github.com/cluster05/linktree/api/model"
	"gorm.io/gorm"
)

type AnalyticsRepository interface {
	CreateAnalytics(model.Analytics) (bool, error)
	ReadAnalytics(model.JWTPayload) ([]model.Analytics, error)
}

type analyticsRepository struct {
	MySqlDB *gorm.DB
}

type AnalyticsRepositoryConfig struct {
	MySqlDB *gorm.DB
}

func NewAnalyticsRepository(config *AnalyticsRepositoryConfig) AnalyticsRepository {
	return &analyticsRepository{
		MySqlDB: config.MySqlDB,
	}
}

func (repo *analyticsRepository) CreateAnalytics(analytics model.Analytics) (bool, error) {
	result := repo.MySqlDB.Create(&analytics)
	if result.Error != nil {
		return false, result.Error
	}
	return true, nil
}

func (repo *analyticsRepository) ReadAnalytics(user model.JWTPayload) ([]model.Analytics, error) {
	analytics := []model.Analytics{}

	result := repo.MySqlDB.Model(&model.Analytics{}).
		Select("linkId").Where("linkId IN (?)",
		repo.MySqlDB.Model(&model.Link{}).Select("linkId").Where("authId=?", user.AuthId)).Find(&analytics)
	if result.Error != nil {
		return []model.Analytics{}, result.Error
	}
	return analytics, nil
}
