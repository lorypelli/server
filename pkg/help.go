package pkg

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"slices"
	"strings"

	"github.com/lorypelli/server/internal"
	"github.com/pterm/pterm"
)

type option struct {
	names []string
	usage string
	value string
}

func Help() {
	pterm.Println()
	pterm.Printfln("%s %s %s", internal.Strong("Usage:"), executable(), internal.Muted("[flags]"))
	pterm.Println()
	rows := [][]string{{"Flag", "Description", "Default"}}
	for _, o := range options() {
		rows = append(rows, []string{strings.Join(o.names, ", "), o.usage, o.value})
	}
	pterm.DefaultTable.WithHasHeader().WithBoxed().WithData(rows).Render()
	pterm.Println()
}

func options() []option {
	var opts []option
	index := make(map[flag.Value]int)
	flag.VisitAll(func(f *flag.Flag) {
		name := dashed(f.Name)
		if i, ok := index[f.Value]; ok {
			opts[i].names = append(opts[i].names, name)
			return
		}
		index[f.Value] = len(opts)
		opts = append(opts, option{names: []string{name}, usage: f.Usage, value: defaultValue(f)})
	})
	for i := range opts {
		slices.SortFunc(opts[i].names, func(a, b string) int {
			return len(b) - len(a)
		})
	}
	return opts
}

func dashed(name string) string {
	if len(name) == 1 {
		return fmt.Sprintf("-%s", name)
	}
	return fmt.Sprintf("--%s", name)
}

func defaultValue(f *flag.Flag) string {
	if f.DefValue == "" || f.DefValue == "false" {
		return internal.Muted("-")
	}
	return f.DefValue
}

func executable() string {
	return strings.TrimSuffix(filepath.Base(os.Args[0]), filepath.Ext(os.Args[0]))
}
