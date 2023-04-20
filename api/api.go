package api

import (
	"net/http"

	docs "github.com/cluster05/linktree/docs"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	swaggerfiles "github.com/swaggo/files"
	ginswagger "github.com/swaggo/gin-swagger"

	"github.com/cluster05/linktree/api/appresponse"
	"github.com/cluster05/linktree/api/middleware"
	"github.com/cluster05/linktree/api/repository"
	"github.com/cluster05/linktree/api/routes"
	"github.com/cluster05/linktree/api/service"
	"github.com/cluster05/linktree/datasource"
	"github.com/cluster05/linktree/pkg/customevalidator"
)

func InitRouter() (*gin.Engine, error) {

	ds, err := datasource.Init()
	if err != nil {
		return nil, err
	}

	router := gin.New()

	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		_ = v.RegisterValidation("useragent", customevalidator.UserAgentValidator)
		_ = v.RegisterValidation("plantype", customevalidator.PlanTypeValidator)
		_ = v.RegisterValidation("subscriptiontype", customevalidator.SubscriptionTypeValidator)
	}

	o := router.Group("/o")

	r := router.Group("/r", middleware.Auth)

	o.POST("/health", func(context *gin.Context) {
		context.JSON(http.StatusOK, appresponse.NewSuccess("up"))
	})

	authRoute := setupAuthRoute(ds)
	linkRoute := setupLinkRoute(ds)
	analyticsRoute := setupAnalyticsRoute(ds)
	planRoute := setupPlanRoute(ds)

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

	o.POST("/createAnalytics", analyticsRoute.CreateAnalytics)
	r.POST("/readAnalytics", analyticsRoute.ReadAnalytics)

	r.POST("/createPlan", planRoute.CreatePlan)
	r.POST("/readPlan", planRoute.ReadPlan)

	docs.SwaggerInfo.BasePath = ""
	router.GET("/docs/*any", ginswagger.WrapHandler(swaggerfiles.Handler))

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

func setupAnalyticsRoute(datasource *datasource.DataSource) routes.AnalyticsRoute {
	analyticsRepository := repository.NewAnalyticsRepository(&repository.AnalyticsRepositoryConfig{MySqlDB: datasource.MySqlDB})
	analyticsService := service.NewAnalyticsService(&service.AnalyticsServiceConfig{AnalyticsRepository: analyticsRepository})
	return routes.NewAnalyticsRoute(&routes.AnalyticsRouteConfig{AnalyticsService: analyticsService})
}

func setupPlanRoute(datasource *datasource.DataSource) routes.PlanRoute {
	planRepository := repository.NewPlanRepository(&repository.PlanRepositoryConfig{MySqlDB: datasource.MySqlDB})
	planService := service.NewPlanService(&service.PlanServiceConfig{PlanRepository: planRepository})
	return routes.NewPlanRoute(&routes.PlanRouteConfig{PlanService: planService})
}
