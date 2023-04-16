package repository

import (
	"gorm.io/gorm"

	"github.com/cluster05/linktree/api/model"
)

type AuthRepository interface {
	FetchAuthByUsername(string) (model.Auth, error)
	FetchAuthByEmail(string) (model.Auth, error)
	Register(model.Auth) (model.Auth, error)
	ChangePassword(string, string) error
}

type authRepository struct {
	MySqlDB *gorm.DB
}

type AuthRepositoryConfig struct {
	MySqlDB *gorm.DB
}

func NewAuthRepository(config *AuthRepositoryConfig) AuthRepository {
	return &authRepository{
		MySqlDB: config.MySqlDB,
	}
}

func (repo *authRepository) FetchAuthByUsername(username string) (model.Auth, error) {
	auth := model.Auth{}
	result := repo.MySqlDB.Where("username=?  AND isDeleted=false", username).Find(&auth)
	if result.Error != nil {
		return model.Auth{}, result.Error
	}

	return auth, nil
}

func (repo *authRepository) FetchAuthByEmail(email string) (model.Auth, error) {
	auth := model.Auth{}
	result := repo.MySqlDB.Where("email=? AND isDeleted=false", email).Find(&auth)
	if result.Error != nil {
		return model.Auth{}, result.Error
	}

	return auth, nil
}

func (repo *authRepository) Register(auth model.Auth) (model.Auth, error) {
	result := repo.MySqlDB.Create(&auth)
	if result.Error != nil {
		return model.Auth{}, result.Error
	}
	return auth, nil
}

func (repo *authRepository) ChangePassword(newPassword string, authId string) error {
	result := repo.MySqlDB.Model(&model.Auth{}).
		Where("authId=? AND isDeleted=false", authId).
		Update("password", newPassword)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
