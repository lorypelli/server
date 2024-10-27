package main

import (
	"flag"
	"os"
	"strconv"
	"strings"

	"github.com/pterm/pterm"
)

const WS_PORT uint16 = 50643

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
	for *dir == "" {
		*dir, _ = pterm.DefaultInteractiveTextInput.WithDefaultValue(".").Show("Provide directory to serve")
		*dir = strings.TrimSpace(*dir)
	}
	if _, err := os.Stat(*dir); err != nil {
		Exit(err)
	}
	extension, _ := pterm.DefaultInteractiveConfirm.WithDefaultValue(true).Show("Do you want to use the HTML extension?")
	realtime, _ := pterm.DefaultInteractiveConfirm.Show("Do you want to have realtime loading for HTML files?")
	if realtime {
		pterm.Warning.Printfln("Port %d can't be used since it's in use by the realtime service!", WS_PORT)
	}
	for *ext != ".html" && *ext != ".htm" {
		*ext, _ = pterm.DefaultInteractiveSelect.WithOptions([]string{".html", ".htm"}).Show("Choose HTML extension")
		*ext = strings.TrimSpace(*ext)
	}
	if *name == "" {
		*name, _ = pterm.DefaultInteractiveTextInput.Show("Provide app name")
		*name = strings.TrimSpace(*name)
	}
	for *port == "" || (realtime && *port == pterm.Sprint(WS_PORT)) {
		*port, _ = pterm.DefaultInteractiveTextInput.WithDefaultValue("53273").Show("Provide port to use")
		*port = strings.TrimSpace(*port)
	}
	p, err := strconv.Atoi(*port)
	if err != nil {
		Exit(err)
	}
	go func() {
		if realtime {
			StartWebsocket(*dir, WS_PORT)
		}
	}()
	Start(*dir, *ext, *name, extension, realtime, uint16(p), WS_PORT)
}
