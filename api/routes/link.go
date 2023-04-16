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

type LinkRoute interface {
	CreateLink(*gin.Context)
	ReadLink(*gin.Context)
	UpdateLink(*gin.Context)
	DeleteLink(*gin.Context)
}

type linkRoute struct {
	linkService service.LinkService
}

type LinkRouteConfig struct {
	LinkService service.LinkService
}

func NewLinkRoute(config *LinkRouteConfig) LinkRoute {
	return &linkRoute{
		linkService: config.LinkService,
	}
}

func (lr *linkRoute) CreateLink(c *gin.Context) {

	var createLinkDTO model.CreateLinkDTO
	if valid := requesthandler.BindData(c, &createLinkDTO); !valid {
		return
	}

	user := query.User(c)

	link, err := lr.linkService.CreateLink(user, createLinkDTO)
	if err != nil {
		c.JSON(http.StatusOK, appresponse.NewBadRequestError(err.Error()))
		return
	}
	c.JSON(http.StatusOK, appresponse.NewSuccess(link))
}

func (lr *linkRoute) ReadLink(c *gin.Context) {

	user := query.User(c)

	links, err := lr.linkService.ReadLink(user)
	if err != nil {
		c.JSON(http.StatusOK, appresponse.NewBadRequestError(err.Error()))
		return
	}
	c.JSON(http.StatusOK, appresponse.NewSuccess(links))
}

func (lr *linkRoute) UpdateLink(c *gin.Context) {
	var updateLinkDTO model.UpdateLinkDTO
	if valid := requesthandler.BindData(c, &updateLinkDTO); !valid {
		return
	}

	user := query.User(c)

	link, err := lr.linkService.UpdateLink(user, updateLinkDTO)
	if err != nil {
		c.JSON(http.StatusOK, appresponse.NewBadRequestError(err.Error()))
		return
	}
	c.JSON(http.StatusOK, appresponse.NewSuccess(link))
}

func (lr *linkRoute) DeleteLink(c *gin.Context) {
	var deleteLinkDTO model.DeleteLinkDTO
	if valid := requesthandler.BindData(c, &deleteLinkDTO); !valid {
		return
	}

	user := query.User(c)

	err := lr.linkService.DeleteLink(user, deleteLinkDTO)
	if err != nil {
		c.JSON(http.StatusOK, appresponse.NewBadRequestError(err.Error()))
		return
	}
	c.JSON(http.StatusOK, appresponse.NewSuccess("deleted successfully"))
}
