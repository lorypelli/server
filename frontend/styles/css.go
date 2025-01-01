package styles

import (
	"embed"

	"github.com/pterm/pterm"
	"github.com/tdewolff/minify/v2/minify"
)

//go:embed *.css
var css embed.FS

func RenderCSS(f string) string {
	var minified string
	switch f {
	case "main":
		{
			file, _ := css.ReadFile("style.css")
			minified, _ = minify.CSS(string(file))
			break
		}
	case "error":
		{
			file, _ := css.ReadFile("error.css")
			minified, _ = minify.CSS(string(file))
			break
		}
	}
	return pterm.Sprintf("<style>%s</style>", minified)
}
