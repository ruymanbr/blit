/*
 * File:    main.go
 *
 * Author:  	Ruymán Borges R. (ruyman21@gmail.com)
 * Date:    	1-7-21
 *
 * Summary of File:
 *
 *  This program interacts with blit CLI, blit backend and blit frontend to handle user CLi and API requests
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

		sizesSli, encap_data, err, totSize	:= HandlePath(path)	
		totFiles 							:= len(sizesSli)
		_, dirList 							:= blit_cli.CleanData(encap_data)

		blit_cli.FileSizeSort(sizesSli, 1)

		sortedSli							:= blit_cli.FastSwitchSli(encap_data, sizesSli, 0)

		blit_cli.RenderData(dirList, sortedSli, totSize, totFiles)

	} else {		
		Openbrowser("http://localhost:8080")
		blit_backend.Start()
	}

}

// HandlePath handles a given path calling functions in package blit_cli
// 
// Takes 1 argument:
// 1: path string	 		(what URI to open in browser)
//
// Returns:
//	1: [][]int 			(File sizes matrix)
//	2: [][]string 		(File info -as in [n_files]{isDir, lastM, fName, size_HR_Format}  )
//	3: error 			(Returns this error when trying to obtain os.Stat(/path/to/file/name/) for each file
//	4: int64 			(Sum of total file sizes in given path)
func HandlePath(path string) ([][]int, [][]string, error, int64) {	
	//err := blit_cli.Handler(path)
	fileInfo, err := blit_cli.GetPathInfo(path)	

	if err != nil {
		HandleMalformedPath(path)
	}

	return blit_cli.EncapData(fileInfo, path)
}

// HandleMalformedPath starts a cascade function call to try to recover from a malformed path (missing slashes)
// 
// Takes 1 argument:
// 1: oldPath string	 	(what URI to open in browser)
//
// Returns: 
//	<No Return>
func HandleMalformedPath(oldPath string) {
	err := TrySlashBefore(oldPath)

	if err != nil {
		log.Fatalln("Program interrupted. Error: ", err)
	}
}

// TrySlashBefore include 1 slash string before old path string and retry HandlePath()
// 
// Takes 1 argument:
// 1: oldPath string	 	(what URI to open in browser)
//
// Returns: 
//	1: error 				(In case it can't open the modified path string)
func TrySlashBefore(oldPath string) error {
	newPath := "/" + oldPath
	err := HandlePath(newPath)
	if err != nil {
		return TrySlashEnd(oldPath)
	}

	return err
}

// TrySlashEnd include 1 slash string after old path string and retry HandlePath()
// 
// Takes 1 argument:
// 1: oldPath string	 	(what URI to open in browser)
//
// Returns: 
//	1: error 				(In case it can't open the modified path string)
func TrySlashEnd(oldPath string) error {
	newPath := oldPath + "/"
	err := HandlePath(newPath)
	if err != nil {
		return TryBothSlashes(oldPath)
	}

	return err
}

// TryBothSlashes include slashes string before and after and retry HandlePath()
// 
// Takes 1 argument:
// 1: oldPath string	 	(what URI to open in browser)
//
// Returns: 
//	1: error 				(In case it can't open the modified path string)
func TryBothSlashes(oldPath string) error {
	newPath := "/" + oldPath + "/"
	err := HandlePath(newPath)

	return err
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