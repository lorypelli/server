package main

import (
	"os"
	"strconv"

	"github.com/pterm/pterm"
)

func main() {
	dir, _ := pterm.DefaultInteractiveTextInput.WithDefaultText("Provide directory to serve").Show()
	if _, err := os.Stat(dir); err != nil {
		pterm.Error.Println(err)
		os.Exit(1)
	}
	name, _ := pterm.DefaultInteractiveTextInput.WithDefaultText("Provide app name").Show()
	port, _ := pterm.DefaultInteractiveTextInput.WithDefaultText("Provide port to use").WithDefaultValue("80").Show()
	p, err := strconv.Atoi(port)
	if err != nil {
		pterm.Error.Println(err)
		os.Exit(1)
	}
	if p < 0 || p > 65535 {
		pterm.Error.Println("Port not in range!")
		os.Exit(1)
	}
	Start(dir, name, uint16(p))
}
