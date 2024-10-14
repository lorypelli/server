package main

import (
	"flag"
	"os"
	"strconv"
	"strings"

	"github.com/pterm/pterm"
)

func main() {
	dir := flag.String("d", "", "Directory to serve")
	name := flag.String("n", "", "App name")
	port := flag.String("p", "", "Port to use")
	flag.Parse()
	if strings.TrimSpace(*dir) == "" {
		*dir, _ = pterm.DefaultInteractiveTextInput.WithDefaultText("Provide directory to serve").WithDefaultValue(".").Show()
	}
	*dir = strings.TrimSpace(*dir)
	if _, err := os.Stat(*dir); err != nil {
		pterm.Error.Println(err)
		os.Exit(1)
	}
	extension, _ := pterm.DefaultInteractiveConfirm.WithDefaultText("Do you want to use the .html extension?").WithDefaultValue(true).Show()
	if strings.TrimSpace(*name) == "" {
		*name, _ = pterm.DefaultInteractiveTextInput.WithDefaultText("Provide app name").Show()
	}
	*name = strings.TrimSpace(*name)
	if strings.TrimSpace(*port) == "" {
		*port, _ = pterm.DefaultInteractiveTextInput.WithDefaultText("Provide port to use").WithDefaultValue("80").Show()
	}
	*port = strings.TrimSpace(*port)
	p, err := strconv.Atoi(*port)
	if err != nil {
		pterm.Error.Println(err)
		os.Exit(1)
	}
	if p < 0 || p > 65535 {
		pterm.Error.Println("Port not in range!")
		os.Exit(1)
	}
	Start(*dir, *name, extension, uint16(p))
}
