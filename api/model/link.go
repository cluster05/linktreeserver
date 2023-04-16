package model

type Link struct {
	AuthId    string `json:"authId" db:"authId"`
	LinkId    string `json:"linkId" db:"linkId"`
	Title     string `json:"title" db:"title"`
	URL       string `json:"url" db:"url"`
	ImageUrl  string `json:"imageUrl" db:"imageUrl"`
	CreatedAt int64  `json:"createdAt" db:"createdAt"`
	UpdatedAt int64  `json:"updatedAt" db:"updatedAt"`
	IsDeleted bool   `json:"isDeleted" db:"isDeleted"`
}

type CreateLinkDTO struct {
	Title    string `json:"title" binding:"required"`
	URL      string `json:"url" binding:"required"`
	ImageUrl string `json:"imageUrl" binding:"required"`
}

type UpdateLinkDTO struct {
	LinkId   string `json:"linkId" binding:"required"`
	Title    string `json:"title" binding:"required"`
	URL      string `json:"url" binding:"required"`
	ImageUrl string `json:"imageUrl" binding:"required"`
}

type DeleteLinkDTO struct {
	LinkId string `json:"linkId" binding:"required"`
}
