package storage

import (
	"github.com/MR5356/tos/response"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
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
	if id, err := uuid.Parse(ctx.Param("id")); err != nil {
		response.Error(ctx, response.CodeParamsError)
	} else {
		if path := ctx.Query("path"); path != "" {
			if res, err := c.service.ListDirectory(id, path); err != nil {
				response.ErrorWithMsg(ctx, response.CodeParamsError, err.Error())
			} else {
				response.Success(ctx, res)
			}
		} else {
			response.Error(ctx, response.CodeParamsError)
		}
	}
}

func (c *Controller) handleExists(ctx *gin.Context) {
	if id, err := uuid.Parse(ctx.Param("id")); err != nil {
		response.Error(ctx, response.CodeParamsError)
	} else {
		if path := ctx.Query("path"); path != "" {
			response.Success(ctx, c.service.Exists(id, path))
		} else {
			response.Error(ctx, response.CodeParamsError)
		}
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

	if id, err := uuid.Parse(ctx.Param("id")); err != nil {
		response.Error(ctx, response.CodeParamsError)
	} else {
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
			if err := c.service.Upload(id, fileName, file, targetPath, mode); err != nil {
				response.ErrorWithMsg(ctx, response.CodeParamsError, err.Error())
			} else {
				response.Success(ctx, nil)
			}
		}
	}
}

func (c *Controller) handleGetSpecialPath(ctx *gin.Context) {
	if id, err := uuid.Parse(ctx.Param("id")); err != nil {
		response.Error(ctx, response.CodeParamsError)
	} else {
		response.Success(ctx, c.service.GetSpecialPath(id))
	}
}

func (c *Controller) handleListStorages(ctx *gin.Context) {
	if res, err := c.service.ListStorages(); err != nil {
		response.ErrorWithMsg(ctx, response.CodeParamsError, err.Error())
	} else {
		response.Success(ctx, res)
	}
}

func (c *Controller) RegisterRoute(group *gin.RouterGroup) {
	api := group.Group("/storage")
	// 获取存储列表
	api.GET("/list", c.handleListStorages)

	// 对指定存储进行操作
	api.POST("/:id/upload", c.handleUpload)
	api.GET("/:id/folder", c.handleListDirectory)
	api.GET("/:id/sp", c.handleGetSpecialPath)
	api.GET("/:id/exists", c.handleExists)
}
