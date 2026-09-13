package pkg

import (
	"flag"
	"fmt"
	"strings"

	"github.com/pterm/pterm"
)

func Help() {
	var lines []string
	flag.VisitAll(func(f *flag.Flag) {
		line := fmt.Sprintf("%s - %s", f.Name, f.Usage)
		if f.DefValue != "" {
			line = fmt.Sprintf("%s (default: %q)", line, f.DefValue)
		}
		lines = append(lines, line)
	})
	pterm.DefaultBox.WithTitle("Help Menu").WithTitleTopCenter().Println(strings.Join(lines, "\n"))
}
