package main

import (
	"flag"
	run "github.com/kimuraz/impactree/cli"
	utils "github.com/kimuraz/impactree/utils"
	"os"
)

func main() {
	cmd := flag.NewFlagSet(os.Args[1], flag.ExitOnError)

	switch cmd.Name() {
	case "help":
		utils.PrintCommandsList()
	case "run":
		run.Run(os.Args[2:])
	}
}
