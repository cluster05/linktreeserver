package model

type AuthType string

var (
	EmailAuthType    AuthType = "EMAIL"
	GoogleAuthType   AuthType = "GOOGLE"
	FacebookAuthType AuthType = "FACEBOOK"
	TwitterAuthType  AuthType = "TWITTER"
)

type Auth struct {
	AuthId    string   `json:"authId,omitempty" db:"authId"`
	Username  string   `json:"username,omitempty" db:"username"`
	Firstname string   `json:"firstname,omitempty" db:"firstname"`
	Lastname  string   `json:"lastname,omitempty" db:"lastname"`
	Email     string   `json:"email,omitempty" db:"email"`
	Password  string   `json:"password,omitempty" db:"password"`
	AuthType  AuthType `json:"authBy,omitempty" db:"authBy"`
	CreatedAt int64    `json:"createdAt,omitempty" db:"createdAt"`
	UpdatedAt int64    `json:"updatedAt,omitempty" db:"updatedAt"`
	IsDeleted bool     `json:"isDeleted" db:"isDeleted"`
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
