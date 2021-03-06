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
	"github.com/ruymanbr/blit/pkg/blit_api"
	"github.com/ruymanbr/blit/pkg/blit_cli"
	"os"
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

		fileInfo, pathCorrect, err	:= blit_cli.HandlePath(path)
		if err != nil {
			log.Fatalf("Couldn't extract any info from %v. Error: %v\n", path, err)
		}
		encap_data, sizesSli, err, totSize 	:= blit_cli.EncapData(fileInfo, pathCorrect)
		if err != nil {
			log.Fatalf("\nPath %v is incorrect. Error: %v\n", path, err)
		}
		totFiles 					:= len(sizesSli)
		_, dirList 					:= blit_cli.CleanData(encap_data)

		blit_cli.FileSizeSort(sizesSli, 1)
		sortedSli					:= blit_cli.FastSwitchSli(encap_data, sizesSli, 0)	
		
		blit_cli.RenderData(dirList, sortedSli, totSize, totFiles)

	} else {		
		app := blit_api.App{}
	    app.InitRouter()
	}

}