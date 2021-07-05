/*
 * File:    blit_cli.go
 *
 * Author:  	Ruymán Borges R. (ruyman21@gmail.com)
 * Date:    	30-6-21
 *
 * Summary of File:
 *
 *  This package include functions to operate as a CLI to list folders and files. 
 *
 */

package blit_cli

import (
	"fmt"
	"os"
    "strconv"
    "github.com/olekukonko/tablewriter"
    "io/fs"
	"runtime" 
	"os/exec"
	"log"
	"path/filepath"
)

type PathError struct {
	err error
	path string
}

type File struct {
	IsDir string `json:"IsDir"`
	LastM string `json:"LastM"`
	FName string `json:"FName"`
	FSize string `json:"FSize"`
}

func (p *PathError) Error() string {
	return fmt.Sprintf("Path %v ...is incorrect", p.path)
}

// GetPath extracts path from CLI argument, if not given it returns current directory path
// 
// Takes 1 argument:
// 1: args []string 			(os.Args)
//
// Returns: 
//	1: string 					(argument path or current working directory)
//	2: bool 					(Yes for argument with path from CLI call to blit program)
func GetPath(args []string) (string, bool) {
	curr_wd, err := os.Getwd()
	if err != nil {
		panic(err)
	} else {
		switch len(args) {
			case 2: 				
				return args[1], true
		}	
	}
	return curr_wd, false	
}

// GetPathInfo extracts info from a given path. 
// 
// Takes 1 argument:
// 1: root string 				(Path to extract info from)
//
// Returns (same as EncapData() :
//	1: []fs.FileInfo			(slice with info from files and folders)
//	2: error 					(not nilfor failing to open or failing reading it)
func GetPathInfo(root string) ([]fs.FileInfo, error) {
	var emptyPath []fs.FileInfo
	f, err := os.Open(root)
	if err != nil {
		fmt.Println("Error after os.Open: ", err)
		return emptyPath, err
	}
	fileInfo, err := f.Readdir(-1)
	defer func() {
		f.Close()
	}()
	if err != nil {
		fmt.Println("Error after f.Readdir: ", err)
		return emptyPath, err
	}	
	return fileInfo, nil	
}

// DirSize obtains Dir size recursively
//
// 1: path string 				(Path where files are located)
//
// Returns:
//	1: int64 					(Sum of total file sizes in given path)
func DirSize(path string) (int64, error) {
	var size int64
    err := filepath.Walk(path, func(_ string, info os.FileInfo, err error) error {
        if err != nil {
            return err
        }
        if !info.IsDir() {
            size += info.Size()
        }
        return err
    })
    return size, err
}

// EncapData extracts data from a []fs.FileInfo dataset in a given path
//
// 1: fileInfo []fs.FileInfo 	(obtained from os.Open File -> Readdir()) 
// 2: path string 				(Path where files are located)
//
// Returns:
//	2: [][]string 				(File info -as in [n_files]{IsDir, LastM, FName, FSize_HR_Format}  )
//	3: error 					(Returns this error when trying to obtain os.Stat(/path/to/file/name/) for each file
//	3: int64 					(Sum of total file sizes in given path)
func EncapData(fileInfo []fs.FileInfo, path string) ([][]string, error, int64) {
    var files [][]string	// data set of all files scanned
    var totSize int64 		// sum of file sizes
    var IsDir string		// y/n to detect if it's a directory, for latter format

	for _, file := range fileInfo {
		var FSize int64
		FName := file.Name()
		stats, err := os.Stat(path + FName)

		if err != nil {
			fmt.Println("os.Stat failed. Err: ", err)
			return files, err, 0
		}

		if stats.IsDir() {
			IsDir = "y"
			DirPath := path + FName
			FSize, err = DirSize(DirPath)
			if err != nil {
				fmt.Println("DirSize() failed to obtain Dir total size. Are you Super User? (Try 'sudo su -' and run again) Err: ", err)
				fmt.Println("Remember to call: 'export PATH=$PATH:/usr/local/go/bin', as super user")
				return files, err, 0
			}
			totSize += FSize
		} else {
			IsDir = "n"
			FSize = file.Size()
			totSize += FSize
		}
		LastM := stats.ModTime().Format("2006-01-02 15:04:05");
		fileLine	:= []string{IsDir, LastM, FName, ByteToReadableSize(FSize)}
		files 		= append(files, fileLine)
		
	}
	return files, nil, totSize
}

// EncapSizes returns a [][]int slice with data from a []fs.FileInfo dataset in a given path
//
// 1: fileInfo []fs.FileInfo (obtained from os.Open File -> Readdir()) 
//
// Returns:
//	1: [][]int 					(File sizes matrix)
func EncapSizes(fileInfo []fs.FileInfo) ([][]int) {
	var sizes [][]int

	for i, file := range fileInfo {
		size := file.Size()		
		sizeN := int(size)			
		sizeLine	:= []int{i, sizeN}
		sizes		= append(sizes, sizeLine)
	}
	return sizes
}

// CleanData removes first column for [][]string matrix. Ideally the format returned by EncapData() function in second position
//
// 1: rawData [][]string 		(Raw data from Encap(), including dirs conditional y/n in first colum
// 
// Returns: 
//	1: [][]string 				(Same matrix without first colum)
//	2: []string 				(Folder y/n confirmation string obtained from argument to this function)
func CleanData(rawData [][]string) ([][]string, []string) {
	var cleanSli [][]string						// 2D slice matrix without the y/n directory column
	var dirSli []string							// []int slice matrix containing y/n directory column
	for _, line := range rawData {
		cleanSli	= append(cleanSli, line[1:])
		dirSli 		= append(dirSli, line[0])
	}	

	return cleanSli, dirSli
}

// ByteToReadableSize transform a byte size into human readable form sizes (kb, Mb, Gb, Tb, Pb). Takes 1 argument and returns a HR string for size
//
// 1: bigNum int64 				(size in bytes)
//
// Returns:
//	1: string					(size in human readable form: Pb, Tb, Gb, etc)
func ByteToReadableSize(bigNum int64) string {
    const unit = 1024
    if bigNum < unit {
        return fmt.Sprintf("%d  B", bigNum)
    }
    div, exp := int64(unit), 0
    for n := bigNum / unit; n >= unit; n /= unit {
        div *= unit
        exp++
    }
    return fmt.Sprintf("%.1f %cb",
        float64(bigNum)/float64(div), "KMGTPE"[exp])
}

// RenderData renders a table in CLI. Takes 4 arguments with information from Files in path given as first argument to the program
//
// 1: []string  				(Slice with y/n values for Directory)
// 2: [][]string 				(Sorted Slice from biggest file to lowest size) 
// 3: int64 					(Total scanned file size combined)
// 4: int 						(Total files in given path)//	
//	
//	<No return>
func RenderData(dirs []string, data [][]string, totSize int64, totFiles int) {

	tSizeStr := ByteToReadableSize(totSize)

	table := tablewriter.NewWriter(os.Stdout)

	table.SetHeaderAlignment(tablewriter.ALIGN_RIGHT)
	
	table.SetAlignment(tablewriter.ALIGN_RIGHT) 

	table.SetHeader([]string{"Is Dir y/n", "LastModified", "Name", "Size"})
	table.SetHeaderColor(tablewriter.Colors{tablewriter.FgHiWhiteColor},
	tablewriter.Colors{tablewriter.FgHiWhiteColor},
	tablewriter.Colors{tablewriter.FgHiWhiteColor},
	tablewriter.Colors{tablewriter.FgHiWhiteColor})

	table.SetColumnColor(tablewriter.Colors{tablewriter.FgHiWhiteColor},
	tablewriter.Colors{tablewriter.FgHiWhiteColor},
	tablewriter.Colors{tablewriter.FgHiWhiteColor},
	tablewriter.Colors{tablewriter.FgHiWhiteColor})
	table.SetFooter([]string{"", "", strconv.Itoa(totFiles) + " files", tSizeStr + " (total)"})
	
	table.SetBorder(false) 
	table.AppendBulk(data) 
	table.Render()
}

// FileSizeSort sorts a [][]int slice matrix of file data, by size.
// 
// Takes 2 arguments:
// 
// 1: sli [][]int 				(size matrix with size and original position as column values in every row)
// 2: sizePort int 				(as first argument (Bigger first, smaller last) by calling Swap() function
//	
//	<No return>
func FileSizeSort(sli [][]int, sizePos int)  {
	var sorted bool = false
	var i, sorted_i int
	
	for !sorted {		
		sorted_i = 0
		for i = 0; i < (len(sli) - 1); i++ {
			if sli[i][sizePos] < sli[i+1][sizePos] {
				Swap(sli, i)
				sorted_i += 1
			}
		}

		if sorted_i == 0 {
			sorted = true
		}			
	}		
}

// Swap switches positions of 2 rows from [][]int slice. Rows swapped are i and i+1 index (Takes i int as second argument) 
// 
//	Takes 2 arguments:
//	1: sli[][]int 				(Slice containing file size information in 2 columns)
//	2: i int 					(i and i+1 positions where rows are going to be swapped)
// 
//	<No return>
func Swap(sli [][]int, i int) {
	var row1, row2 []int

	row1		= sli[i]
	row2 		= sli[i+1]

	sli[i] 		= row2
	sli[i+1]	= row1
}

// FastSwitchSli sorts a [n_files][5]string dataset obtained from <- GetPathInfo() <- EncapData().
// 
// Takes 3 arguments:
// 1: [][]string 				( Unordered string matrix with folder files data)
// 2: [][]int 					( Sorted slice with file size and original position in primitive raw data slice)
// 3: int 						( original position of files, in ordered fileSize slice's rows. Basically its col_index )
//
// Returns: 
//	1: [][]string 				(Fully formatted array with file data. Ordered by size. later derived to RenderData() function for CLI display purpose)
func FastSwitchSli(strUnordered [][]string, orderedSli [][]int, origPos int) [][]string {
	var sortedSli [][]string
	for _, row := range orderedSli {
		sortThisID	:= row[origPos]
		sortedSli 	= append(sortedSli, strUnordered[sortThisID])
	}

	return sortedSli
}

// HandlePath handles a given path calling functions in package blit_cli
// 
// Takes 1 argument:
// 1: path string	 			(what system path to be listed)
//
// Returns:
//	1: []fs.FileInfo			(Data from files listed)
//	2: string					(Sanitized path. Returned from  SanitizeLastSlash() with proper slashing format)
//	3: error 					(Returns this error when trying to obtain os.Stat(/path/to/file/name/) for each file
func HandlePath(path string) ([]fs.FileInfo, string, error) {
	var fileInfo []fs.FileInfo
	
	path = SanitizeLastSlash(path)	
	fileInfo, err := GetPathInfo(path)	

	return fileInfo, path, err
}

// SanitizeLastSlash verifies that last slash is added to given path or returns it with it
//
// Takes 1 argument:
//	1: path string				(what system path to be listed)
//
// Returns:
//	1: string					(Sanitized path with slash at the end)
func SanitizeLastSlash(path string) string {
	if path[len(path)-1:] != "/" {
		path += "/"
	}

	if path[:1] != "/" {
		path = "/" + path
	}
	return path
}


// Openbrowser opens default browser in system at a given URL
// 
// Takes 1 argument:
// 1: url string	 			(what URI to open in brwoser)
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

// StructurizeFiles converts [][]string data from files into []File struct type so it can be converted into Json
// 
// Takes 1 argument:
// 1: files [][]string	 	(files in 2D array format)
//
// Returns: 
//	1: []File 				(Preformated to be json capable)
func StructurizeFiles(filesStr [][]string) []File {
	var Files []File
	for _, file := range filesStr {
		IsDir		:= file[0]
		LastM 		:= file[1]
		FName		:= file[2]
		FSize_HR 	:= file[3]

		var File 	= File{IsDir, LastM, FName, FSize_HR}
		Files 		= append(Files, File)
	}

	return Files
}