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
	"sync"
	"time"
)

type Controller struct {
}

func NewController() *Controller {
	return &Controller{}
}

type LocalSession struct {
	pty    *os.File
	stdin  io.WriteCloser
	stdout io.Reader
	stderr io.Reader
	cmd    *exec.Cmd
	wg     sync.WaitGroup
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
	if l.cmd != nil {
		return l.cmd.Wait()
	}
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
	master, slave, err := pty.Open()
	if err != nil {
		fmt.Printf("Failed to open pty: %v\n", err)
		return nil
	}

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
	cmd.Stdin = slave
	cmd.Stdout = slave
	cmd.Stderr = slave

	if err := cmd.Start(); err != nil {
		fmt.Printf("Failed to start command: %v\n", err)
		slave.Close()
		master.Close()
		return nil
	}

	return &LocalSession{
		pty:    master,
		stdin:  master,
		stdout: master,
		stderr: master,
		cmd:    cmd,
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

	session := NewLocalSession()
	if session == nil {
		_ = webConn.WriteMessage(websocket.TextMessage, []byte("Failed to create session"))
		webConn.Close()
		return
	}

	t.Websocket = webConn
	t.Session = session

	stdin, err := session.StdinPipe()
	if err != nil {
		_ = webConn.WriteMessage(websocket.TextMessage, []byte(fmt.Sprintf("Error creating stdin pipe: %v", err)))
		session.Close()
		webConn.Close()
		return
	}

	stdout, err := session.StdoutPipe()
	if err != nil {
		_ = webConn.WriteMessage(websocket.TextMessage, []byte(fmt.Sprintf("Error creating stdout pipe: %v", err)))
		session.Close()
		webConn.Close()
		return
	}

	stderr, err := session.StderrPipe()
	if err != nil {
		_ = webConn.WriteMessage(websocket.TextMessage, []byte(fmt.Sprintf("Error creating stderr pipe: %v", err)))
		session.Close()
		webConn.Close()
		return
	}

	sshOut := new(terminal.WsBufferWriter)
	t.Stdin = stdin
	t.Stdout = sshOut

	// 使用WaitGroup来等待所有goroutine完成
	var wg sync.WaitGroup

	// 启动输出转发
	wg.Add(2)
	go func() {
		defer wg.Done()
		io.Copy(sshOut, stdout)
	}()
	go func() {
		defer wg.Done()
		io.Copy(sshOut, stderr)
	}()

	// 监听进程退出
	go func() {
		// 等待命令执行完成
		if err := session.cmd.Wait(); err != nil {
			fmt.Printf("Command finished with error: %v\n", err)
		}

		// 关闭WebSocket连接
		_ = webConn.WriteControl(websocket.CloseMessage,
			websocket.FormatCloseMessage(websocket.CloseNormalClosure, "Process exited"),
			time.Now().Add(time.Second))
		webConn.Close()

		// 清理资源
		session.Close()
		wg.Wait()
	}()

	err = session.Start()
	if err != nil {
		_ = webConn.WriteMessage(websocket.TextMessage, []byte(fmt.Sprintf("Error starting session: %v", err)))
		session.Close()
		webConn.Close()
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
