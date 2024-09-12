package terminal

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

type Controller struct {
}

func NewController() *Controller {
	return &Controller{}
}

func (c *Controller) handleTerminal(ctx *gin.Context) {
	upgrader := websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
		Subprotocols: []string{"tos"},
	}

	webConn, err := upgrader.Upgrade(ctx.Writer, ctx.Request, nil)
	if err != nil {
		fmt.Printf("WebSocket upgrade error: %v\n", err)
		return
	}

	_ = webConn.WriteMessage(websocket.TextMessage, []byte(fmt.Sprintf("windows not support")))
}

func (c *Controller) RegisterRoute(group *gin.RouterGroup) {
	group.GET("/terminal", c.handleTerminal)
}
