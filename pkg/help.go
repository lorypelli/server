package pkg

import (
	"flag"
	"strings"

	"github.com/pterm/pterm"
)

func Help() {
	box := pterm.DefaultBox.WithTitle("Help Menu").WithTitleTopCenter()
	var msg string
	flag.VisitAll(func(f *flag.Flag) {
		msg += pterm.Sprintf("%s - %s", f.Name, f.Usage)
		initial := f.DefValue
		if initial != "" {
			msg += " "
			msg += pterm.Sprintf("(default: %q)", initial)
		}
		msg += "\n"
	})
	msg = strings.TrimSuffix(msg, "\n")
	box.Println(msg)
}
