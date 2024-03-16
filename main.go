package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
)

func handleError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
func main() {
	helpCommands := []string{"-h", "h", "-help", "help"}
	for _, arg := range os.Args[1:] {
		for _, command := range helpCommands {
			if arg == command {
				fmt.Println("example gocrop usage: gocrop <file> <width> <height> <left> <top>")
				return
			}
		}
	}

	flag.Parse()

	if len(flag.Args()) != 5 {
		log.Fatal("Invalid number of arguments")
	}
	file := flag.Arg(0)
	width, err := strconv.Atoi(flag.Arg(1))
	handleError(err)
	height, err := strconv.Atoi(flag.Arg(2))
	handleError(err)
	left, err := strconv.Atoi(flag.Arg(3))
	handleError(err)
	right, err := strconv.Atoi(flag.Arg(4))
	handleError(err)

	fmt.Println(file, width, height, left, right)
}
