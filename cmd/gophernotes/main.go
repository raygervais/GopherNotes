package main

import (
	"flag"
	"fmt"
	"github.com/raygervais/gophernotes/pkg/cli"
	"os"
)

var (
	helpFlag = flag.Bool("help", false, "Display application usage material")
)

func main() {
	flag.Parse()

	cli := cli.InitCLI()

	if *helpFlag || len(os.Args) == 1 {
		cli.Help()
		return
	}

	if err := cli.Handler(); err != nil {
		fmt.Printf("Command error: %v\n", err)
		os.Exit(1)
	}
}
