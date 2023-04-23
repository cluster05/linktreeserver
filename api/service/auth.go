package service

import (
	"errors"
	"fmt"
	"gorm.io/gorm"

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
	ErrAccountAlreadyExists          = fmt.Errorf("account already exists with us. try to login")
	ErrUsernameAlreadyExists         = fmt.Errorf("uesrname is already taken. please choose different username")
	ErrorAccountNotExists            = fmt.Errorf("account not exists with us. try to register")
	ErrOldAndNewPasswordCannotBeSame = fmt.Errorf("old and new password cannot be same")
	ErrorInvalidCredentials          = fmt.Errorf("invalid credentails")
)

func NewAuthService(config *AuthServiceConfig) AuthService {
	return &authService{
		authRepository: config.AuthRepository,
	}
}

func (as authService) Register(registerDTO model.RegisterDTO) (string, error) {

	_, err := as.authRepository.FetchAuthByEmail(registerDTO.Email)
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return "", ErrAccountAlreadyExists
	}

	_, err = as.authRepository.FetchAuthByUsername(registerDTO.Username)
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return "", ErrUsernameAlreadyExists
	}

	hashPassword, err := hash.CreatePasswordHash(registerDTO.Password)
	if err != nil {
		return "", err
	}

	auth := model.Auth{
		Username:  registerDTO.Username,
		Firstname: registerDTO.Firstname,
		Lastname:  registerDTO.Lastname,
		Email:     registerDTO.Email,
		Password:  hashPassword,
		AuthMode:  model.EmailAuthMode,
	}

	auth, err = as.authRepository.Register(auth)
	if err != nil {
		return "", err
	}

	jwtToken, err := GenerateToken(auth)
	if err != nil {
		return "", err
	}

	return jwtToken, nil
}

func (as authService) Login(loginDTO model.LoginDTO) (string, error) {

	auth, err := as.authRepository.FetchAuthByEmail(loginDTO.Email)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return "", ErrorAccountNotExists
	}

	if valid := hash.CheckPasswordHash(loginDTO.Password, auth.Password); !valid {
		return "", ErrorInvalidCredentials
	}

	jwtToken, err := GenerateToken(auth)
	if err != nil {
		return "", err
	}

	return jwtToken, nil
}

func (as authService) ForgotPassword(forgotPasswordDTO model.ForgotPasswordDTO) (string, error) {
	_, err := as.authRepository.FetchAuthByEmail(forgotPasswordDTO.Email)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return "", ErrorAccountNotExists
	}
	// TODO : send email for change password
	return "email is send to register email id", nil
}

func (as authService) ChangePassword(changePasswordDTO model.ChangePasswordDTO) (string, error) {

	if changePasswordDTO.OldPassword == changePasswordDTO.NewPassword {
		return "", ErrOldAndNewPasswordCannotBeSame
	}

	auth, err := as.authRepository.FetchAuthByEmail(changePasswordDTO.Email)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return "", ErrorAccountNotExists
	}

	if valid := hash.CheckPasswordHash(changePasswordDTO.OldPassword, auth.Password); !valid {
		return "", ErrorInvalidCredentials
	}

	hashPassword, err := hash.CreatePasswordHash(changePasswordDTO.NewPassword)
	if err != nil {
		return "", err
	}

	err = as.authRepository.ChangePassword(hashPassword, auth.AuthId)
	if err != nil {
		return "", err
	}

	jwtToken, err := GenerateToken(auth)
	if err != nil {
		return "", err
	}

	return jwtToken, nil
}
