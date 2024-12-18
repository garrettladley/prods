package server

import (
	"fmt"
	"log/slog"
	"time"

	"github.com/garrettladley/prods/internal/constants"
	"github.com/garrettladley/prods/internal/server/handlers"
	"github.com/garrettladley/prods/internal/xerr"
	go_json "github.com/goccy/go-json"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/limiter"

	"github.com/gofiber/fiber/v2/middleware/recover"
	slogfiber "github.com/samber/slog-fiber"

	_ "embed"
)

type Config struct {
	Logger *slog.Logger
}

func New(cfg *Config) *fiber.App {
	app := fiber.New(fiber.Config{
		JSONEncoder:       go_json.Marshal,
		JSONDecoder:       go_json.Unmarshal,
		ErrorHandler:      xerr.ErrorHandler,
		PassLocalsToViews: true,
	})
	setupMiddleware(app, cfg)
	setupHealthCheck(app)

	// register routes here before 404 handler
	service := handlers.NewService(nil)
	app.Route(fmt.Sprintf("/api/v%d", constants.Version), func(r fiber.Router) {
		r.Post("/register", service.Register)
		r.Get("/token", service.Token)
		r.Route("/:token", func(r fiber.Router) {
			r.Get("/prompt", service.Prompt)
			r.Post("/submit", service.Submit)
		})
		r.Get("/products", service.Products)
	})

	setup404Handler(app)

	return app
}

func setupMiddleware(app *fiber.App, cfg *Config) {
	app.Use(recover.New())
	app.Use(slogfiber.New(cfg.Logger))
	app.Use(limiter.New(limiter.Config{
		Max:               20,
		Expiration:        30 * time.Second,
		LimiterMiddleware: limiter.SlidingWindow{},
	}))
	app.Use(compress.New(compress.Config{
		Level: compress.LevelBestSpeed,
	}))
}

func setupHealthCheck(app *fiber.App) {
	app.Get("/health", func(c *fiber.Ctx) error {
		return c.SendStatus(fiber.StatusOK)
	})
}

func setup404Handler(app *fiber.App) {
	app.Use(func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Page not found",
			"path":  c.Path(),
		})
	})
}
