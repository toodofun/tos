//go:build !windows

package terminal

import (
	"fmt"
	"github.com/MR5356/tos/util/terminal"
	"github.com/creack/pty"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"golang.org/x/crypto/ssh"
	"io"
	"net/http"
	"os"
	"os/exec"
)

type Controller struct {
}

func NewController() *Controller {
	return &Controller{}
}

type LocalSession struct {
	pty    *os.File // 伪终端文件描述符
	stdin  io.WriteCloser
	stdout io.Reader
	stderr io.Reader
}

func (l *LocalSession) StdinPipe() (io.WriteCloser, error) {
	return l.stdin, nil
}

func (l *LocalSession) StdoutPipe() (io.Reader, error) {
	return l.stdout, nil
}

func (l *LocalSession) StderrPipe() (io.Reader, error) {
	return l.stderr, nil
}

func (l *LocalSession) Start() error {
	return nil
}

func (l *LocalSession) WindowChange(h, w int) error {
	return nil
}

func (l *LocalSession) Wait() error {
	return nil
}

func (l *LocalSession) Close() error {
	if l.pty != nil {
		return l.pty.Close()
	}
	return nil
}

func (l *LocalSession) RequestPty(term string, h, w int, modes ssh.TerminalModes) error {
	return nil
}

func NewLocalSession() *LocalSession {
	// 创建伪终端
	master, slave, err := pty.Open()
	if err != nil {
		fmt.Printf("Failed to open pty: %v\n", err)
		return nil
	}

	// 执行 shell 程序
	shell := os.Getenv("SHELL")
	if shell == "" {
		shell = "/bin/bash"
	}

	homeDir := os.Getenv("HOME")
	if homeDir == "" {
		homeDir = "/"
	}

	cmd := exec.Command(shell)
	cmd.Dir = homeDir

	// 将子进程的 stdin/stdout/stderr 连接到伪终端
	cmd.Stdin = slave
	cmd.Stdout = slave
	cmd.Stderr = slave

	if err := cmd.Start(); err != nil {
		fmt.Printf("Failed to start command: %v\n", err)
		return nil
	}

	return &LocalSession{
		pty:    master,
		stdin:  master,
		stdout: master,
		stderr: master,
	}
}

func (c *Controller) handleTerminal(ctx *gin.Context) {
	t := terminal.NewTerminal()
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

	defer func() {
		if err := recover(); err != nil {
			_ = webConn.WriteMessage(websocket.TextMessage, []byte(fmt.Sprintf("create terminal error: %v", err)))
		}
	}()

	// 创建一个本地 session
	sshOut := new(terminal.WsBufferWriter)
	session := NewLocalSession()

	t.Websocket = webConn
	t.Session = session

	stdin, err := session.StdinPipe()
	if err != nil {
		_ = webConn.WriteMessage(websocket.TextMessage, []byte(fmt.Sprintf("Error creating stdin pipe: %v", err)))
		return
	}

	stdout, err := session.StdoutPipe()
	if err != nil {
		_ = webConn.WriteMessage(websocket.TextMessage, []byte(fmt.Sprintf("Error creating stdout pipe: %v", err)))
		return
	}

	stderr, err := session.StderrPipe()
	if err != nil {
		_ = webConn.WriteMessage(websocket.TextMessage, []byte(fmt.Sprintf("Error creating stderr pipe: %v", err)))
		return
	}

	go io.Copy(sshOut, stdout)
	go io.Copy(sshOut, stderr)

	t.Stdin = stdin
	t.Stdout = sshOut

	err = session.Start()
	if err != nil {
		_ = webConn.WriteMessage(websocket.TextMessage, []byte(fmt.Sprintf("Error starting session: %v", err)))
		return
	}

	webConn.SetCloseHandler(func(code int, text string) error {
		if err := t.CloseHandler(code, text); err != nil {
			return err
		}
		return session.Close()
	})
	go t.Send2SSH()
	go t.Send2Web()
}

func (c *Controller) RegisterRoute(group *gin.RouterGroup) {
	group.GET("/terminal", c.handleTerminal)
}
