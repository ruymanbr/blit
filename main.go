/*
 * File:    main.go
 *
 * Author:  	Ruymán Borges R. (ruyman21@gmail.com)
 * Date:    	1-7-21
 *
 * Summary of File:
 *
 *  This program interacts with blit CLI, blit backend and blit frontend to handle user CLi and API requests.
 * 	
 * 	It lists files and folders ordered by size and displays size and last modification date of each element, count of files
 *	and total size of path folder and files.
 *  
 *	It operates with CLI and also through a frontend (based on React.js)
 * 
 *	It runs on a unix system (i.e. ubuntu)
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

		fileInfo, pathCorrect, err	:= HandlePathFormat(path)
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

// HandlePathFormat handles a given path calling functions in package blit_cli
// 
// Takes 1 argument:
// 1: path string	 		(what URI to open in browser)
//
// Returns:
//	1: []fs.FileInfo	(Data from files listed)
//	3: error 			(Returns this error when trying to obtain os.Stat(/path/to/file/name/) for each file
func HandlePathFormat(path string) ([]fs.FileInfo, string, error) {
	var fileInfo []fs.FileInfo
	SlashBefore := "/" + path
	SlashAfter 	:= path + "/"
	BothSlashes	:= "/" + path + "/"

	var paths = []struct{
		newPath string
	}{
		{SlashBefore},
		{SlashAfter},
		{BothSlashes},
	}
	
	fileInfo, err := blit_cli.GetPathInfo(path)	

	if err != nil {
		fmt.Println("Error: ", err)
		for _, tryPath := range paths {
			
			fileInfo, err = blit_cli.GetPathInfo(tryPath.newPath)
			
			if err == nil {
				if tryPath.newPath[len(tryPath.newPath)-1:] != "/" {
					tryPath.newPath += "/"
				}			
				return fileInfo, tryPath.newPath, nil
			}
		}
	}
	return fileInfo, path, err
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