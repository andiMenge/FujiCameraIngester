package main

import (
	"fmt"
	"github.com/kr/fs"
	"github.com/urfave/cli"
	"log"
	"os"
	"os/exec"
	"regexp"
)

//Vars

// RegEx pattern that is used to find files
const JpgRegExPattern = "\\.(JPG)"
const RafRegExPattern = "\\.(RAF)"

//here can you define your working folders. They will be created by func createDirs()
var destPaths = map[string]string{
	"jpegPath": "1_source/Fuji_XE2_JPEG",
	"rawPath":  "1_source/Fuji_XE2_RAW",
	"outPath":  "2_delivery",
}

//debug functions
func printSlice(pathSlice []string) {
	fmt.Printf("DEBUG: Slice stats:\nlen=%d cap=%d \n", len(pathSlice), cap(pathSlice))
}

// error checker
func checkError(err error) {
	if err != nil {
		log.Fatal("ERROR: ", err)
	}
}

//parses searchPath recursive and looks for file endings specified in egExPattern.
//if found filepath gets extended to to slice array. Returns slice array with absolute file paths.
func findFiles(path string) ([]string, []string) {
	JpgPathSlice := make([]string, 0)                //create empty string slice for JPG files
	RafPathSlice := make([]string, 0)                //create empty string slice for RAF files
	JpgRegEx, err := regexp.Compile(JpgRegExPattern) // compile regEx pattern
	checkError(err)                                  //checks for errors from regex compiling
	RafRegEx, err := regexp.Compile(RafRegExPattern) // compile regEx pattern
	checkError(err)                                  //checks for errors from regex compiling
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
	pwd, err := os.Getwd()
	checkError(err)
	for i := range destPaths { //gets keys stored in map
		fmt.Printf("INFO: creating folder: %s/%s\n", pwd, destPaths[i])
		err := os.MkdirAll(destPaths[i], 0755)
		checkError(err)
	}
}

// takes src and destination path as input and copys files with cp tool
func copyFiles(src []string, dest string) {
	bin, lookErr := exec.LookPath("cp") //finds absolut path to binary
	checkError(lookErr)
	args := "-av" //args for cp command
	for _, i := range src {
		copyCmd := exec.Command(bin, args, i, dest) //preps cp cmd
		cmdOut, cpErr := copyCmd.Output()           // executes cp cmd collects stdin from cp cmd
		checkError(cpErr)
		fmt.Printf(string(cmdOut)) //prints stdout from cp to terminal
	}
}

//main programm flow
func mainFunc(searchPath string) {
	createDirs()
	JpgFilePathSlice, RafFilePathSlice := findFiles(searchPath)
	copyFiles(JpgFilePathSlice, destPaths["jpegPath"])
	copyFiles(RafFilePathSlice, destPaths["rawPath"])
	//DEBUG
	//printSlice(JpgFilePathSlice)
	//printSlice(RafFilePathSlice)
}

func main() {
	app := cli.NewApp()
	app.Name = "fujiCp"
	app.Usage = "copies images from fuji cameras"
	app.Version = "1.0.0"
	app.UsageText = "fujiCp /path/to/source/folder/or/sdCard"

	//calls main function
	app.Action = func(c *cli.Context) error {
		searchPath := c.Args().Get(0)
		mainFunc(searchPath)
		return nil
	}

	app.Run(os.Args)
}
