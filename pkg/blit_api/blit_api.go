/*
 * File:    blit_api.go
 *
 * Author:  	Ruymán Borges R. (ruyman21@gmail.com)
 * Date:    	2-7-21
 *
 * Summary of File:
 *
 *  This program runs a backend server for package blit to handle http requests from a foreground counterpart (blit_frontend)
 *
 */


package blit_api

import (
    "encoding/json"
	"github.com/ruymanbr/blit/pkg/blit_cli"
	"log"
    "net/http"
	"io/ioutil"
	"fmt"
	"strconv"
    "github.com/gin-gonic/gin"
    "github.com/gin-contrib/cors"
)

// App export
type App struct {
    
}

const (
	EmptyStr string = ""
)

func (app *App) InitRouter() {
	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
	}))

	api := router.Group("/api/v1")
	{
		api.GET("/", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H {
			"message": "pong",
			})
		})
	}
	api.POST("/post", FrontHandler)
	api.GET("/post", FrontHandler)

	router.Run(":8080")
}

// FrontHandler retrieves a list of files. 
func FrontHandler(c *gin.Context) {
	c.Header("Content-Type", "application/json")

	var response map[string]interface{}
    
    jsonData, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
	    
	    c.JSON(http.StatusInternalServerError, gin.H {
			"message":"Internal Server Error 500",
			"error": err,
		})	
	}
	json.Unmarshal(jsonData, &response)
	
	if response["path"] != nil {
		
		path := fmt.Sprintf("%v", response["path"])
		if (path == "") {
			c.JSON(http.StatusNotFound, gin.H {
				"message":"Not found 404",
			})	
		}

		jsonToSend, totFiles, tSizeStr, err := GetFilesData(path)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H {
				"message":"Unauthorized Path 401",
				"error": err,
			})	
		} else {
			c.JSON(http.StatusOK, gin.H {
				"data": jsonToSend,
				"totFiles": totFiles,
				"totSize": tSizeStr,
			})	
		}		

	} else {
		c.JSON(http.StatusNotFound, gin.H {
			"message":"Not found 404",
		})
	}	
}

// GetFilesData calls for
func GetFilesData(path string) ([]blit_cli.File, string, string, error) {
	if (path == "") {
		err := fmt.Errorf("Empty search is not allowed")
		return []blit_cli.File{}, EmptyStr, EmptyStr, err		
	}
	fileInfo, pathCorrect, err	:= blit_cli.HandlePath(path)
	if err != nil {
		return []blit_cli.File{}, EmptyStr, EmptyStr, err
	}
	
	encap_data, sizesSli, err, totSize 	:= blit_cli.EncapData(fileInfo, pathCorrect)
	if err != nil {
		log.Fatalf("\nPath %v is incorrect. Error: %v\n", path, err)

		return []blit_cli.File{}, EmptyStr, EmptyStr, err
	}
	tSizeStr := blit_cli.ByteToReadableSize(totSize)
	
	totFiles 					:= len(sizesSli)
	totFilesStr					:= strconv.Itoa(totFiles) + " files"

	blit_cli.FileSizeSort(sizesSli, 1)
	sortedSli					:= blit_cli.FastSwitchSli(encap_data, sizesSli, 0)	
	SortedFiles					:= blit_cli.StructurizeFiles(sortedSli)

	return SortedFiles, totFilesStr, tSizeStr, nil
}
