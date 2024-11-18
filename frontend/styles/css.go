package styles

import (
	_ "embed"

	"github.com/pterm/pterm"
	"github.com/tdewolff/minify/v2/minify"
)

//go:embed style.css
var css string

func RenderCSS() string {
	minified, _ := minify.CSS(css)
	return pterm.Sprintf("<style>%s</style>", minified)
}
