package utils

import "fmt"

var argsRun = map[string]string{
	"--help":      "-h, Show list of accepted arguments",
	"--json":      "Generate a JSON report",
	"--csv":       "Generate a CSV report",
	"--no-report": "Don't generate a report",
	"--dir":       "-d, Directory to scan",
	"--exclude":   "-e, Exclude files or directories according to a regex pattern",
	"--include":   "-i, Include files or directories according to a regex pattern",
	"--verbose":   "-v, Verbose mode",
}

// List of accepted commands and its description
var commandsList = map[string]map[string]string{
	"help": map[string]string{
		"desc": "Show this help",
	},
	"run": map[string]string{
		"desc":   "Run the scanner and generate the report depending on the arguments",
		"--help": "Show list of accepted arguments",
	},
}

func PrintCommandsList() {
	for cmd, desc := range commandsList {
		fmt.Printf("%s\t%s\n", cmd, desc["desc"])
	}
}

func PrintRunArgs() {
	for arg, desc := range argsRun {
		fmt.Printf("%s\t%s\n", arg, desc)
	}
}
