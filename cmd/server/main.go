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
	dir := flag.String("d", internal.DEFAULT_DIR, "Directory to serve")
	ext := flag.String("e", internal.DEFAULT_EXT, "Extension to use")
	name := flag.String("n", "", "App name")
	port := flag.String("p", internal.DEFAULT_PORT, "Port to use")
	skip := flag.Bool("y", false, "Skip questions")
	flag.Parse()
	*dir = strings.TrimSpace(*dir)
	*ext = strings.TrimSpace(*ext)
	*name = strings.TrimSpace(*name)
	*port = strings.TrimSpace(*port)
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
		*dir = ""
		*ext = ""
		*port = ""
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
	pkg.Start(*dir, *ext, *name, extension, network, realtime, uint16(p), internal.WS_PORT)
}
