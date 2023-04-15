package repository

import (
	"github.com/jmoiron/sqlx"

	"github.com/cluster05/linktree/api/model"
)

type AuthRepository interface {
	FetchAuthByUsername(string) (model.Auth, error)
	FetchAuthByEmail(string) (model.Auth, error)
	Register(model.Auth) (model.Auth, error)
	ChangePassword(string, string) error
}

type authRepository struct {
	MySqlDB *sqlx.DB
}

type AuthRepositoryConfig struct {
	MySqlDB *sqlx.DB
}

func NewAuthRepository(config *AuthRepositoryConfig) AuthRepository {
	return &authRepository{
		MySqlDB: config.MySqlDB,
	}
}

func (repo *authRepository) FetchAuthByUsername(username string) (model.Auth, error) {
	findAuthByEmailQuery := `SELECT authId,username,firstname,lastname,email,password,authBy 
		FROM auth
		WHERE username=?  AND isDeleted=false;`

	auth := model.Auth{}
	err := repo.MySqlDB.Get(&auth, findAuthByEmailQuery, username)
	if err != nil {
		return model.Auth{}, err
	}

	return auth, err
}

func (repo *authRepository) FetchAuthByEmail(email string) (model.Auth, error) {
	findAuthByEmailQuery := `SELECT authId,username,firstname,lastname,email,password,authBy 
		FROM auth
		WHERE email=?  AND isDeleted=false;`

	auth := model.Auth{}
	err := repo.MySqlDB.Get(&auth, findAuthByEmailQuery, email)
	if err != nil {
		return model.Auth{}, err
	}

	return auth, err
}

func (repo *authRepository) Register(auth model.Auth) (model.Auth, error) {
	registerUserQuery := `INSERT INTO 
		auth(authId,username,firstname,lastname,email,password,authBy,createdAt,updatedAt,isDeleted) VALUES 
		(:authId,:username,:firstname,:lastname,:email,:password,:authBy,:createdAt,:updatedAt,:isDeleted);`

	_, err := repo.MySqlDB.NamedExec(registerUserQuery, auth)
	if err != nil {
		return model.Auth{}, err
	}
	return auth, nil
}

func (repo *authRepository) ChangePassword(newPassword string, authId string) error {
	changePasswordQuery := `UPDATE auth SET password=? 
		WHERE authId=? AND isDeleted=false;`
	_, err := repo.MySqlDB.Exec(changePasswordQuery, newPassword, authId)
	return err
}
