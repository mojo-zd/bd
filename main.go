package main

import (
	"os"

	"github.com/mojo-zd/bd/cmd"
	"github.com/urfave/cli"
)

func main() {
	parseArgs()

	app := cli.NewApp()
	app.Name = "bd"
	app.Usage = "bai du translate of command line"
	app.Action = cmd.Action
	app.Run(os.Args)
}

func parseArgs() {
	if len(os.Args) == 1 {
		cmd.DisplayLogo()
	} else if len(os.Args) == 2 {
		if os.Args[1] == "-h" || os.Args[1] == "--help" || os.Args[1] == "h" || os.Args[1] == "help" {
			cmd.DisplayLogo()
		}
	}
}
