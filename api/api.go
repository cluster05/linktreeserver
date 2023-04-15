package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	
	"github.com/cluster05/linktree/api/middleware"
	"github.com/cluster05/linktree/api/repository"
	"github.com/cluster05/linktree/api/routes"
	"github.com/cluster05/linktree/api/service"
	"github.com/cluster05/linktree/datasource"
)

func InitRouter() (*gin.Engine, error) {

	datasource, err := datasource.Init()
	if err != nil {
		return nil, err
	}

	router := gin.New()

	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	o := router.Group("/o")
	r := router.Group("/r", middleware.Auth)

	o.POST("/health", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{})
	})

	r.POST("/health", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{})
	})

	authRoute := setupAuthService(datasource)

	o.POST("/login", authRoute.Login)
	o.POST("/register", authRoute.Register)
	o.POST("/forgotPassword", authRoute.ForgotPassword)

	r.POST("/changePassword", authRoute.ChangePassword)

	return router, nil
}

func setupAuthService(datasource *datasource.DataSource) routes.AuthRoute {
	authRepository := repository.NewAuthRepository(&repository.AuthRepositoryConfig{MySqlDB: datasource.MySqlDB})
	authService := service.NewAuthService(&service.AuthServiceConfig{AuthRepository: authRepository})
	return routes.NewAuthRoute(&routes.AuthRouteConfig{AuthService: authService})
}
