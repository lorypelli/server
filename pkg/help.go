package pkg

import (
	"flag"
	"fmt"
	"strings"

	"github.com/pterm/pterm"
)

func Help() {
	var b strings.Builder
	flag.VisitAll(func(f *flag.Flag) {
		fmt.Fprintf(&b, "%s - %s", f.Name, f.Usage)
		if f.DefValue != "" {
			fmt.Fprintf(&b, " (default: %q)", f.DefValue)
		}
		b.WriteByte('\n')
	})
	pterm.DefaultBox.WithTitle("Help Menu").WithTitleTopCenter().Println(strings.TrimSuffix(b.String(), "\n"))
}
