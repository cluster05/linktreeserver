package model

import "gorm.io/plugin/soft_delete"

type AuthType string

var (
	EmailAuthType    AuthType = "EMAIL"
	GoogleAuthType   AuthType = "GOOGLE"
	FacebookAuthType AuthType = "FACEBOOK"
	TwitterAuthType  AuthType = "TWITTER"
)

type Auth struct {
	AuthId    string                `json:"authId,omitempty" gorm:"column:authId"`
	Username  string                `json:"username,omitempty" gorm:"column:username"`
	Firstname string                `json:"firstname,omitempty" gorm:"column:firstname"`
	Lastname  string                `json:"lastname,omitempty" gorm:"column:lastname"`
	Email     string                `json:"email,omitempty" gorm:"column:email"`
	Password  string                `json:"password,omitempty" gorm:"column:password"`
	AuthType  AuthType              `json:"authBy,omitempty" gorm:"column:authBy"`
	CreatedAt int64                 `json:"createdAt,omitempty" gorm:"column:createdAt"`
	UpdatedAt int64                 `json:"updatedAt,omitempty" gorm:"column:updatedAt"`
	IsDeleted soft_delete.DeletedAt `json:"isDeleted" gorm:"column:isDeleted;softDelete:flag"`
}

type RegisterDTO struct {
	Username  string `json:"username" binding:"required"`
	Firstname string `json:"firstname" binding:"required"`
	Lastname  string `json:"lastname" binding:"required"`
	Email     string `json:"email" binding:"required,email"`
	Password  string `json:"password" binding:"required"`
}

type LoginDTO struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type ForgotPasswordDTO struct {
	Email string `json:"email" binding:"required,email"`
}

type ChangePasswordDTO struct {
	Email       string `json:"email" binding:"required,email"`
	OldPassword string `json:"oldPassword" binding:"required"`
	NewPassword string `json:"newPassword" binding:"required"`
}
