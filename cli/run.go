package cli

import (
	"fmt"
	utils "github.com/kimuraz/impactree/utils"
)

func Run(args []string) {
	if len(args) == 0 {
		utils.PrintRunArgs()
	}
	for idx, arg := range args {
		switch arg {
		case "--help", "-h":
			utils.PrintRunArgs()
		case "--json":
			// Generate JSON report
			fmt.Println("Generating JSON report")
		case "--html":
			// Generate HTML report
			fmt.Println("Generating HTML report")
		case "--csv":
			// Generate CSV report
			fmt.Println("Generating CSV report")
		case "--no-report":
			// Don't generate a report
			fmt.Println("Not generating a report")
		case "--dir", "-d":
			// Directory to scan
			fmt.Println("Scanning directory: " + args[idx+1])
		case "--exclude", "-e":
			// Exclude files or directories according to a regex pattern
			fmt.Println("Excluding files or directories according to a regex pattern: " + args[idx+1])
		}
	}
}
