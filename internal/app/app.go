package app

import (
	"context"
	"fmt"
	"net/http"
	"onelab2/internal/config"
	"onelab2/internal/repository"
	routes "onelab2/internal/transport/http"
	"os"
	"os/signal"
	"time"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go.uber.org/zap"
)

const (
	ctxTimeout = 10 * time.Second
)

func prepareEnv() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
}

func Run() {
	prepareEnv()
	cfg, err := config.New()
	if err != nil {
		panic(err)
	}

	var (
		l *zap.Logger
	)

	if cfg.AppMode == "debug" {
		l, _ = zap.NewDevelopment()
	} else {
		l, _ = zap.NewProduction()
	}
	defer func(l *zap.Logger) {
		err = l.Sync()
		if err != nil {
			panic(err)
		}
	}(l)

	l.Info("app started")

	e := NewServer()

	// repostiory

	userRepo := repository.NewUserRepository(l)
	routes.NewUser(userRepo, l).Register(e)

	go func() {
		if err = e.Start(fmt.Sprintf(":%s", cfg.AppPort)); err != nil {
			l.Fatal("shutting down the server", zap.Error(err))
		}
	}()

	// Graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	l.Info("shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), ctxTimeout)
	defer cancel()

	if err = e.Shutdown(ctx); err != nil {
		l.Fatal("could not stop server gracefully", zap.Error(err))
	}
	l.Info("server stopped")
}

func NewServer() *echo.Echo {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.GET("/ping", func(c echo.Context) error {
		return c.String(http.StatusOK, "pong")
	})

	return e
}
