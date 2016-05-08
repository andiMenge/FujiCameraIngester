package main

import (
	"fmt"
	"github.com/kr/fs"
	"log"
	"flag"
	"regexp"
	"os"
)

//debug functions
func printSlice(pathSlice []string) {
	fmt.Printf("DEBUG: Slice stats:\nlen=%d cap=%d \n", len(pathSlice), cap(pathSlice))
}

//Vars
const JpgRegExPattern = "\\.(JPG)"
const RafRegExPattern = "\\.(RAF)"

// error checker
func checkError(err error) {
	if err != nil {
		log.Fatal("ERROR: ",err)
	}
}

//parses searchPath recursive and looks for file endings specified in egExPattern.
//if found filepath gets extended to to slice array. Returns slice array with absolute file paths.
func findFiles(path string) ([]string, []string) {
	JpgPathSlice := make([]string, 0) //create empty string slice for JPG files
	RafPathSlice := make([]string, 0) //create empty string slice for RAF files
	JpgRegEx, err := regexp.Compile(JpgRegExPattern) // compile regEx pattern
	checkError(err) //checks for errors from regex compiling
	RafRegEx, err := regexp.Compile(RafRegExPattern) // compile regEx pattern
	checkError(err) //checks for errors from regex compiling
	walker := fs.Walk(path)
	for walker.Step() {
	    err := walker.Err()
	    checkError(err) //checks errors from walker.Step
	  	switch {
	  	case JpgRegEx.MatchString(walker.Path()) == true:
	  		JpgPathSlice = append(JpgPathSlice, walker.Path())
	  	case RafRegEx.MatchString(walker.Path()) == true:
	  		RafPathSlice = append(RafPathSlice, walker.Path())
	  	}
	}
	return JpgPathSlice, RafPathSlice //returns slice of paths
}

// create target folder in ./
func createDirs() {
	folderNames := [2]string {"Fuji_XE2_JPEG", "Fuji_XE2_RAW"}
	pwd, err := os.Getwd()
	checkError(err)
	for _, i := range folderNames {
		fmt.Println("INFO: creating folder:" + pwd + "/" + i)
		err := os.Mkdir(i, 0755)
		checkError(err)
	}
}

func main() {
	flag.Parse()
	searchPath := flag.Arg(0)
	createDirs()
	JpgFilePathSlice, RafFilePathSlice := findFiles(searchPath)
	printSlice(JpgFilePathSlice)
	printSlice(RafFilePathSlice)
}