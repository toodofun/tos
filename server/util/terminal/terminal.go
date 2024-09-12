package terminal

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/gorilla/websocket"
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/ssh"
	"io"
	"sync"
	"time"
)

const (
	EventResize = "resize"
)

var Modes = ssh.TerminalModes{
	ssh.ECHO:          1,
	ssh.TTY_OP_ISPEED: 14400,
	ssh.TTY_OP_OSPEED: 14400,
}

type SessionInterface interface {
	StdinPipe() (io.WriteCloser, error)
	StdoutPipe() (io.Reader, error)
	StderrPipe() (io.Reader, error)
	Start() error
	Wait() error
	Close() error
	WindowChange(h, w int) error
	RequestPty(term string, h, w int, modes ssh.TerminalModes) error
}

type Terminal struct {
	Websocket *websocket.Conn
	Stdin     io.WriteCloser
	Stdout    *WsBufferWriter
	Session   SessionInterface
	ctx       context.Context
	cancel    context.CancelFunc
}

func NewTerminal() *Terminal {
	ctx, cancel := context.WithCancel(context.Background())
	return &Terminal{
		ctx:    ctx,
		cancel: cancel,
	}
}

type TerminalEvent struct {
	Operate string `json:"operate"`
	Cols    int    `json:"cols"`
	Rows    int    `json:"rows"`
}

type WsBufferWriter struct {
	buffer bytes.Buffer
	mutex  sync.RWMutex
}

func (w *WsBufferWriter) Write(p []byte) (int, error) {
	w.mutex.Lock()
	defer w.mutex.Unlock()
	return w.buffer.Write(p)
}

func (w *WsBufferWriter) Reset() {
	w.mutex.Lock()
	defer w.mutex.Unlock()
	w.buffer.Reset()
}

func (t *Terminal) Send2SSH() {
	defer func() {
		if err := recover(); err != nil {
			logrus.Errorf("[recover] send websocket data to sshutil error: %v", err)
		}
	}()
	for {
		select {
		case <-t.ctx.Done():
			logrus.Infof("terminal closed ctx")
			return
		default:
			if t.Websocket != nil {
				_, wsData, err := t.Websocket.ReadMessage()

				if err != nil {
					return
				}
				var event TerminalEvent
				err = json.Unmarshal(wsData, &event)
				if err == nil {
					switch event.Operate {
					case EventResize:
						// 重新设置窗口大小
						logrus.Debugf("resize terminal: %+v", event)
						err := t.Session.WindowChange(event.Rows, event.Cols)
						if err != nil {
							logrus.Errorf("resize terminal error: %v", err)
						}
					default:
						logrus.Infof("unknown event: %+v", event)
					}
				} else {
					_, err = t.Stdin.Write(wsData)
					if err != nil {
						logrus.Errorf("send websocket data to sshutil error: %v", err)
					}
				}
			}
		}
	}
}

func (t *Terminal) Send2Web() {
	timer := time.Tick(time.Millisecond * 100)
	defer func() {
		if err := recover(); err != nil {
			logrus.Errorf("[recover] send websocket data error: %v", err)
		}
	}()
	for {
		select {
		case <-t.ctx.Done():
			return
		case <-timer:
			if t.Stdout.buffer.Len() > 0 && t.Websocket != nil {
				err := t.Websocket.WriteMessage(websocket.TextMessage, t.Stdout.buffer.Bytes())
				if err != nil {
					logrus.Errorf("send websocket data error: %v", err)
				}
				t.Stdout.Reset()
			}
		}
	}
}

func (t *Terminal) CloseHandler(code int, msg string) error {
	logrus.Infof("close terminal")
	t.cancel()
	err := t.Session.Close()
	if err != nil {
		logrus.Errorf("close session error: %v", err)
	}
	err = t.Websocket.Close()
	if err != nil {
		logrus.Errorf("close websocket error: %v", err)
	}
	return err
}
