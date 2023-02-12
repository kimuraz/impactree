package cli

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// Function that reads a directory and returns a list of files
func ReadDir(dir string, ch chan []string, verbose bool) {
	var filesNames []string
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}
		filesNames = append(filesNames, path)
		if verbose {
			fmt.Println(path)
		}
		return nil
	})
	if err != nil {
		panic(err)
	}
	ch <- filesNames
}

func ExtractFilesFeatsAndRel(file string, ch chan Rel, verbose bool) {
	var rel Rel
	rel.File = file
	f, err := os.Open(file)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	if verbose {
		fmt.Println("Extracting features and relations from " + file)
	}

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, "@feat") {
			rel.Feat = append(rel.Feat, FeatParser(line)...)
			if verbose {
				fmt.Println(line)
			}
		} else if strings.Contains(line, "@rel") {
			rel.Rel = append(rel.Rel, FeatParser(line)...)
			if verbose {
				fmt.Println(line)
			}
		}
	}
	ch <- rel
}

func GetAllFiles(dirs []string, verbose bool) []string {
	var files []string
	for _, dir := range dirs {
		ch := make(chan []string)
		go ReadDir(dir, ch, verbose)
		files = append(files, <-ch...)
	}
	return files
}
