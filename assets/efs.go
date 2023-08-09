package assets

import (
	"embed"
)

//go:embed "migrations" "templates" "static"
var EmbeddedFiles embed.FS
