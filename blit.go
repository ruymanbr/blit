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
	"log"
	"os"
    "path/filepath"  
    "github.com/ruymanbr/tables.git"
    //"io/ioutil"
)

type File struct {
	name string
	size string
}

type Folder struct {
	name string
}

func main() {
    //var files []string
    //var files []File
    //var folders []Folder

	
	config, cli_on := configPath(os.Args)
	if (cli_on) {
		//fmt.Println("CLI was used")
		//fmt.Println("Path: ", config)
	} else {
		//fmt.Println("CLI NOT used")
		fmt.Println("Current path: ", config)
	}
		
	
	
	//root := "/home/ruyman/go/src/qoH-assignment"
	//root := config
	/*
    
    err := filepath.Walk(root, visit(&files))
    if err != nil {
        panic(err)
    }
    for _, file := range files {
        fmt.Println(file)
    }

    dirname := config
    f, err := os.Open(dirname)
    if err != nil {
        log.Fatal(err)
    }
    files, err := f.Readdir(-1)
    f.Close()
    if err != nil {
        log.Fatal(err)
    }
	*/
	/*
	dirname := config
	//fmt.Println(dirname)
	
    files, err := ioutil.ReadDir(dirname)
    if err != nil {
        log.Fatal(err)
    }

    for _, file := range files {
        fmt.Println(file.Name())
    }
    */

    files_dir, err := OSReadDir(config)
    if err == nil {
    	for _, file := range files_dir {
			//files = append(files, file.Name())
			fmt.Println(file)
		}
	} else {
		fmt.Println("err: ", err)
	}
}

func visit(files *[]string) filepath.WalkFunc {
    return func(path string, info os.FileInfo, err error) error {
        if err != nil {
            log.Fatal(err)
        }
        *files = append(*files, path)
        return nil
    }
}



func configPath(a []string) (string, bool) {
	curr_wd, err := os.Getwd()
	if err != nil {
		fmt.Println("err: ", err)
	} else {
		switch len(a) {
			case 1: return curr_wd, false
			default: 
				return os.Args[1], true
		}	
	}
	return curr_wd, false
	
}

func OSReadDir(root string) ([]string, error) {
    var files []string
	f, err := os.Open(root)
	if err != nil {
		return files, err
	}
	fileInfo, err := f.Readdir()
	f.Close()
	if err != nil {
		return files, err
	}
	for _, file := range fileInfo {
		size := file.Size()
		//fmt.Printf("%v \n", ByteCountSI(size))
		files = append(files, file.Name() + "\t\t" + ByteCountSI(size))
	}

	return files, nil
}

// ByteCountSI transform a byte size into human readable form sizes (kb, Mb, Gb, Tb, Pb)
func ByteCountSI(b int64) string {
    const unit = 1000
    if b < unit {
        return fmt.Sprintf("%d B", b)
    }
    div, exp := int64(unit), 0
    for n := b / unit; n >= unit; n /= unit {
        div *= unit
        exp++
    }
    return fmt.Sprintf("%.1f %cB",
        float64(b)/float64(div), "kMGTPE"[exp])
}