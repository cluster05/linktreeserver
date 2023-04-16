package repository

import (
	"gorm.io/gorm"

	"github.com/cluster05/linktree/api/model"
)

type PlanRepository interface {
	CreatePayment(model.Payment) (model.Payment, error)
	CreatePlan(model.Plan) (model.Plan, error)
	ReadPlan(string) ([]model.Plan, error)
}

type planRepository struct {
	MySqlDB *gorm.DB
}

type PlanRepositoryConfig struct {
	MySqlDB *gorm.DB
}

func NewPlanRepository(config *PlanRepositoryConfig) PlanRepository {
	return &planRepository{
		MySqlDB: config.MySqlDB,
	}
}

func (repo *planRepository) CreatePayment(payment model.Payment) (model.Payment, error) {
	result := repo.MySqlDB.Create(&payment)
	if result.Error != nil {
		return model.Payment{}, result.Error
	}
	return payment, nil
}

func (repo *planRepository) CreatePlan(plan model.Plan) (model.Plan, error) {
	result := repo.MySqlDB.Create(&plan)
	if result.Error != nil {
		return model.Plan{}, result.Error
	}
	return plan, nil
}

func (repo *planRepository) ReadPlan(userId string) ([]model.Plan, error) {
	plans := []model.Plan{}
	result := repo.MySqlDB.Where("authId=?", userId).Find(&plans)

	if result.Error != nil {
		return []model.Plan{}, result.Error
	}
	return plans, nil
}
