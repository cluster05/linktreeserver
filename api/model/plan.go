package model

import (
	"github.com/lithammer/shortuuid"
	"gorm.io/gorm"
)

type PlanType string

type SubscriptionType string

var (
	Free  PlanType = "FREE"
	Basic PlanType = "BASIC"
	Pro   PlanType = "PRO"
)

var (
	Monthly SubscriptionType = "MONTHLY"
	Yearly  SubscriptionType = "YEARLY"
)

type Payment struct {
	PaymentId string `json:"paymentId,omitempty"`
}

type Plan struct {
	PlanId           string           `json:"planId" gorm:"planId"`
	AuthId           string           `json:"authId" gorm:"authId"`
	PlanType         PlanType         `json:"planType" gorm:"planType"`
	Start            int64            `json:"start" gorm:"start"`
	End              int64            `json:"end" gorm:"end"`
	SubscriptionType SubscriptionType `json:"subscriptionType" gorm:"subscriptionType"`
	CreateAt         int64            `json:"createAt" gorm:"createAt"`
	UpdatedAt        int64            `json:"updatedAt" gorm:"updatedAt"`
	PaymentId        string           `json:"paymentId" gorm:"paymentId"`
	ActiveStatus     bool             `json:"activeStatus" gorm:"activeStatus"`
}

func (plan *Plan) BeforeCreate(tx *gorm.DB) error {
	plan.PlanId = shortuuid.New()
	return nil
}

type CreatePlanDTO struct {
	PlanType         PlanType         `json:"planType" binding:"required,plantype"`
	SubscriptionType SubscriptionType `json:"subscriptionType" binding:"required,subscriptiontype"`
}
