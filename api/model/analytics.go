package model

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
	AnalyticsId   string    `json:"analyticsId" gorm:"analyticsId" `
	LinkId        string    `json:"linkId" gorm:"linkId"`
	ContinentCode string    `json:"continentCode" gorm:"continentCode"`
	CountryCode   string    `json:"countryCode" gorm:"countryCode"`
	RegionCode    string    `json:"regionCode" gorm:"regionCode"`
	City          string    `json:"city" gorm:"city"`
	Pincode       string    `json:"pincode" gorm:"pincode"`
	Latitude      string    `json:"latitude" gorm:"latitude"`
	Longitude     string    `json:"longitude" gorm:"longitude"`
	UserAgent     UserAgent `json:"userAgent" gorm:"userAgent"`
	OS            string    `json:"os" gorm:"os"`
	CreatedAt     int64     `json:"createdAt" gorm:"column:createdAt"`
	UpdatedAt     int64     `json:"updatedAt" gorm:"column:updatedAt"`
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
