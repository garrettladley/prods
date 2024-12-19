package handlers

import (
	"time"

	"github.com/garrettladley/prods/internal/constants"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cache"
	"github.com/gofiber/fiber/v2/middleware/etag"
	"github.com/gofiber/fiber/v2/utils"
)

func (s *Service) Routes(r fiber.Router) {
	cache := cache.New(cache.Config{
		KeyGenerator: func(c *fiber.Ctx) string { return utils.CopyString(c.OriginalURL()) },
		Expiration:   time.Hour * 24 * 365, // 1 year
		CacheControl: true,
	})

	r.Route("/", func(router fiber.Router) {
		r.Use(etag.New())
		r.Use(cache)
		r.Get("/", s.Home)
		r.Get("/frontend", s.Frontend)
		r.Get("/backend", s.Backend)
	})

	r.Route(constants.APIVersion, func(r fiber.Router) {
		r.Post("/register", s.Register)
		r.Get("/token", s.Token)

		r.Route("/:token", func(r fiber.Router) {
			r.Get("/prompt", s.Prompt)
			r.Post("/submit", s.Submit)
		})

		r.Route("/products", func(r fiber.Router) {
			r.Get("/", s.Product)
			r.Get("/:id", s.Product)
		})
	},
	)
}
