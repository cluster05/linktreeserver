package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/cluster05/linktree/api/appresponse"
	"github.com/cluster05/linktree/api/model"
	"github.com/cluster05/linktree/api/service"
	"github.com/cluster05/linktree/pkg/requesthandler"
)

type AuthRoute interface {
	Register(*gin.Context)
	Login(*gin.Context)
	ForgotPassword(*gin.Context)
	ChangePassword(*gin.Context)
}

type authRoute struct {
	AuthService service.AuthService
}

type AuthRouteConfig struct {
	AuthService service.AuthService
}

func NewAuthRoute(config *AuthRouteConfig) AuthRoute {
	return &authRoute{
		AuthService: config.AuthService,
	}
}

// Login		 godoc
//
//	@Summary		Login in linktree
//	@Description	Login in linktree
//	@Tags			Authentication
//	@Accept			application/json
//	@Produce		application/json
//	@Param			request	body		model.LoginDTO						true	"Login request"
//	@Success		200		{object}	appresponse.Response{data=string}	"Success response"
//	@Router			/o/login [post]
func (ar *authRoute) Login(c *gin.Context) {

	var loginDTO model.LoginDTO
	if valid := requesthandler.BindData(c, &loginDTO); !valid {
		return
	}

	token, err := ar.AuthService.Login(loginDTO)
	if err != nil {
		c.JSON(http.StatusOK, appresponse.NewInternalError(err.Error()))
		return
	}

	c.JSON(http.StatusOK, appresponse.NewSuccess(token))
}

// Register		 godoc
//
//	@Summary		Create new account in linktree
//	@Description	Create new account in linktree
//	@Tags			Authentication
//	@Accept			application/json
//	@Produce		application/json
//	@Param			request	body		model.RegisterDTO					true	"Create user request"
//	@Success		200		{object}	appresponse.Response{data=string}	"Success response"
//	@Router			/o/register [post]
func (ar *authRoute) Register(c *gin.Context) {

	var registerDTO model.RegisterDTO
	if valid := requesthandler.BindData(c, &registerDTO); !valid {
		return
	}

	token, err := ar.AuthService.Register(registerDTO)
	if err != nil {
		c.JSON(http.StatusOK, appresponse.NewInternalError(err.Error()))
		return
	}

	c.JSON(http.StatusOK, appresponse.NewSuccess(token))
}

// ForgotPassword	godoc
//
//	@Summary		ForgotPassword of linktree
//	@Description	ForgotPassword of linktree
//	@Tags			Authentication
//	@Accept			application/json
//	@Produce		application/json
//	@Param			request	body		model.ForgotPasswordDTO				true	"Forgot user request"
//	@Success		200		{object}	appresponse.Response{data=string}	"Success response"
//	@Router			/o/forgotPassword [post]
func (ar *authRoute) ForgotPassword(c *gin.Context) {

	var forgotPasswordDTO model.ForgotPasswordDTO
	if valid := requesthandler.BindData(c, &forgotPasswordDTO); !valid {
		return
	}

	response, err := ar.AuthService.ForgotPassword(forgotPasswordDTO)
	if err != nil {
		c.JSON(http.StatusOK, appresponse.NewInternalError(err.Error()))
		return
	}

	c.JSON(http.StatusOK, appresponse.NewSuccess(response))

}

// ChangePassword	godoc
//
//	@Summary		ChangePassword of linktree
//	@Description	ChangePassword of linktree
//	@Tags			Authentication
//	@Accept			application/json
//	@Produce		application/json
//	@Param			Authorization	header	string	false	"Bearer token"
//	@Security		BearerAuth
//	@Param			request	body		model.ChangePasswordDTO				true	"Change user request"
//	@Success		200		{object}	appresponse.Response{data=string}	"Success response"
//	@Router			/r/changePassword [post]
func (ar *authRoute) ChangePassword(c *gin.Context) {
	var changePasswordDTO model.ChangePasswordDTO
	if valid := requesthandler.BindData(c, &changePasswordDTO); !valid {
		return
	}

	response, err := ar.AuthService.ChangePassword(changePasswordDTO)
	if err != nil {
		c.JSON(http.StatusOK, appresponse.NewInternalError(err.Error()))
		return
	}

	c.JSON(http.StatusOK, appresponse.NewSuccess(response))
}
