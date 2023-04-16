package service

import (
	"github.com/cluster05/linktree/api/model"
	"github.com/cluster05/linktree/api/repository"
)

type PlanService interface {
	CreatePayment(model.Payment) (model.Payment, error)
	CreatePlan(model.JWTPayload, model.CreatePlanDTO) (model.Plan, error)
	ReadPlan(model.JWTPayload) ([]model.Plan, error)
}

type planService struct {
	planRepository repository.PlanRepository
}

type PlanServiceConfig struct {
	PlanRepository repository.PlanRepository
}

func NewPlanService(config *PlanServiceConfig) PlanService {
	return &planService{
		planRepository: config.PlanRepository,
	}
}

func (ps *planService) CreatePayment(payment model.Payment) (model.Payment, error) {
	return ps.planRepository.CreatePayment(payment)
}

func (ps *planService) CreatePlan(user model.JWTPayload, createPlanDTO model.CreatePlanDTO) (model.Plan, error) {
	// TODO: make payment, create payment entry
	payment := model.Payment{}
	ps.CreatePayment(payment)

	// TODO
	plan := model.Plan{
		PaymentId: payment.PaymentId,
	}
	return ps.planRepository.CreatePlan(plan)
}

func (ps *planService) ReadPlan(user model.JWTPayload) ([]model.Plan, error) {
	return ps.planRepository.ReadPlan(user.AuthId)
}
