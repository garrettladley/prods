package main

import (
	"context"
	"embed"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/garrettladley/prods/docs"
	"github.com/garrettladley/prods/internal/constants"
	"github.com/garrettladley/prods/internal/server"
	"github.com/garrettladley/prods/internal/settings"
	"github.com/garrettladley/prods/internal/storage/postgres"
	"github.com/garrettladley/prods/internal/xslog"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/filesystem"
)

//	@title			Prods
//	@version		1.0
//	@description	Generate Spring 2025 Coding Challenge

//	@contact.name	Garrett Ladley
//	@contact.email	ladley.g@northeastern.edu

//	@license.name	MIT
//	@license.url	https://opensource.org/licenses/MIT

//	@host		prods.garrettladley.com
//	@BasePath	/

// @schemes https http

// @externalDocs.description	OpenAPI
// @externalDocs.url			https://swagger.io/resources/open-api/
func main() {
	docs.SwaggerInfo.Version = constants.Version

	handler := slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level:     slog.LevelDebug,
		AddSource: true,
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

	store, err := postgres.New(postgres.Config{
		DSN:             settings.DB.DSN,
		MaxOpenConns:    settings.DB.MaxOpenConns,
		MaxIdleConns:    settings.DB.MaxIdleConns,
		ConnMaxLifetime: settings.DB.ConnMaxLifetime,
	})
	if err != nil {
		slog.LogAttrs(
			ctx,
			slog.LevelError,
			"failed to connect to database",
			xslog.Error(err),
		)
		os.Exit(1)
	}

	app := server.New(&server.Config{
		Storage:  &store,
		Settings: &settings,
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

func static(app *fiber.App) {
	app.Use("/public", filesystem.New(filesystem.Config{
		Root:       http.FS(PublicFS),
		PathPrefix: "public",
		Browse:     true,
	}))
}
