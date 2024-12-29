//go:build !dev

package main

import (
	"embed"
	"net/http"
)

//go:embed public
var publicFS embed.FS // nolint:unused

func public() http.Handler {
	return http.FileServerFS(publicFS)
}
