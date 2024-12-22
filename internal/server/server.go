package server

import (
	"log/slog"
	"time"

	"github.com/garrettladley/prods/internal/constants"
	"github.com/garrettladley/prods/internal/handlers"
	"github.com/garrettladley/prods/internal/settings"
	"github.com/garrettladley/prods/internal/storage"
	"github.com/garrettladley/prods/internal/views/x404"
	"github.com/garrettladley/prods/internal/xerr"
	"github.com/garrettladley/prods/internal/xtempl"

	go_json "github.com/goccy/go-json"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/etag"
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
	setupRobotsTxt(app)
	setupSiteMap(app)

	service := handlers.NewService(cfg.Storage)
	app.Get(constants.APIVersion+"/docs/*", etag.New(), swagger.HandlerDefault)
	service.Routes(app)
	cfg.StaticFn(app)

	setup404Handler(app)

	return app
}

func setupMiddleware(app *fiber.App, cfg *Config) {
	app.Use(recover.New())
	app.Use(slogfiber.NewWithConfig(cfg.Logger,
		slogfiber.Config{
			WithUserAgent:      true,
			WithRequestID:      true,
			WithRequestBody:    true,
			WithRequestHeader:  true,
			WithResponseHeader: true,
		},
	))
	app.Use(limiter.New(limiter.Config{
		Max:               256,
		Expiration:        30 * time.Second,
		LimiterMiddleware: limiter.SlidingWindow{},
	}))
	app.Use(compress.New(compress.Config{
		Level: compress.LevelBestSpeed,
	}))
}

func setupHealthCheck(app *fiber.App) {
	app.Get("/health", func(c *fiber.Ctx) error { return c.SendStatus(fiber.StatusOK) })
}

func setupFavicon(app *fiber.App) {
	app.Get("/favicon.ico", func(c *fiber.Ctx) error { return c.SendStatus(fiber.StatusNoContent) })
}

func setup404Handler(app *fiber.App) {
	app.Use(func(c *fiber.Ctx) error { return xtempl.Render(c, x404.Index()) })
}

//go:embed artifacts/robots.txt
var robotsTxt string

func setupRobotsTxt(app *fiber.App) {
	app.Get("/robots.txt", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).SendString(robotsTxt)
	})
}

//go:embed artifacts/sitemap.xml
var sitemapXml string

func setupSiteMap(app *fiber.App) {
	app.Get("/sitemap.xml", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).SendString(sitemapXml)
	})
}
