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

// CreateLink		 godoc
//
//	@Summary		Create new link in linktree
//	@Description	Create new link in linktree
//	@Tags			Link
//	@Accept			application/json
//	@Produce		application/json
//	@Param			Authorization	header	string	false	"Bearer token"
//	@Security		BearerAuth
//	@Param			request	body		model.CreateLinkDTO						true	"CreateLink request"
//	@Success		200		{object}	appresponse.Response{data=[]model.Link}	"Success response"
//	@Router			/r/createLink [post]
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

// ReadLink		 godoc
//
//	@Summary		Read new link in linktree
//	@Description	Read new link in linktree
//	@Tags			Link
//	@Param			Authorization	header	string	false	"Bearer token"
//	@Security		BearerAuth
//	@Produce		application/json
//	@Success		200	{object}	appresponse.Response{data=[]model.Link}	"Success response"
//	@Router			/r/readLink [post]
func (lr *linkRoute) ReadLink(c *gin.Context) {

	user := query.User(c)

	links, err := lr.linkService.ReadLink(user)
	if err != nil {
		c.JSON(http.StatusOK, appresponse.NewBadRequestError(err.Error()))
		return
	}
	c.JSON(http.StatusOK, appresponse.NewSuccess(links))
}

// UpdateLink		 godoc
//
//	@Summary		Update new link in linktree
//	@Description	Update new link in linktree
//	@Tags			Link
//	@Accept			application/json
//	@Produce		application/json
//	@Param			Authorization	header	string	false	"Bearer token"
//	@Security		BearerAuth
//	@Param			request	body		model.UpdateLinkDTO						true	"UpdateLink request"
//	@Success		200		{object}	appresponse.Response{data=model.Link}	"Success response"
//	@Router			/r/updateLink [post]
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

// DeleteLink		 godoc
//
//	@Summary		Delete new link in linktree
//	@Description	Delete new link in linktree
//	@Tags			Link
//	@Accept			application/json
//	@Produce		application/json
//	@Param			Authorization	header	string	false	"Bearer token"
//	@Security		BearerAuth
//	@Param			request	body		model.DeleteLinkDTO					true	"DeleteLink request"
//	@Success		200		{object}	appresponse.Response{data=string}	"Success response"
//	@Router			/r/deleteLink [post]
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
