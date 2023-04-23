package model

import (
	"github.com/lithammer/shortuuid"
	"gorm.io/gorm"
	"gorm.io/plugin/soft_delete"
)

type Link struct {
	AuthId    string                `json:"authId" gorm:"column:authId"`
	LinkId    string                `json:"linkId" gorm:"column:linkId"`
	Title     string                `json:"title" gorm:"column:title"`
	URL       string                `json:"url" gorm:"column:url"`
	ImageURL  string                `json:"ImageURL" gorm:"column:ImageURL"`
	CreatedAt int64                 `json:"createdAt" gorm:"column:createdAt"`
	UpdatedAt int64                 `json:"updatedAt" gorm:"column:updatedAt"`
	IsDeleted soft_delete.DeletedAt `json:"isDeleted" gorm:"column:isDeleted;softDelete:flag" swaggerignore:"true"`
}

func (link *Link) BeforeCreate(tx *gorm.DB) error {
	link.LinkId = shortuuid.New()
	return nil
}

type CreateLinkDTO struct {
	Title    string `json:"title" binding:"required,gte=2,max=30"`
	URL      string `json:"url" binding:"required,url"`
	ImageURL string `json:"ImageURL" binding:"required,url"`
}

type UpdateLinkDTO struct {
	LinkId   string `json:"linkId" binding:"required"`
	Title    string `json:"title" binding:"required,gte=2,max=30""`
	URL      string `json:"url" binding:"required,url"`
	ImageURL string `json:"ImageURL" binding:"required,url"`
}

type DeleteLinkDTO struct {
	LinkId string `json:"linkId" binding:"required"`
}
