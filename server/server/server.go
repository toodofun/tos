package server

import (
	"context"
	"embed"
	"errors"
	"fmt"
	"github.com/MR5356/tos/config"
	"github.com/MR5356/tos/module/application"
	"github.com/MR5356/tos/module/storage"
	"github.com/MR5356/tos/module/system"
	"github.com/MR5356/tos/module/terminal"
	"github.com/MR5356/tos/persistence/database"
	"github.com/MR5356/tos/server/ginmiddleware"
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/sirupsen/logrus"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type Server struct {
	engine *gin.Engine
}

type Service interface {
	Initialize() error
}

type Controller interface {
	RegisterRoute(group *gin.RouterGroup)
}

//go:embed static
var fs embed.FS

func New(cfg *config.Config) (server *Server, err error) {
	if cfg.Server.Debug {
		gin.SetMode(gin.DebugMode)
		logrus.SetLevel(logrus.DebugLevel)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	database.NewDatabase(cfg)

	engine := gin.New()
	engine.MaxMultipartMemory = 8 << 20

	engine.Use(
		gin.Recovery(),
		ginmiddleware.Static("/", ginmiddleware.NewStaticFileSystem(fs, "static")),
	)

	api := engine.Group(cfg.Server.Prefix)
	api.Use(
		gzip.Gzip(gzip.DefaultCompression),
	)

	// metrics
	engine.GET(fmt.Sprintf("/%s/metrics", cfg.Server.Prefix), func(handler http.Handler) gin.HandlerFunc {
		return func(context *gin.Context) {
			handler.ServeHTTP(context.Writer, context.Request)
		}
	}(promhttp.Handler()))

	// services
	services := []Service{
		system.GetService(),
		storage.GetService(),
		application.GetService(),
	}

	// controllers
	controllers := []Controller{
		terminal.NewController(),
		system.NewController(),
		storage.NewController(),
		application.NewController(),
	}

	// initialize
	for _, svc := range services {
		if err := svc.Initialize(); err != nil {
			return nil, err
		}
	}
	for _, ctl := range controllers {
		ctl.RegisterRoute(api)
	}

	server = &Server{
		engine: engine,
	}
	return
}

func (s *Server) Run() error {
	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", config.Current().Server.Port),
		Handler: s.engine,
	}

	go func() {
		logrus.Infof("server running on port %d", config.Current().Server.Port)
		if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			logrus.Fatalf("server listening error: %v", err)
		}
	}()

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)

	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(config.Current().Server.GracePeriod)*time.Second)
	defer cancel()

	ch := <-sig
	logrus.Infof("server receive signal: %s", ch.String())
	return server.Shutdown(ctx)
}
