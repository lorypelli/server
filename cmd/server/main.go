package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/lorypelli/server/internal"
	"github.com/lorypelli/server/pkg"
	"github.com/pterm/pterm"
)

type flags struct {
	dir      string
	ext      string
	name     string
	port     uint16
	username string
	password string
	yes      bool
	given    map[string]bool
}

func main() {
	f := parse()
	var o pkg.Options
	if f.yes || confirm("Do you want to use defaults options?", true) {
		o = defaults(f)
	} else {
		o = interactive(f)
	}
	pkg.Start(o)
}

func parse() flags {
	var f flags
	stringFlag(&f.dir, "dir", "d", internal.DefaultDir, "Directory to serve")
	stringFlag(&f.ext, "ext", "e", internal.DefaultExt, "Extension to use")
	stringFlag(&f.name, "name", "n", "", "App name")
	portFlag(&f.port, "port", "p", internal.DefaultPort, "Port to use")
	stringFlag(&f.username, "username", "user", "", "Username for authentication")
	stringFlag(&f.password, "password", "pwd", "", "Password for authentication")
	flag.BoolVar(&f.yes, "yes", false, "Skip questions")
	flag.BoolVar(&f.yes, "y", false, "Alias for --yes (-y)")
	flag.CommandLine.Usage = pkg.Help
	flag.Parse()
	f.given = make(map[string]bool)
	flag.Visit(func(fl *flag.Flag) {
		f.given[fl.Name] = true
	})
	for _, s := range []*string{&f.dir, &f.ext, &f.name, &f.username, &f.password} {
		*s = strings.TrimSpace(*s)
	}
	return f
}

func stringFlag(p *string, name, alias, value, usage string) {
	flag.StringVar(p, name, value, usage)
	flag.StringVar(p, alias, value, fmt.Sprintf("Alias for --%s (-%s)", name, alias))
}

func portFlag(p *uint16, name, alias string, value uint16, usage string) {
	*p = value
	v := (*portValue)(p)
	flag.Var(v, name, usage)
	flag.Var(v, alias, fmt.Sprintf("Alias for --%s (-%s)", name, alias))
}

type portValue uint16

func (p *portValue) String() string {
	return strconv.Itoa(int(*p))
}

func (p *portValue) Set(s string) error {
	port, err := parsePort(s)
	*p = portValue(port)
	return err
}

func (f flags) provided(name, alias string) bool {
	return f.given[name] || f.given[alias]
}

func defaults(f flags) pkg.Options {
	checkDir(f.dir)
	return pkg.Options{
		Dir:       f.dir,
		Ext:       f.ext,
		Name:      f.name,
		Username:  f.username,
		Password:  f.password,
		Port:      f.port,
		Extension: internal.DefaultUseExt,
		Realtime:  internal.DefaultUseRealtime,
		Network:   internal.DefaultExposeNetwork,
	}
}

func interactive(f flags) pkg.Options {
	dir := f.dir
	if !f.provided("dir", "d") || dir == "" {
		dir = askRequired("Provide directory to serve", internal.DefaultDir)
	}
	checkDir(dir)
	extension := confirm("Do you want to use the HTML extension?", internal.DefaultUseExt)
	realtime := confirm("Do you want to have realtime loading for HTML files?", internal.DefaultUseRealtime)
	if realtime {
		pterm.Warning.Printfln("Port %d can't be used since it's in use by the realtime service!", internal.WSPort)
	}
	network := confirm("Do you want to expose also to the local network?", internal.DefaultExposeNetwork)
	ext := f.ext
	if !f.provided("ext", "e") || ext == "" {
		ext = askExtension()
	}
	name := f.name
	if name == "" {
		name = ask("Provide app name", "")
	}
	port := f.port
	if !f.provided("port", "p") {
		port = askPort("Provide port to use", internal.DefaultPort)
	}
	username := f.username
	if username == "" {
		username = ask("Provide username for authentication", "")
	}
	password := f.password
	if password == "" {
		password = ask("Provide password for authentication", "")
	}
	return pkg.Options{
		Dir:       dir,
		Ext:       ext,
		Name:      name,
		Username:  username,
		Password:  password,
		Port:      port,
		Extension: extension,
		Realtime:  realtime,
		Network:   network,
	}
}

func checkDir(dir string) {
	if _, err := os.Stat(dir); err != nil {
		internal.Exit(err)
	}
}

func parsePort(s string) (uint16, error) {
	port, err := strconv.ParseUint(strings.TrimSpace(s), 10, 16)
	return uint16(port), err
}

func confirm(question string, value bool) bool {
	answer, _ := pterm.DefaultInteractiveConfirm.WithDefaultValue(value).Show(question)
	return answer
}

func ask(question, value string) string {
	answer, _ := pterm.DefaultInteractiveTextInput.WithDefaultValue(value).Show(question)
	return strings.TrimSpace(answer)
}

func askRequired(question, value string) string {
	for {
		if answer := ask(question, value); answer != "" {
			return answer
		}
	}
}

func askPort(question string, value uint16) uint16 {
	for {
		if port, err := parsePort(ask(question, strconv.Itoa(int(value)))); err == nil {
			return port
		}
	}
}

const customExt = "..."

func askExtension() string {
	for {
		choice, _ := pterm.DefaultInteractiveSelect.WithOptions([]string{".html", ".htm", customExt}).WithDefaultOption(internal.DefaultExt).Show("Choose HTML extension")
		if choice == customExt {
			choice = ask("Provide extension to use", "")
		}
		if choice != "" {
			return choice
		}
	}
}
