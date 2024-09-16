package application

import (
	"github.com/MR5356/tos/response"
	"github.com/gin-gonic/gin"
)

type Controller struct {
	service *Service
}

func NewController() *Controller {
	return &Controller{
		service: GetService(),
	}
}

func (c *Controller) handleListApplications(ctx *gin.Context) {
	if res, err := c.service.ListApps(); err != nil {
		response.ErrorWithMsg(ctx, response.CodeParamsError, err.Error())
	} else {
		response.Success(ctx, res)
	}
}

func (c *Controller) RegisterRoute(group *gin.RouterGroup) {
	api := group.Group("/application")
	api.GET("/list", c.handleListApplications)
}
