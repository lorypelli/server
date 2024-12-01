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

const (
	WS_PORT                uint16 = 50643
	DEFAULT_DIR                   = "."
	DEFAULT_USE_EXT               = true
	DEFAULT_USE_REALTIME          = true
	DEFAULT_EXPOSE_NETWORK        = true
	DEFAULT_EXT                   = ".html"
	DEFAULT_PORT                  = "53273"
)

func main() {
	dir := flag.String("d", "", "Directory to serve")
	ext := flag.String("e", "", "Extension to use")
	name := flag.String("n", "", "App name")
	port := flag.String("p", "", "Port to use")
	flag.Parse()
	*dir = strings.TrimSpace(*dir)
	*ext = strings.TrimSpace(*ext)
	*name = strings.TrimSpace(*name)
	*port = strings.TrimSpace(*port)
	var defaults bool
	var extension bool
	var realtime bool
	var network bool
	if *dir == "" && *ext == "" && *port == "" {
		defaults, _ = pterm.DefaultInteractiveConfirm.WithDefaultValue(true).Show("Do you want to use defaults options?")
		if defaults {
			*dir = DEFAULT_DIR
			extension = DEFAULT_USE_EXT
			realtime = DEFAULT_USE_REALTIME
			network = DEFAULT_EXPOSE_NETWORK
			*ext = DEFAULT_EXT
			*port = DEFAULT_PORT
		}
	}
	if !defaults {
		for *dir == "" {
			*dir, _ = pterm.DefaultInteractiveTextInput.WithDefaultValue(DEFAULT_DIR).Show("Provide directory to serve")
			*dir = strings.TrimSpace(*dir)
		}
		if _, err := os.Stat(*dir); err != nil {
			internal.Exit(err)
		}
		extension, _ = pterm.DefaultInteractiveConfirm.WithDefaultValue(DEFAULT_USE_EXT).Show("Do you want to use the HTML extension?")
		realtime, _ = pterm.DefaultInteractiveConfirm.WithDefaultValue(DEFAULT_USE_REALTIME).Show("Do you want to have realtime loading for HTML files?")
		if realtime {
			pterm.Warning.Printfln("Port %d can't be used since it's in use by the realtime service!", WS_PORT)
		}
		network, _ = pterm.DefaultInteractiveConfirm.WithDefaultValue(DEFAULT_EXPOSE_NETWORK).Show("Do you want to expose also to the local network?")
		for *ext != ".html" && *ext != ".htm" {
			*ext, _ = pterm.DefaultInteractiveSelect.WithOptions([]string{".html", ".htm"}).WithDefaultOption(DEFAULT_EXT).Show("Choose HTML extension")
			*ext = strings.TrimSpace(*ext)
		}
		if *name == "" {
			*name, _ = pterm.DefaultInteractiveTextInput.Show("Provide app name")
			*name = strings.TrimSpace(*name)
		}
		for *port == "" || (realtime && *port == pterm.Sprint(WS_PORT)) {
			*port, _ = pterm.DefaultInteractiveTextInput.WithDefaultValue(DEFAULT_PORT).Show("Provide port to use")
			*port = strings.TrimSpace(*port)
		}
	}
	p, err := strconv.Atoi(*port)
	if err != nil {
		internal.Exit(err)
	}
	go func() {
		if realtime {
			pkg.StartWebsocket(*dir, WS_PORT)
		}
	}()
	pkg.Start(*dir, *ext, *name, extension, network, realtime, uint16(p), WS_PORT)
}
