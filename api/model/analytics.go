package model

import (
	"github.com/lithammer/shortuuid"
	"gorm.io/gorm"
)

type UserAgent string

var (
	Mobile    UserAgent = "Mobile"
	Tablet    UserAgent = "Tablet"
	Desktop   UserAgent = "Desktop"
	SmartTV   UserAgent = "SmartTV"
	Raspberry UserAgent = "Raspberry"
	Bot       UserAgent = "Bot"
	OTHER     UserAgent = "Other"
)

type Analytics struct {
	AnalyticsId   string    `json:"analyticsId" gorm:"column:analyticsId" `
	LinkId        string    `json:"linkId" gorm:"column:linkId"`
	ContinentCode string    `json:"continentCode" gorm:"column:continentCode"`
	CountryCode   string    `json:"countryCode" gorm:"column:countryCode"`
	RegionCode    string    `json:"regionCode" gorm:"column:regionCode"`
	City          string    `json:"city" gorm:"column:city"`
	Pincode       string    `json:"pincode" gorm:"column:pincode"`
	Latitude      string    `json:"latitude" gorm:"column:latitude"`
	Longitude     string    `json:"longitude" gorm:"column:longitude"`
	UserAgent     UserAgent `json:"userAgent" gorm:"column:userAgent"`
	OS            string    `json:"os" gorm:"column:os"`
	CreatedAt     int64     `json:"createdAt" gorm:"column:createdAt"`
	UpdatedAt     int64     `json:"updatedAt" gorm:"column:updatedAt"`
}

func (analytics *Analytics) BeforeCreate(tx *gorm.DB) error {
	analytics.AnalyticsId = shortuuid.New()
	return nil
}

type CreateAnalyticsDTO struct {
	LinkId        string    `json:"linkId" binding:"required"`
	ContinentCode string    `json:"continentCode"`
	CountryCode   string    `json:"countryCode"`
	RegionCode    string    `json:"regionCode"`
	City          string    `json:"city"`
	Pincode       string    `json:"pincode"`
	Latitude      string    `json:"latitude"`
	Longitude     string    `json:"longitude"`
	UserAgent     UserAgent `json:"userAgent" binding:"useragent"`
	OS            string    `json:"os"`
}

type ReadAnalyticsDTO struct {
	AuthId string `json:"authId"`
}

type ReadAnalyticsResult struct {
	// TODO some complex struct
}
