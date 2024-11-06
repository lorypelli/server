package styles

import (
	"embed"

	"github.com/pterm/pterm"
	"github.com/tdewolff/minify/v2/minify"
)

//go:embed style.css
var css embed.FS

func RenderCSS() string {
	body, _ := css.ReadFile("style.css")
	css, _ := minify.CSS(string(body))
	return pterm.Sprintf("<style>%s</style>", css)
}
