package system

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

func (c *Controller) handleGetTimestamp(ctx *gin.Context) {
	response.Success(ctx, c.service.GetTimestamp())
}

func (c *Controller) handleGetSystemInfo(ctx *gin.Context) {
	response.Success(ctx, c.service.GetSystemInfo())
}

func (c *Controller) handleGetNetworkInfo(ctx *gin.Context) {
	response.Success(ctx, c.service.GetNetworkInfo())
}

func (c *Controller) handleGetHolidayAPI(ctx *gin.Context) {
	response.Success(ctx, c.service.GetHolidayAPI())
}

func (c *Controller) RegisterRoute(group *gin.RouterGroup) {
	api := group.Group("/system")
	api.GET("/info", c.handleGetSystemInfo)
	api.GET("/timestamp", c.handleGetTimestamp)
	api.GET("/network", c.handleGetNetworkInfo)
	api.GET("/holiday", c.handleGetHolidayAPI)
}
