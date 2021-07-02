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
		err := ShowFilesInShell(path)
		if err != nil {
			HandleMalformedPath(path)
		}
	} else {
		
		Openbrowser("http://localhost:8080")
		blit_backend.Start()
	}

}

// ShowFilesInShell passes a given path string to a handler function in CLI package for blit
// 
// Takes 1 argument:
// 1: path string	 		(what URI to open in browser)
//
// Returns: 
//	<No Return>
func ShowFilesInShell(path string) error {	
	err := blit_cli.Handler(path)
	
	return err
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

// TrySlashBefore include 1 slash string before old path string and retry ShowFilesInShell()
// 
// Takes 1 argument:
// 1: oldPath string	 	(what URI to open in browser)
//
// Returns: 
//	1: error 				(In case it can't open the modified path string)
func TrySlashBefore(oldPath string) error {
	newPath := "/" + oldPath
	err := ShowFilesInShell(newPath)
	if err != nil {
		return TrySlashEnd(oldPath)
	}

	return err
}

// TrySlashEnd include 1 slash string after old path string and retry ShowFilesInShell()
// 
// Takes 1 argument:
// 1: oldPath string	 	(what URI to open in browser)
//
// Returns: 
//	1: error 				(In case it can't open the modified path string)
func TrySlashEnd(oldPath string) error {
	newPath := oldPath + "/"
	err := ShowFilesInShell(newPath)
	if err != nil {
		return TryBothSlashes(oldPath)
	}

	return err
}

// TryBothSlashes include slashes string before and after and retry ShowFilesInShell()
// 
// Takes 1 argument:
// 1: oldPath string	 	(what URI to open in browser)
//
// Returns: 
//	1: error 				(In case it can't open the modified path string)
func TryBothSlashes(oldPath string) error {
	newPath := "/" + oldPath + "/"
	err := ShowFilesInShell(newPath)

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