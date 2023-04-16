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

type PlanRoute interface {
	CreatePlan(*gin.Context)
	ReadPlan(*gin.Context)
}

type planRoute struct {
	planService service.PlanService
}

type PlanRouteConfig struct {
	PlanService service.PlanService
}

func NewPlanRoute(config *PlanRouteConfig) PlanRoute {
	return &planRoute{
		planService: config.PlanService,
	}
}

func (pr *planRoute) CreatePlan(c *gin.Context) {

	var createPlanDTO model.CreatePlanDTO
	if isValid := requesthandler.BindData(c, &createPlanDTO); !isValid {
		return
	}

	user := query.User(c)

	result, err := pr.planService.CreatePlan(user, createPlanDTO)
	if err != nil {
		c.JSON(http.StatusOK, appresponse.NewBadRequestError(err.Error()))
		return
	}
	c.JSON(http.StatusOK, appresponse.NewSuccess(result))
}

func (pr *planRoute) ReadPlan(c *gin.Context) {
	user := query.User(c)

	result, err := pr.planService.ReadPlan(user)
	if err != nil {
		c.JSON(http.StatusOK, appresponse.NewBadRequestError(err.Error()))
		return
	}
	c.JSON(http.StatusOK, appresponse.NewSuccess(result))
}
