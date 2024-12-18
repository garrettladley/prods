package handlers

import (
	"fmt"

	"github.com/garrettladley/prods/internal/constants"
	"github.com/gofiber/fiber/v2"
)

func (s *Service) Routes(r fiber.Router) {
	r.Get("/", s.Home)
	r.Route("/challenges", func(r fiber.Router) {
		r.Get("/frontend", s.Frontend)
	})

	r.Route(fmt.Sprintf("/api/v%d", constants.Version), func(router fiber.Router) {
		r.Post("/register", s.Register)
		r.Get("/token", s.Token)

		r.Route("/:token", func(r fiber.Router) {
			r.Get("/prompt", s.Prompt)
			r.Post("/submit", s.Submit)
		})

		r.Get("/products", s.Products)
	},
	)
}
