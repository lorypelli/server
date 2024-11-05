package internal

import (
	"os"

	"github.com/pterm/pterm"
)

func Exit(err error) {
	pterm.Error.Println(err)
	os.Exit(1)
}
