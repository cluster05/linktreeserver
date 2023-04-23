package model

import (
	"github.com/lithammer/shortuuid"
	"gorm.io/gorm"
	"gorm.io/plugin/soft_delete"
)

type AuthMode string

var (
	EmailAuthMode    AuthMode = "EMAIL"
	GoogleAuthMode   AuthMode = "GOOGLE"
	FacebookAuthMode AuthMode = "FACEBOOK"
	TwitterAuthMode  AuthMode = "TWITTER"
)

type Auth struct {
	AuthId    string                `json:"authId,omitempty" gorm:"column:authId"`
	Username  string                `json:"username,omitempty" gorm:"column:username"`
	Firstname string                `json:"firstname,omitempty" gorm:"column:firstname"`
	Lastname  string                `json:"lastname,omitempty" gorm:"column:lastname"`
	Email     string                `json:"email,omitempty" gorm:"column:email"`
	Password  string                `json:"password,omitempty" gorm:"column:password"`
	AuthMode  AuthMode              `json:"authMode,omitempty" gorm:"column:authMode"`
	CreatedAt int64                 `json:"createdAt,omitempty" gorm:"column:createdAt"`
	UpdatedAt int64                 `json:"updatedAt,omitempty" gorm:"column:updatedAt"`
	IsDeleted soft_delete.DeletedAt `json:"isDeleted" gorm:"column:isDeleted;softDelete:flag" swaggerignore:"true"`
}

func (auth *Auth) BeforeCreate(tx *gorm.DB) error {
	auth.AuthId = shortuuid.New()

	return nil
}

type RegisterDTO struct {
	Username  string `json:"username" binding:"required,gte=2,max=10"`
	Firstname string `json:"firstname" binding:"required,gte=2,max=25"`
	Lastname  string `json:"lastname" binding:"required,gte=2,max=25"`
	Email     string `json:"email" binding:"required,email,max=50"`
	Password  string `json:"password" binding:"required,gte=8,max=20"`
}

type LoginDTO struct {
	Email    string `json:"email" binding:"required,email,max=50"`
	Password string `json:"password" binding:"required,gte=8,max=20"`
}

type ForgotPasswordDTO struct {
	Email string `json:"email" binding:"required,email,max=50"`
}

type ChangePasswordDTO struct {
	Email       string `json:"email" binding:"required,email,max=50"`
	OldPassword string `json:"oldPassword" binding:"required,gte=8,max=20"`
	NewPassword string `json:"newPassword" binding:"required,gte=8,max=20"`
}
