package main

import (
	"context"
	"embed"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/garrettladley/prods/internal/server"
	"github.com/garrettladley/prods/internal/settings"
	"github.com/garrettladley/prods/internal/xslog"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/filesystem"
)

func main() {
	handler := slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelDebug,
	})
	logger := slog.New(handler)
	slog.SetDefault(logger)

	ctx, cancel := context.WithCancel(context.Background())

	settings, err := settings.Load()
	if err != nil {
		slog.LogAttrs(
			ctx,
			slog.LevelError,
			"failed to load settings",
			xslog.Error(err),
		)
		os.Exit(1)
	}

	app := server.New(&server.Config{
		Logger:   logger,
		StaticFn: static,
	})

	go func() {
		if err := app.Listen(":" + settings.App.Port); err != nil {
			slog.LogAttrs(
				ctx,
				slog.LevelError,
				"failed to start server",
				xslog.Error(err),
			)
			os.Exit(1)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	<-quit
	slog.LogAttrs(
		ctx,
		slog.LevelInfo,
		"stopping server",
	)
	cancel()

	if err := app.Shutdown(); err != nil {
		slog.LogAttrs(
			ctx,
			slog.LevelError,
			"failed to shutdown server",
			xslog.Error(err),
		)
	}

	slog.LogAttrs(
		ctx,
		slog.LevelInfo,
		"server shutdown",
	)
}

//go:embed public
var PublicFS embed.FS

//go:embed deps
var DepsFS embed.FS

func static(app *fiber.App) {
	app.Use("/public", filesystem.New(filesystem.Config{
		Root:       http.FS(PublicFS),
		PathPrefix: "public",
		Browse:     true,
	}))
	app.Use("/deps", filesystem.New(filesystem.Config{
		Root:       http.FS(DepsFS),
		PathPrefix: "deps",
		Browse:     true,
	}))
}
