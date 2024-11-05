package utils

import (
	"os"

	"github.com/pterm/pterm"
	"github.com/tdewolff/minify/v2/minify"
)

func RenderCSS() string {
	body, _ := os.ReadFile("frontend/styles/style.css")
	css, _ := minify.CSS(string(body))
	return pterm.Sprintf("<style>%s</style>", css)
}
