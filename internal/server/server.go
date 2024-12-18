package server

import (
	"fmt"
	"log/slog"
	"net/http"
	"strings"
	"time"

	"github.com/garrettladley/prods/internal/constants"
	"github.com/garrettladley/prods/internal/handlers"
	"github.com/garrettladley/prods/internal/settings"
	"github.com/garrettladley/prods/internal/storage"
	"github.com/garrettladley/prods/internal/xerr"
	go_json "github.com/goccy/go-json"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/limiter"

	_ "embed"

	"github.com/gofiber/swagger"

	_ "github.com/garrettladley/prods/docs"
	"github.com/gofiber/fiber/v2/middleware/recover"
	slogfiber "github.com/samber/slog-fiber"
)

type Config struct {
	Settings *settings.Settings
	Storage  storage.Storage
	Logger   *slog.Logger
	StaticFn func(*fiber.App)
}

// TODO: cache
func New(cfg *Config) *fiber.App {
	app := fiber.New(fiber.Config{
		JSONEncoder:       go_json.Marshal,
		JSONDecoder:       go_json.Unmarshal,
		ErrorHandler:      xerr.ErrorHandler,
		PassLocalsToViews: true,
	})
	setupMiddleware(app, cfg)
	setupHealthCheck(app)
	setupFavicon(app)

	service := handlers.NewService(cfg.Storage)
	app.Get(fmt.Sprintf("/api/v%d/docs/*", constants.Major), swagger.HandlerDefault)
	service.Routes(app)
	cfg.StaticFn(app)

	setup404Handler(app)

	return app
}

func setupMiddleware(app *fiber.App, cfg *Config) {
	app.Use(recover.New())
	app.Use(slogfiber.New(cfg.Logger))
	app.Use(limiter.New(limiter.Config{
		Max:               256,
		Expiration:        30 * time.Second,
		LimiterMiddleware: limiter.SlidingWindow{},
	}))
	app.Use(compress.New(compress.Config{
		Level: compress.LevelBestSpeed,
	}))
	app.Use(cors.New(cors.Config{
		AllowOrigins:     strings.Join([]string{"https://prods.garrettladley.com", "http://prods.garrettladley.com", "http://127.0.0.1"}, ","),
		AllowMethods:     strings.Join([]string{http.MethodGet, http.MethodPost, http.MethodPut, http.MethodDelete, http.MethodOptions}, ","),
		AllowHeaders:     strings.Join([]string{"Accept", "Authorization", "Content-Type"}, ","),
		AllowCredentials: true,
		MaxAge:           300,
	}))
}

func setupHealthCheck(app *fiber.App) {
	app.Get("/health", func(c *fiber.Ctx) error {
		return c.SendStatus(fiber.StatusOK)
	})
}

func setupFavicon(app *fiber.App) {
	app.Get("/favicon.ico", func(c *fiber.Ctx) error {
		return c.SendStatus(http.StatusNoContent)
	})
}

// TODO: create templ view
func setup404Handler(app *fiber.App) {
	app.Use(func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Page not found",
			"path":  c.Path(),
		})
	})
}
