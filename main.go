package main

import (
	"flag"
	"fmt"
	"strings"
)

func showHelp() {
	fmt.Println("To crop an image supply the following arguments:")
	fmt.Println("gocrop source.jpg width height left top destination.jpg")
}

func main() {
	flag.Parse()

	if strings.ToLower(flag.Arg(0)) == "help" || flag.NArg() == 0 {
		showHelp()
	}
}
