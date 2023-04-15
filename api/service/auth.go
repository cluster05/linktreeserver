package service

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/google/uuid"

	"github.com/cluster05/linktree/api/model"
	"github.com/cluster05/linktree/api/repository"
	"github.com/cluster05/linktree/pkg/hash"
)

type AuthService interface {
	Register(model.RegisterDTO) (string, error)
	Login(model.LoginDTO) (string, error)
	ForgotPassword(model.ForgotPasswordDTO) (string, error)
	ChangePassword(model.ChangePasswordDTO) (string, error)
}

type authService struct {
	authRepository repository.AuthRepository
}

type AuthServiceConfig struct {
	AuthRepository repository.AuthRepository
}

var (
	errAccountAlreadyExists          = fmt.Errorf("account already exists with us. try to login")
	errUsernameAlreadyExists         = fmt.Errorf("uesrname is already taken. please choose different username")
	errorAccountNotExists            = fmt.Errorf("account not exists with us. try to register")
	errOldAndNewPasswordCannotBeSame = fmt.Errorf("old and new password cannot be same")
	errorInvalidCredentials          = fmt.Errorf("invalid credentails")
)

func NewAuthService(config *AuthServiceConfig) AuthService {
	return &authService{
		authRepository: config.AuthRepository,
	}
}

func (as authService) Register(registerDTO model.RegisterDTO) (string, error) {

	auth, _ := as.authRepository.FetchAuthByEmail(registerDTO.Email)
	if auth.Email != "" {
		return "", errAccountAlreadyExists
	}

	auth, _ = as.authRepository.FetchAuthByUsername(registerDTO.Username)
	if auth.Username != "" {
		return "", errUsernameAlreadyExists
	}

	hashPassword, err := hash.HashPassword(registerDTO.Password)
	if err != nil {
		return "", err
	}

	auth = model.Auth{
		AuthId:    uuid.NewString(),
		Username:  registerDTO.Username,
		Firstname: registerDTO.Firstname,
		Lastname:  registerDTO.Lastname,
		Email:     registerDTO.Email,
		Password:  hashPassword,
		AuthType:  model.EmailAuthType,
		CreatedAt: time.Now().Unix(),
		UpdatedAt: time.Now().Unix(),
		IsDeleted: false,
	}

	auth, err = as.authRepository.Register(auth)
	if err != nil {
		return "", err
	}

	jwtToken, err := generateToken(auth)
	if err != nil {
		return "", err
	}

	return jwtToken, nil
}

func (as authService) Login(loginDTO model.LoginDTO) (string, error) {

	auth, err := as.authRepository.FetchAuthByEmail(loginDTO.Email)
	if err == sql.ErrNoRows {
		return "", errorAccountNotExists
	}

	if valid := hash.CheckPasswordHash(loginDTO.Password, auth.Password); !valid {
		return "", errorInvalidCredentials
	}

	jwtToken, err := generateToken(auth)
	if err != nil {
		return "", err
	}

	return jwtToken, nil
}

func (as authService) ForgotPassword(forgotPasswordDTO model.ForgotPasswordDTO) (string, error) {
	_, err := as.authRepository.FetchAuthByEmail(forgotPasswordDTO.Email)
	if err == sql.ErrNoRows {
		return "", errorAccountNotExists
	}
	// TODO : send email for change password
	return "email is send to register email id", nil
}

func (as authService) ChangePassword(changePasswordDTO model.ChangePasswordDTO) (string, error) {

	if changePasswordDTO.OldPassword == changePasswordDTO.NewPassword {
		return "", errOldAndNewPasswordCannotBeSame
	}

	auth, err := as.authRepository.FetchAuthByEmail(changePasswordDTO.Email)
	if err == sql.ErrNoRows {
		return "", errorAccountNotExists
	}

	if valid := hash.CheckPasswordHash(changePasswordDTO.OldPassword, auth.Password); !valid {
		return "", errorInvalidCredentials
	}

	hashPassword, err := hash.HashPassword(changePasswordDTO.NewPassword)
	if err != nil {
		return "", err
	}

	err = as.authRepository.ChangePassword(hashPassword, auth.AuthId)
	if err != nil {
		return "", err
	}

	jwtToken, err := generateToken(auth)
	if err != nil {
		return "", err
	}

	return jwtToken, nil
}