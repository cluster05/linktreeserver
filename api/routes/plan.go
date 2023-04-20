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

// CreatePlan		 godoc
//
//	@Summary		Create new plan in linktree
//	@Description	Create new plan in linktree
//	@Tags			Plan
//	@Accept			application/json
//	@Produce		application/json
//	@Param			Authorization	header	string	false	"Bearer token"
//	@Security		BearerAuth
//	@Param			request	body		model.CreatePlanDTO						true	"CreatePlan request"
//	@Success		200		{object}	appresponse.Response{data=model.Plan}	"Success response"
//	@Router			/r/createPlan [post]
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

// ReadPlan		 godoc
//
//	@Summary		Read plan in linktree
//	@Description	Read plan in linktree
//	@Tags			Plan
//	@Accept			application/json
//	@Produce		application/json
//	@Param			Authorization	header	string	false	"Bearer token"
//	@Security		BearerAuth
//	@Success		200	{object}	appresponse.Response{data=[]model.Plan}	"Success response"
//	@Router			/r/readPlan [post]
func (pr *planRoute) ReadPlan(c *gin.Context) {
	user := query.User(c)

	result, err := pr.planService.ReadPlan(user)
	if err != nil {
		c.JSON(http.StatusOK, appresponse.NewBadRequestError(err.Error()))
		return
	}
	c.JSON(http.StatusOK, appresponse.NewSuccess(result))
}
