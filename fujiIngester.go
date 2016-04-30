package main

import (
	"log"
	"path/filepath"
	"fmt"
	"flag"
)

// error checker
func checkError(err error) {
	if err != nil {
		log.Fatal("ERROR: ",err)
	}
}

// takes file path as string and returns string array with file paths
func findFiles(path string) []string {
	files, err := filepath.Glob(path)
	checkError(err)
	return files
}

//takes string array from findFiles and prints them to std out. If string array is nil then error gehts thrown
func printFiles(searchPattern *string) {
	files := findFiles(*searchPattern)
	if files == nil {
		log.Fatal("ERROR: no files found. Search pattern was:" + *searchPattern + "\nCheck your search pattern and try again\n")
	}
	for _, i := range files {
		fmt.Println(i)
	}
}

func main() {
	// Commandline Flag Definition
	filesToLookforPtr := flag.String("f", "./*", "file pattern to look for")
	flag.Parse()

	printFiles(filesToLookforPtr)

	

}