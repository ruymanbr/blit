/*
 * File:    main.go
 *
 * Author:  	Ruymán Borges R. (ruyman21@gmail.com)
 * Date:    	1-7-21
 *
 * Summary of File:
 *
 *  This program interacts with blit CLI, blit backend and blit frontend to handle user CLi and API requests.
 * 	It displays folder and files in a given path, including: Dir (y/n) - Last modified date - Name - Size.
 *  It runs on Linux (Ubuntu, etc)
 *	
 *	CLI operates displaying info in console. Frontend activates when no path has been given. It then operates through a UI
 *	to display a given path through default browser (http://localhost:8080, specifically). It can be altered to work from 
 *	a different location in future versions.
 *
 */

package main

import (
	"github.com/ruymanbr/blit/pkg/blit_backend"
	"github.com/ruymanbr/blit/pkg/blit_cli"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"log"
	"io/fs"
)

var CLI_active bool

func init() {
	if len(os.Args) > 1 {
		CLI_active = true
	} else {
		CLI_active = false
	}
}

func main() {
	if CLI_active {
		path, _ := blit_cli.GetPath(os.Args)

		fileInfo, pathCorrect, err	:= HandlePath(path)
		if err != nil {
			log.Fatalf("Couldn't extract any info from %v. Error: %v\n", path, err)
		}
		fmt.Printf("Passing fileInfo length %v with path '%v' to EncapData()\n", len(fileInfo), pathCorrect)
		encap_data, err, totSize 	:= blit_cli.EncapData(fileInfo, pathCorrect)
		if err != nil {
			log.Fatalf("\nPath %v is incorrect. Error: %v\n", path, err)
		}
		sizesSli 					:= blit_cli.EncapSizes(fileInfo)
		totFiles 					:= len(sizesSli)
		_, dirList 					:= blit_cli.CleanData(encap_data)

		blit_cli.FileSizeSort(sizesSli, 1)
		sortedSli					:= blit_cli.FastSwitchSli(encap_data, sizesSli, 0)

		blit_cli.RenderData(dirList, sortedSli, totSize, totFiles)

	} else {		
		Openbrowser("http://localhost:8080")
		blit_backend.Start()
	}

}

// HandlePath handles a given path calling functions in package blit_cli
// 
// Takes 1 argument:
// 1: path string	 	(what system path to be listed)
//
// Returns:
//	1: []fs.FileInfo	(Data from files listed)
//	2: string			(Sanitized path. Returned from  SanitizeLastSlash() with proper slashing format)
//	3: error 			(Returns this error when trying to obtain os.Stat(/path/to/file/name/) for each file
func HandlePath(path string) ([]fs.FileInfo, string, error) {
	var fileInfo []fs.FileInfo
	
	//fmt.Println("Alternative paths: ", paths)
	path = SanitizeLastSlash(path)
	fmt.Println("Trying path: ", path)
	
	fileInfo, err := blit_cli.GetPathInfo(path)	

	return fileInfo, path, err
}

// SanitizeLastSlash verifies that last slash is added to given path or returns it with it
//
// Takes 1 argument:
//	1: path string		(what system path to be listed)
//
// Returns:
//	1: string			(Sanitized path with slash at the end)
func SanitizeLastSlash(path string) string {
	if path[len(path)-1:] != "/" {
		fmt.Println("Last char of path is: ", path[len(path)-1:])
		path += "/"
	}

	if path[:1] != "/" {
		fmt.Println("First char of path is: ", path[:1])
		path = "/" + path
	}
	return path
}


// Openbrowser opens default browser in system at a given URL
// 
// Takes 1 argument:
// 1: url string	 	(what URI to open in brwoser)
//
// Returns: 
//	<No Return>
func Openbrowser(url string) {
	var err error

	switch runtime.GOOS {
	case "linux":
		err = exec.Command("xdg-open", url).Start()
	default:
		err = fmt.Errorf("Unsupported platform")
	}
	if err != nil {
		log.Fatal(err)
	}
}