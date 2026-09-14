package internal

import "os"

func Exit(err error) {
	Error.Println(err)
	os.Exit(1)
}
