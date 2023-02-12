package cli

import "io/ioutil"

// Function that reads a directory and returns a list of files
func ReadDir(dir string) {
	ioutil.ReadDir(dir)
}
