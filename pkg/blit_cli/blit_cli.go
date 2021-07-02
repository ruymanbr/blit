/*
 * File:    blit_cli.go
 *
 * Author:  	Ruymán Borges R. (ruyman21@gmail.com)
 * Date:    	30-6-21
 *
 * Summary of File:
 *
 *  This program lists the folders and files in a source folder. The list is ordered by size, 
 *	and displays size and last modification date of each element. A count of the files and the 
 *	total size is also provided.
 *  
 *	It comes with a CLI to choose the folder and output the result
 * 
 *	It runs on a unix system (i.e. ubuntu)
 *
 */

package blit_cli

import (
	"fmt"
	"os"
    "strconv"
    "github.com/olekukonko/tablewriter"
    "io/fs"    
)

type PathError struct {
	err error
	path string
}

func (p *PathError) Error() string {
	return fmt.Sprintf("Path %v ...is incorrect", p.path)
}

// Handler calls blit_cli functions to process request for a given path.
// 
// Takes 1 argument:
// 1: path []string 	(path to handle)
//
// Returns: 
//	1: error 			(if path is not a valid or existant folder in system)
func Handler(path string) error {
    fmt.Println("Trying path: ", path)
	
	
	fileInfo, err := GetPathInfo(path)

    if err != nil {
    	return err
	}
	sizesSli, encap_data, err, totSize, totFiles := EncapData(fileInfo, path)	
	if err != nil {
    	return err
	}
	_, dirList 		:= CleanData(encap_data)
	FileSizeSort(sizesSli, 1)
	sortedSli		:= FastSwitchSli(encap_data, sizesSli, 0)
	RenderData(dirList, sortedSli, totSize, totFiles)
	return nil	
}

// GetPath extracts path from CLI argument, if not given it returns current directory path
// 
// Takes 1 argument:
// 1: args []string 	(os.Args)
//
// Returns: 
//	1: string 			(argument path or current working directory)
//	2: bool 			(Yes for argument with path from CLI call to blit program)
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
// 1: root string 		(Path to extract info from)
//
// Returns (same as EncapData() :
//	1: []fs.FileInfo	(slice with info from files and folders)
//	2: error 			(not nilfor failing to open or failing reading it)
func GetPathInfo(root string) ([]fs.FileInfo, error) {
	var emptyPath []fs.FileInfo
	f, err := os.Open(root)
	if err != nil {
		return emptyPath, err
	}
	fileInfo, err := f.Readdir(-1)
	defer func() {
		f.Close()
	}()
	if err != nil {
		return emptyPath, err
	}
	//if (cli_ON) {
		return fileInfo, nil
	//} 
	//return EncapData(fileInfo, "")
}

// EncapData extracts data from a []fs.FileInfo dataset in a given path (string). 
//
// 1: fileInfo []fs.FileInfo (obtained from os.Open File -> Readdir()) 
// 2: root string (Path where files are located)
//
// Returns:
//	1: [][]int 			(File sizes matrix)
//	2: [][]string 		(File info -as in [n_files]{isDir, lastM, fName, size_HR_Format}  )
//	3: error 			(Returns this error when trying to obtain os.Stat(/path/to/file/name/) for each file
//	4: int64 			(Sum of total file sizes in given path)
//	5: int (			total number of files in given path)
func EncapData(fileInfo []fs.FileInfo, root string) ([][]int, [][]string, error, int64, int) {
    var files [][]string	// data set of all files scanned
    var sizes [][]int 		// data set of [][original_order, size] of [][]ints
    var totSize int64 		// sum of file sizes
    var totFiles int = 0 	
    var isDir string		// y/n to detect if it's a directory, for latter format


	for i, file := range fileInfo {
		fName := file.Name()
		stats, err := os.Stat(root + fName)

		if err != nil {
			fmt.Println("err: ", err)
			return sizes, files, err, 0, 0
		}

		if stats.IsDir() {
			isDir = "y"
		} else {
			isDir = "n"
		}

		lastM := stats.ModTime().Format("2006-01-02 15:04:05");
	 
		size := file.Size()
		totSize += size
		totFiles += 1
		
		sizeN := int(size)
		fileLine	:= []string{isDir, lastM, fName, ByteToReadableSize(size)}
		files 		= append(files, fileLine)
		
		sizeLine	:= []int{i, sizeN}
		sizes		= append(sizes, sizeLine)
		
		
	}

	return sizes, files, nil, totSize, totFiles
}

// CleanData removes first column for [][]string matrix. Ideally the format returned by EncapData() function in second position
//
// 1: [][]string 	(Raw data from Encap(), including dirs conditional y/n in first colum
// 
// Returns: 
//	1: [][]string 	(Same matrix without first colum)
//	2: []string 	(Folder y/n confirmation string obtained from argument to this function)
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
// 1: bigNum int64 		(size in bytes)
//
// Returns:
//	1: string			(size in human readable form: Pb, Tb, Gb, etc)
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
// 1: []string  		(Slice with y/n values for Directory)
// 2: [][]string 		(Sorted Slice from biggest file to lowest size) 
// 3: int64 			(Total scanned file size combined)
// 4: int 				(Total files in given path)//	
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
// 1: sli [][]int 		(size matrix with size and original position as column values in every row)
// 2: sizePort int 		(as first argument (Bigger first, smaller last) by calling Swap() function
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
//	1: sli[][]int 	(Slice containing file size information in 2 columns)
//	2: i int 		(i and i+1 positions where rows are going to be swapped)
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
// 1: [][]string 	( Unordered string matrix with folder files data)
// 2: [][]int 		( Sorted slice with file size and original position in primitive raw data slice)
// 3: int 			( original position of files, in ordered fileSize slice's rows. Basically its col_index )
//
// Returns: 
//	1: [][]string 	(Fully formatted array with file data. Ordered by size. later derived to RenderData() function for CLI display purpose)
func FastSwitchSli(strUnordered [][]string, orderedSli [][]int, origPos int) [][]string {
	var sortedSli [][]string
	for _, row := range orderedSli {
		sortThisID	:= row[origPos]
		sortedSli 	= append(sortedSli, strUnordered[sortThisID])
	}

	return sortedSli
}