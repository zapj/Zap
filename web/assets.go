package web

import (
	"embed"
)

//go:embed static/assets/*
var ASSETS_FS embed.FS

//go:embed static/index.html
var INDEX_BYTES []byte

//go:embed static/favicon.ico
var FAVICON_BYTES []byte
