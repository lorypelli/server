package main

import (
	"flag"
	"os"
	"strconv"
	"strings"

	"github.com/pterm/pterm"
)

const WS_PORT = 50643

func main() {
	dir := flag.String("d", "", "Directory to serve")
	ext := flag.String("e", "", "Extension to use")
	name := flag.String("n", "", "App name")
	port := flag.String("p", "", "Port to use")
	flag.Parse()
	if *dir == "" {
		*dir, _ = pterm.DefaultInteractiveTextInput.WithDefaultText("Provide directory to serve").WithDefaultValue(".").Show()
	}
	if _, err := os.Stat(*dir); err != nil {
		pterm.Error.Println(err)
		os.Exit(1)
	}
	extension, _ := pterm.DefaultInteractiveConfirm.WithDefaultText("Do you want to use the HTML extension?").WithDefaultValue(true).Show()
	realtime, _ := pterm.DefaultInteractiveConfirm.WithDefaultText("Do you want to have realtime loading for HTML files?").Show()
	if *ext != ".html" && *ext != ".htm" {
		*ext, _ = pterm.DefaultInteractiveSelect.WithDefaultText("Choose HTML extension").WithOptions([]string{".html", ".htm"}).Show()
	}
	if realtime {
		pterm.Warning.Printfln("Port %d can't be used since it's in use by the realtime service!", WS_PORT)
	}
	if *name == "" {
		*name, _ = pterm.DefaultInteractiveTextInput.WithDefaultText("Provide app name").Show()
	}
	if *port == "" {
		*port, _ = pterm.DefaultInteractiveTextInput.WithDefaultText("Provide port to use").WithDefaultValue("53273").Show()
	}
	*dir = strings.TrimSpace(*dir)
	*ext = strings.TrimSpace(*ext)
	*name = strings.TrimSpace(*name)
	*port = strings.TrimSpace(*port)
	p, err := strconv.Atoi(*port)
	if err != nil {
		pterm.Error.Println(err)
		os.Exit(1)
	}
	go func() {
		if realtime {
			StartWebsocket(*dir, WS_PORT)
		}
	}()
	Start(*dir, *ext, *name, extension, realtime, p, WS_PORT)
}
