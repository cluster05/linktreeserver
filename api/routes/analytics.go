package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/cluster05/linktree/api/appresponse"
	"github.com/cluster05/linktree/api/model"
	"github.com/cluster05/linktree/api/query"
	"github.com/cluster05/linktree/api/service"
	"github.com/cluster05/linktree/pkg/requesthandler"
)

type AnalyticsRoute interface {
	CreateAnalytics(*gin.Context)
	ReadAnalytics(*gin.Context)
}

type analyticsRoute struct {
	analyticsService service.AnalyticsService
}

type AnalyticsRouteConfig struct {
	AnalyticsService service.AnalyticsService
}

func NewAnalyticsRoute(config *AnalyticsRouteConfig) AnalyticsRoute {
	return &analyticsRoute{
		analyticsService: config.AnalyticsService,
	}
}

// CreateAnalytics		 godoc
//
//	@Summary		Create new Analytics in linktree
//	@Description	Create new Analytics in linktree
//	@Tags			Analytics
//	@Accept			application/json
//	@Produce		application/json
//	@Param			request	body		model.CreateAnalyticsDTO		true	"CreateAnalytics request"
//	@Success		200		{object}	appresponse.Response{data=bool}	"Success response"
//	@Router			/o/createAnalytics [post]
func (ar *analyticsRoute) CreateAnalytics(c *gin.Context) {
	var createAnalyticsDTO model.CreateAnalyticsDTO
	if valid := requesthandler.BindData(c, &createAnalyticsDTO); !valid {
		return
	}

	isSaved, err := ar.analyticsService.CreateAnalytics(createAnalyticsDTO)
	if err != nil {
		c.JSON(http.StatusOK, appresponse.NewBadRequestError(err.Error()))
		return
	}

	c.JSON(http.StatusOK, isSaved)
}

// ReadAnalytics		 godoc
//
//	@Summary		Read new Analytics in linktree
//	@Description	Read new Analytics in linktree
//	@Tags			Analytics
//	@Accept			application/json
//	@Produce		application/json
//	@Param			Authorization	header	string	false	"Bearer token"
//	@Security		BearerAuth
//	@Success		200	{object}	appresponse.Response{data=[]model.Analytics}	"Success response"
//	@Router			/r/readAnalytics [post]
func (ar *analyticsRoute) ReadAnalytics(c *gin.Context) {

	user := query.User(c)

	result, err := ar.analyticsService.ReadAnalytics(user)
	if err != nil {
		c.JSON(http.StatusOK, appresponse.NewBadRequestError(err.Error()))
		return
	}

	c.JSON(http.StatusOK, result)
}
