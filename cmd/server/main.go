package main

import (
	"flag"
	"os"
	"strconv"
	"strings"

	"github.com/lorypelli/server/internal"
	"github.com/lorypelli/server/pkg"
	"github.com/pterm/pterm"
)

func main() {
	dir := flag.String("dir", internal.DEFAULT_DIR, "Directory to serve")
	flag.StringVar(dir, "d", *dir, "Alias for --dir (-d)")
	ext := flag.String("ext", internal.DEFAULT_EXT, "Extension to use")
	flag.StringVar(ext, "e", *ext, "Alias for --ext (-e)")
	name := flag.String("name", "", "App name")
	flag.StringVar(name, "n", *name, "Alias for --name (-n)")
	port := flag.String("port", internal.DEFAULT_PORT, "Port to use")
	flag.StringVar(port, "p", *port, "Alias for --port (-p)")
	username := flag.String("username", "", "Username for authentication")
	flag.StringVar(username, "user", *username, "Alias for --username (-user)")
	password := flag.String("password", "", "Password for authentication")
	flag.StringVar(password, "pwd", *password, "Alias for --password (-pwd)")
	skip := flag.Bool("yes", false, "Skip questions")
	flag.BoolVar(skip, "y", *skip, "Alias for --yes (-y)")
	flag.CommandLine.Usage = pkg.Help
	flag.Parse()
	*dir = strings.TrimSpace(*dir)
	*ext = strings.TrimSpace(*ext)
	*name = strings.TrimSpace(*name)
	*port = strings.TrimSpace(*port)
	*username = strings.TrimSpace(*username)
	*password = strings.TrimSpace(*password)
	var hasDir bool
	var hasExt bool
	var hasPort bool
	var defaults bool
	var extension bool
	var realtime bool
	var network bool
	if *skip {
		defaults = true
	} else {
		defaults, _ = pterm.DefaultInteractiveConfirm.WithDefaultValue(true).Show("Do you want to use defaults options?")
	}
	if defaults {
		if _, err := os.Stat(*dir); err != nil {
			internal.Exit(err)
		}
		extension = internal.DEFAULT_USE_EXT
		realtime = internal.DEFAULT_USE_REALTIME
		network = internal.DEFAULT_EXPOSE_NETWORK
	} else {
		flag.Visit(func(f *flag.Flag) {
			switch f.Name {
			case "d":
			case "dir":
				{
					hasDir = true
					break
				}
			case "e":
			case "ext":
				{
					hasExt = true
					break
				}
			case "p":
			case "port":
				{
					hasPort = true
					break
				}
			}
		})
		if !hasDir {
			*dir = ""
		}
		if !hasExt {
			*ext = ""
		}
		if !hasPort {
			*port = ""
		}
		for *dir == "" {
			*dir, _ = pterm.DefaultInteractiveTextInput.WithDefaultValue(internal.DEFAULT_DIR).Show("Provide directory to serve")
			*dir = strings.TrimSpace(*dir)
		}
		if _, err := os.Stat(*dir); err != nil {
			internal.Exit(err)
		}
		extension, _ = pterm.DefaultInteractiveConfirm.WithDefaultValue(internal.DEFAULT_USE_EXT).Show("Do you want to use the HTML extension?")
		realtime, _ = pterm.DefaultInteractiveConfirm.WithDefaultValue(internal.DEFAULT_USE_REALTIME).Show("Do you want to have realtime loading for HTML files?")
		if realtime {
			pterm.Warning.Printfln("Port %d can't be used since it's in use by the realtime service!", internal.WS_PORT)
		}
		network, _ = pterm.DefaultInteractiveConfirm.WithDefaultValue(internal.DEFAULT_EXPOSE_NETWORK).Show("Do you want to expose also to the local network?")
		for *ext == "" {
			*ext, _ = pterm.DefaultInteractiveSelect.WithOptions([]string{".html", ".htm", "..."}).WithDefaultOption(internal.DEFAULT_EXT).Show("Choose HTML extension")
			*ext = strings.TrimSpace(*ext)
			if *ext == "..." {
				*ext, _ = pterm.DefaultInteractiveTextInput.Show("Provide extension to use")
				*ext = strings.TrimSpace(*ext)
			}
		}
		if *name == "" {
			*name, _ = pterm.DefaultInteractiveTextInput.Show("Provide app name")
			*name = strings.TrimSpace(*name)
		}
		for *port == "" {
			*port, _ = pterm.DefaultInteractiveTextInput.WithDefaultValue(internal.DEFAULT_PORT).Show("Provide port to use")
			*port = strings.TrimSpace(*port)
		}
		if *username == "" {
			*username, _ = pterm.DefaultInteractiveTextInput.Show("Provide username for authentication")
			*username = strings.TrimSpace(*username)
		}
		if *password == "" {
			*password, _ = pterm.DefaultInteractiveTextInput.Show("Provide password for authentication")
			*password = strings.TrimSpace(*password)
		}
	}
	p, err := strconv.Atoi(*port)
	if err != nil {
		internal.Exit(err)
	}
	go func() {
		if realtime {
			pkg.StartWebsocket(*dir, internal.WS_PORT)
		}
	}()
	pkg.Start(*dir, *ext, *name, *username, *password, extension, network, realtime, uint16(p))
}
