package templ

import (
	"github.com/a-h/templ"
	"github.com/garrettladley/prods/internal/xhttp"
	"github.com/gofiber/fiber/v2"
)

func Render(c *fiber.Ctx, component templ.Component) error {
	c.Set(xhttp.HeaderContentType, xhttp.HeaderTextHTML+"; "+xhttp.HeaderCharsetUTF8)
	return component.Render(c.Context(), c.Response().BodyWriter())
}
