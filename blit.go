/*
 * File:    blit.go
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

package main

import (
	"fmt"
	"os"
    "strconv"
    "github.com/olekukonko/tablewriter"
    "io/fs"
)

func main() {
	var err error
    path, cli_on := GetPath(os.Args)
    
	if (cli_on) {
		fmt.Println("Path: ", path)
	} else {
		fmt.Println("Current folder: ", path)

	}	
	sizesSli, encap_data, err, totSize, totFiles	:= GetPathInfo(path, cli_on)

    if err != nil {
    	panic(err)
	} else {
		_, dirList 		:= CleanData(encap_data)
		FileSizeSort(sizesSli, 1)
		sortedSli		:= FastSwitchSli(encap_data, sizesSli, 0)
		RenderData(dirList, sortedSli, totSize, totFiles)
	}
	
}

// GetPath extracts path from CLI argument, if not given it returns current directory path
func GetPath(a []string) (string, bool) {
	curr_wd, err := os.Getwd()
	if err != nil {
		panic(err)
	} else {
		switch len(a) {
			case 2: 				
				return a[1], true
			default: 
				return curr_wd, false
		}	
	}
	return curr_wd, false
	
}

// GetPathInfo extracts a info from files in a given dir path and returns in cascade from EncapData() as -> Matrix [][]int for Sizes; [][]string for File Data as [n_files]{isDir, lastM, fName, size_HR_Format}; error(err); total file_size(int64); total files(int)
func GetPathInfo(root string, cli_ON bool) ([][]int, [][]string, error, int64, int) {

	f, err := os.Open(root)
	if err != nil {
		panic(err)
	}
	fileInfo, err := f.Readdir(-1)
	f.Close()
	if err != nil {
		panic(err)
	}
	if (cli_ON) {
		return EncapData(fileInfo, root)
	} else {
		return EncapData(fileInfo, "")
	}
	return EncapData(fileInfo, "")
}

// EncapData extracts data from a []fs.FileInfo dataset in a given path. Returns values: Matrix [][]int for Sizes; [][]string for File Data as [n_files]{isDir, lastM, fName, size_HR_Format}; error(err); total file_size(int64); total files(int)
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

// CleanData removes first column for [][]string matrix as in returned by EncapData() second returned value, and returns the separated datasets: [][]string, []string
func CleanData(rawData [][]string) ([][]string, []string) {
	var cleanSli [][]string						// 2D slice matrix without the y/n directory column
	var dirSli []string							// []int slice matrix containing y/n directory column
	for _, line := range rawData {
		cleanSli	= append(cleanSli, line[1:])
		dirSli 		= append(dirSli, line[0])
	}	

	return cleanSli, dirSli
}

// ByteToReadableSize transform a byte size into human readable form sizes (kb, Mb, Gb, Tb, Pb)
func ByteToReadableSize(b int64) string {
    const unit = 1024
    if b < unit {
        return fmt.Sprintf("%d  B", b)
    }
    div, exp := int64(unit), 0
    for n := b / unit; n >= unit; n /= unit {
        div *= unit
        exp++
    }
    return fmt.Sprintf("%.1f %cb",
        float64(b)/float64(div), "KMGTPE"[exp])
}

// RenderData outputs in Shell CLI a [][]string dataset as a table
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
	table.SetFooter([]string{"", "", strconv.Itoa(totFiles) + " files", tSizeStr + " (total)"}) // Add Footer
	
	table.SetBorder(false)                              // Set Border to false
	table.AppendBulk(data)                              // Add Bulk Data
	table.Render()
}

// FileSizeSort sorts a [][]int slice matrix as first argument (Bigger first, smaller last). '<sizePos> int' gives the position of filesize (int type)
func FileSizeSort(sli [][]int, sizePos int)  {
	var sorted bool = false
	var i, sorted_i int
	
	for sorted == false {		
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

// Swap switches positions of [][]int slice rows. Rows swapped are i and i+1 index (Takes i int as second argument) 
func Swap(sli [][]int, i int) {
	var row1, row2 []int

	row1		= sli[i]
	row2 		= sli[i+1]

	sli[i] 		= row2
	sli[i+1]	= row1
}

// FastSwitchSli sorts a [n_files][5]string dataset returned in cascade from <- GetPathInfo() <- EncapData(). Takes ([n_files][5]string [n_files][2]int, int) as arguments. 2nd argument a sorted blueprint, 3rd for size_val as col_index in 2nd argument
func FastSwitchSli(strUnordered [][]string, orderedSli [][]int, origPos int) [][]string {
	var sortedSli [][]string
	for _, row := range orderedSli {
		sortThisID	:= row[origPos]
		sortedSli 	= append(sortedSli, strUnordered[sortThisID])
	}

	return sortedSli
}