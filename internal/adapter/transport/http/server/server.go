package server

import (
	"context"
	"errors"
	"github.com/goccy/go-json"

	"fmt"
	"gateway/internal/adapter/config"
	"gateway/internal/core/port/http"
	"github.com/gofiber/fiber/v3"
	"go.uber.org/zap"
	std_http "net/http"
	"time"
)

const (
	viewPath   = "../../internal/adapter/transport/http/web/views"
	renderType = ".html"
)

var (
	_ http.ServerMaker = (*server)(nil)
)

type (
	server struct {
		ctx              context.Context
		cfg              *config.Config
		gatewayLogger    *zap.Logger
		authLogger       *zap.Logger
		managementLogger *zap.Logger
		app              *fiber.App
	}
)

func NewHTTPServer(
	ctx context.Context,
	cfg *config.Config,
	gatewayLogger *zap.Logger,
	authLogger *zap.Logger,
	managementLogger *zap.Logger,
) http.ServerMaker {
	return &server{
		ctx:              ctx,
		cfg:              cfg,
		gatewayLogger:    gatewayLogger,
		authLogger:       authLogger,
		managementLogger: managementLogger,
	}
}

func (s *server) Start(ctx context.Context) error {
	app := fiber.New(fiber.Config{
		ReadTimeout:   time.Minute * time.Duration(s.cfg.Settings.ServerReadTimeout),
		StrictRouting: false,
		CaseSensitive: true,
		BodyLimit:     4 * 1024 * 1024,
		JSONEncoder:   json.Marshal,
		JSONDecoder:   json.Unmarshal,
		AppName:       "Go-Gateway",
		Immutable:     true,
	})

	s.app = app
	fiberConnURL := fmt.Sprintf("%s:%d", s.cfg.HTTP.Host, s.cfg.HTTP.Port)

	go func() {
		zap.S().Info("Starting HTTP server on ", fiberConnURL)
		if err := s.app.Listen(fiberConnURL); err != nil {
			if errors.Is(err, std_http.ErrServerClosed) {
				return
			}
			zap.S().Fatal("server listen error: %w", err)
		}
	}()

	err := s.HTTPMiddleware()
	if err != nil {
		zap.S().Fatal("middleware error:", err)
	}
	s.SetupRouter()

	return nil
}

func (s *server) Close(ctx context.Context) error {
	zap.S().Info("HTTP-Server Context is done. Shutting down server...")
	if err := s.app.ShutdownWithContext(ctx); err != nil {
		zap.S().Info("server shutdown error: %w", err)
		return err
	}
	return nil
}
