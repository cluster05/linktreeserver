package model

import "gorm.io/plugin/soft_delete"

type Link struct {
	AuthId    string                `json:"authId" gorm:"column:authId"`
	LinkId    string                `json:"linkId" gorm:"column:linkId"`
	Title     string                `json:"title" gorm:"column:title"`
	URL       string                `json:"url" gorm:"column:url"`
	ImageUrl  string                `json:"imageUrl" gorm:"column:imageUrl"`
	CreatedAt int64                 `json:"createdAt" gorm:"column:createdAt"`
	UpdatedAt int64                 `json:"updatedAt" gorm:"column:updatedAt"`
	IsDeleted soft_delete.DeletedAt `json:"isDeleted" gorm:"column:isDeleted;softDelete:flag"`
}

type CreateLinkDTO struct {
	Title    string `json:"title" binding:"required"`
	URL      string `json:"url" binding:"required,url"`
	ImageUrl string `json:"imageUrl" binding:"required,url"`
}

type UpdateLinkDTO struct {
	LinkId   string `json:"linkId" binding:"required"`
	Title    string `json:"title" binding:"required"`
	URL      string `json:"url" binding:"required,url"`
	ImageUrl string `json:"imageUrl" binding:"required,url"`
}

type DeleteLinkDTO struct {
	LinkId string `json:"linkId" binding:"required"`
}
