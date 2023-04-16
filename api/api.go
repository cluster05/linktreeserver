package api

import (
	"github.com/cluster05/linktree/api/appresponse"
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
		context.JSON(http.StatusOK, appresponse.NewSuccess("up"))
	})

	authRoute := setupAuthRoute(datasource)
	linkRoute := setupLinkRoute(datasource)

	o.POST("/login", authRoute.Login)
	o.POST("/register", authRoute.Register)
	o.POST("/forgotPassword", authRoute.ForgotPassword)

	r.POST("/health", func(context *gin.Context) {
		context.JSON(http.StatusOK, appresponse.NewSuccess("up"))
	})

	r.POST("/changePassword", authRoute.ChangePassword)
	r.POST("/createLink", linkRoute.CreateLink)
	r.POST("/readLink", linkRoute.ReadLink)
	r.POST("/updateLink", linkRoute.UpdateLink)
	r.POST("/deleteLink", linkRoute.DeleteLink)

	return router, nil
}

func setupAuthRoute(datasource *datasource.DataSource) routes.AuthRoute {
	authRepository := repository.NewAuthRepository(&repository.AuthRepositoryConfig{MySqlDB: datasource.MySqlDB})
	authService := service.NewAuthService(&service.AuthServiceConfig{AuthRepository: authRepository})
	return routes.NewAuthRoute(&routes.AuthRouteConfig{AuthService: authService})
}

func setupLinkRoute(datasource *datasource.DataSource) routes.LinkRoute {
	linkRepository := repository.NewLinkRepository(&repository.LinkRepositoryConfig{MySqlDB: datasource.MySqlDB})
	linkService := service.NewLinkService(&service.LinkServiceConfig{LinkRepository: linkRepository})
	return routes.NewLinkRoute(&routes.LinkRouteConfig{LinkService: linkService})
}
