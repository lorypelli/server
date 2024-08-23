package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	var dir, port string
	fmt.Print("Provide directory to serve: ")
	fmt.Scan(&dir)
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Print("Provide port to use: ")
	fmt.Scan(&port)
	p, err := strconv.Atoi(port)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	if p < 0 || p > 65535 {
		fmt.Println("Port not in range!")
		os.Exit(1)
	}
	Start(dir, uint16(p))
}
