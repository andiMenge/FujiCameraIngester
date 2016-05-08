package main

import (
	"fmt"
	"github.com/kr/fs"
	"log"
	"flag"
	"regexp"
)

//debug functions
func printSlice(pathSlice []string) {
	fmt.Printf("INFO: Slice stats:\nlen=%d cap=%d %v\n\n", len(pathSlice), cap(pathSlice), pathSlice)
}

//Vars
const regExPattern = "\\.(JPG|RAF)"

// error checker
func checkError(err error) {
	if err != nil {
		log.Fatal("ERROR: ",err)
	}
}

func findFiles(path string) []string {
	pathSlice := make([]string, 0) //create empty slice
	regEx, err := regexp.Compile(regExPattern) // compile regEx pattern from regExPattern
	checkError(err) //checks for errors from regex compiling
	walker := fs.Walk(path)
	for walker.Step() {
	    err := walker.Err()
	    checkError(err) //checks errors from walker.Step
	  	if regEx.MatchString(walker.Path()) == true { //if filepath matches regExPattern path gets added to slice
	  		pathSlice = append(pathSlice, walker.Path())
	  	}
	}
	return pathSlice //returns slice of paths
}

func main() {
	flag.Parse()
	searchPath := flag.Arg(0)
	filePathSlice := findFiles(searchPath)
	printSlice(filePathSlice)
}