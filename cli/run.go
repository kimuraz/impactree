package cli

import (
	"encoding/json"
	"fmt"
	utils "github.com/kimuraz/impactree/utils"
)

func BuildFeatureMap(relations []Rel, verbose bool) map[string][]string {
	if verbose {
		fmt.Println("Building feature tree")
	}
	featureMap := make(map[string][]string)
	for _, rel := range relations {
		for _, feat := range rel.Feat {
			featureMap[feat] = append(featureMap[feat], rel.File)
		}
	}

	return featureMap
}

func BuildFileMap(relations []Rel, verbose bool) map[string]map[string][]string {
	if verbose {
		fmt.Println("Building file tree")
	}
	fileMap := make(map[string]map[string][]string)
	featuresMap := BuildFeatureMap(relations, verbose)
	for _, rel := range relations {
		fileMap[rel.File] = make(map[string][]string)
		fileMap[rel.File]["features"] = append(fileMap[rel.File]["features"], rel.Feat...)
		fileMap[rel.File]["relations"] = append(fileMap[rel.File]["relations"], rel.Rel...)
		for _, feat := range rel.Rel {
			fileMap[rel.File]["impacts_on"] = append(fileMap[rel.File]["impacts_on"], featuresMap[feat]...)
			for _, file := range featuresMap[feat] {
				fileMap[rel.File]["impacted_by"] = append(fileMap[file]["impacted_by"], rel.File)
			}
		}
	}

	return fileMap
}

func BuildImpactTree(relations []Rel, verbose bool) string {
	if verbose {
		fmt.Println("Building impact tree")
	}
	fileMapJson, err := json.Marshal(BuildFileMap(relations, verbose))
	if err != nil {
		panic(err)
	}

	return string(fileMapJson)
}

func Run(args []string) {
	if len(args) == 0 {
		utils.PrintRunArgs()
	}
	var verbose bool = false
	var dirs []string
	var output_dir string = "impact_tree_reports"
	var report_type string = "json"

	for idx, arg := range args {
		switch arg {
		case "--help", "-h":
			utils.PrintRunArgs()
		case "--json":
			// Generate JSON report
			fmt.Println("Generating JSON report")
			report_type = "json"
		case "--csv":
			// Generate CSV report
			fmt.Println("Generating CSV report")
			report_type = "csv"
		case "--no-report":
			// Don't generate a report
			fmt.Println("Not generating a report")
			report_type = "none"
		case "--dir", "-d":
			// Directory to scan
			fmt.Println("Including directory " + args[idx+1] + " to scan")
			dirs = append(dirs, args[idx+1])
		case "--exclude", "-e":
			// Exclude files or directories according to a regex pattern
			fmt.Println("Excluding files or directories according to a regex pattern: " + args[idx+1])
		case "--verbose", "-v":
			// Verbose mode
			fmt.Println("Verbose mode")
			verbose = true
		case "--output-dir", "-o":
			// Output directory
			fmt.Println("Output directory: " + args[idx+1])
			output_dir = args[idx+1]
		}
	}

	// Print settings
	fmt.Println("Settings:")
	fmt.Println("Report type: " + report_type)
	fmt.Println("Output directory: " + output_dir)

	var relations []Rel
	files := GetAllFiles(dirs, verbose)
	for _, file := range files {
		ch := make(chan Rel)
		go ExtractFilesFeatsAndRel(file, ch, verbose)
		relations = append(relations, <-ch)
	}
	report := BuildImpactTree(relations, verbose)
	fmt.Println(report)
}
