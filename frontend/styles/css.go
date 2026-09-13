package styles

import (
	_ "embed"
	"fmt"
)

//go:embed style.css
var source string

var CSS = fmt.Sprintf("<style>%s</style>", source)
