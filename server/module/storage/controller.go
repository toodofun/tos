package storage

import (
	"github.com/MR5356/tos/response"
	"github.com/gin-gonic/gin"
	"strings"
)

type Controller struct {
	service *Service
}

func NewController() *Controller {
	return &Controller{
		service: GetService(),
	}
}

func (c *Controller) handleListDirectory(ctx *gin.Context) {
	if path := ctx.Query("path"); path != "" {
		if res, err := c.service.ListDirectory(path); err != nil {
			response.ErrorWithMsg(ctx, response.CodeParamsError, err.Error())
		} else {
			response.Success(ctx, res)
		}
	} else {
		response.Error(ctx, response.CodeParamsError)
	}
}

func (c *Controller) handleExists(ctx *gin.Context) {
	if path := ctx.Query("path"); path != "" {
		response.Success(ctx, c.service.Exists(path))
	} else {
		response.Error(ctx, response.CodeParamsError)
	}
}

func (c *Controller) handleUpload(ctx *gin.Context) {
	fileName := ctx.PostForm("fileName")
	targetPath := ctx.PostForm("targetPath")
	mode := ctx.PostForm("mode")

	if len(mode) == 0 {
		response.ErrorWithMsg(ctx, response.CodeParamsError, "mode is required")
		return
	}

	file, header, err := ctx.Request.FormFile("file")
	if err != nil {
		response.Error(ctx, response.CodeParamsError)
	} else {
		if len(strings.TrimSpace(fileName)) == 0 {
			fileName = header.Filename
		}
		if len(fileName) == 0 {
			response.ErrorWithMsg(ctx, response.CodeParamsError, "file name is required")
			return
		}
		if err := c.service.Upload(fileName, file, targetPath, mode); err != nil {
			response.ErrorWithMsg(ctx, response.CodeParamsError, err.Error())
		} else {
			response.Success(ctx, nil)
		}
	}
}

func (c *Controller) handleGetSpecialPath(ctx *gin.Context) {
	response.Success(ctx, c.service.GetSpecialPath())
}

func (c *Controller) RegisterRoute(group *gin.RouterGroup) {
	api := group.Group("/storage")
	api.POST("/upload", c.handleUpload)
	api.GET("/folder", c.handleListDirectory)
	api.GET("/sp", c.handleGetSpecialPath)
	api.GET("/exists", c.handleExists)
}
