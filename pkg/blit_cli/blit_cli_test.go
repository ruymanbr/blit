/*
 * File:    blit_cli_test.go
 *
 * Author:  	Ruymán Borges R. (ruyman21@gmail.com)
 * Date:    	30-6-21
 *
 * Summary of File:
 *
 *  This program tests blit_cli.go and its functions 
 *  
 *
 */

package blit_cli

import (
	"testing"
	"os"
	"fmt"
	"io/fs"
)

var WhiteNotPassed	= "\033[97m NOT PASSED: "
var NotPassed 		= "\033[31m NOT PASSED: \033[97m"
var Passed 			= "\033[32m PASSED: \033[97m"
var TotalPassed 	= 0
var TotalNotPassed 	= 0

// TestGetPath_tests GetPath from package blit_cli
func TestGetPath(t *testing.T) {
	var tests = []struct {
		args []string
		wantedPath string
		pathWasGiven bool
	}{
		{[]string{"cmd", ""}, "", false},
		{[]string{"cmd", "/home/"}, "/home/", true},
		{[]string{"cmd", "/usr/local/"}, "/usr/local/", true},
		{[]string{"cmd", "/home/yourUserName/"}, "/home/yourUserName/", true},
		{[]string{"cmd", "/etc/"}, "/etc/", true},
	}
	oldArgs := os.Args
	defer func() { os.Args = oldArgs }()
	os.Args = []string{"cmd"}
	
	curr_wd, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	
	for i, test := range tests {
		got, pathGiven := GetPath(test.args)
		
		if (pathGiven) {
			if got != test.wantedPath {
				t.Errorf("%v Test %v - GetPath(%v) == GOT: %v, %v. // WANT: %v, %v\n", NotPassed, i+1, test.args, got, pathGiven, test.wantedPath, test.pathWasGiven)
				TotalNotPassed += 1
			} else {
				fmt.Printf("%v Test %v - GetPath(%v) -> GOT: %v, %v. // WANT: %v, %v\n", Passed, i+1, test.args, got, pathGiven, test.wantedPath, pathGiven)
			}
		} else {
			if got != curr_wd {
				t.Errorf("%v Test %v - GetPath(%v) -> GOT: %v, %v. // WANT: %v, %v\n", NotPassed, i+1, test.args, got, pathGiven, curr_wd, test.pathWasGiven)
				TotalNotPassed += 1
			} else {
				fmt.Printf("%v Test %v - GetPath(%v) -> GOT: %v, %v. // WANT: %v, %v\n", Passed, i+1, test.args, got, pathGiven, curr_wd, test.pathWasGiven)
				TotalPassed += 1
			}
		}
		
	}
}

// TestGetPathInfo tests function GetPathInfo from package blit_cli
func TestGetPathInfo(t *testing.T) {
	var tests = []struct {
		path string
	}{
		{"/home/"},
		{"/usr/"},
		{"/usr/local/"},
		{"/etc/"},
		{"/var/"},
		{"/usr/"},
		{"/lib/"},
	}

	for i, test := range tests {
		fileInfo, err := GetPathInfo(test.path)

		if err != nil {
			t.Errorf("%v Test %v - Error in GetPathInfo call: %v\n", NotPassed, i+1, err)
			TotalNotPassed += 1
		} else {
			totFiles := len(fileInfo)
			if totFiles > 0  {
				fmt.Printf("%v Test %v - GOT: %v files // WANT: > 0 files in %v\n", Passed, i+1, totFiles, test.path)
				TotalPassed += 1
			}	
		}
		
	}
}

// BenchmarkGetPathInfo tests GetPathInfo performance
func BenchmarkGetPathInfo(b *testing.B) {
	path := "/usr/"

	for i := 0; i < b.N; i++ {
		GetPathInfo(path)
	}
}

// TestEncapData tests function EncapData from package blit_cli
func TestEncapData(t *testing.T) {
	
    type Test struct {
    	files []fs.FileInfo
    	path string
    }

    var tests []Test
    

    var paths = []string{"/home/", "/usr/local/", "/etc/"}

    for _, path := range paths {
    	f, err := os.Open(path)
		if err != nil {
			panic(err)
		}
		fileInfo, err := f.Readdir(-1)
		if err != nil {
			panic(err)
		}
		f.Close()
		test := new(Test)
		test.files 	= fileInfo
		test.path 	= path
    	tests = append(tests, *test)
    	
    }
    
    for i, test := range tests {
    	allFiles, err, totSize := EncapData(test.files, test.path)

    	if err != nil {
    		t.Errorf("%v Test %v - Couldn't complete. GOT: %v files, %v total size  // WANT: >0 files, >0 (B, Kb, Mb, ...) total size\n", NotPassed, i+1, len(allFiles), ByteToReadableSize(totSize))
    		TotalNotPassed += 1
    	} else {
    		fmt.Printf("%v Test %v - GOT: %v files, %v total size  // WANT: >0 files, >0 (B, Kb, Mb, ...) total size\n", Passed, i+1, len(allFiles), ByteToReadableSize(totSize))
    		TotalPassed += 1
    	}
    }
}

// TestCleanData tests function CleanData from package blit_cli
func TestCleanData(t *testing.T) {
	var tests = []struct {
		fullSli [][]string
		cutSli	[][]string
		partSli []string
	}{
		{ 	[][]string{{"0_a","0_b","0_c"},{"1_a","1_b","1_c"},{"2_a","2_b","2_c"}},
			[][]string{{"0_b","0_c"},{"1_b","1_c"},{"2_b","2_c"}},
			[]string{"0_a","1_a","2_a"},
		},
		{ 	[][]string{{"aa","bb","cc"},{"dd","ee","ff"},{"gg","hh","ii"}},
			[][]string{{"bb","cc"},{"ee","ff"},{"hh","ii"}},
			[]string{"aa","dd","gg"},
		},
	}	
	for i, test := range tests {
		cleanSli, _ := CleanData(test.fullSli)
		
		fullRowNum := len(test.fullSli)
		fullColNum := len(test.fullSli[0])
		cleanRowNum := len(cleanSli)
		cleanColNum := len(cleanSli[0])
		wantCol		:= len(test.cutSli[0])
		
		if cleanColNum == fullColNum {
			t.Errorf("%v Test %v - Couldn't complete. GOT: [%v][%v] // WANT: [%v][%v]\n", NotPassed, i+1, cleanRowNum, cleanColNum, fullRowNum, wantCol)
			TotalNotPassed += 1
		} else {
			fmt.Printf("%v Test %v - GOT: [%v][%v] // WANT: [%v][%v]\n", Passed, i+1, cleanRowNum, cleanColNum, fullRowNum, wantCol)
			TotalPassed += 1
		}
	}	
}

// TestByteToReadableSize tests function ByteToReadableSize from package blit_cli
func TestByteToReadableSize(t *testing.T) {
	var tests = []struct {
		num int64
		want string
	}{
		{1229, "1.2 Kb"},
		{1048576, "1.0 Mb"},
		{1073741824, "1.0 Gb"},
		{2097152, "2.0 Mb"},
		{2202009, "2.1 Mb"},
		{3113851290, "2.9 Gb"},
	}


	for i, test := range tests {
		got := ByteToReadableSize(test.num)

		if got != test.want {
			t.Errorf("%v Test %v - GOT: %v // WANT: %v\n", NotPassed, i+1, got, test.want)
			TotalNotPassed += 1
		} else {
			fmt.Printf("%v Test %v - GOT: %v // WANT: %v\n", Passed, i+1, got, test.want)
			TotalPassed += 1
		}
	}
}


// TestFileSizeSort tests function FileSizeSort from package blit_cli
func TestFileSizeSort(t *testing.T)  {
	var tests = []struct{
		sizeSli [][]int
		size_pos int
	}{
		{ 	[][]int{{1229,0},{1048576,1},{3113851290,2},{2202009,3},{1073741824,4},{2097152,5}}, 
			0,
		},
		{ 	[][]int{{1048576,1},{1229,0},{3113851290,2},{2202009,3},{1073741824,4},{2097152,5}}, 
			0,
		},
		{ 	[][]int{{2097152,5},{1229,0},{1048576,1},{1073741824,4},{3113851290,2},{2202009,3}}, 
			0,
		},
	}

	ordExpectSli := []int{2,4,3,5,1,0}

	for j, test := range tests {
		FileSizeSort(test.sizeSli, test.size_pos)
		var order []int
		for i, row := range test.sizeSli {
			//if row[1] != ordExpectSli[i][1] 
			oldPos := row[1]
			order = append(order, oldPos)
			if oldPos != ordExpectSli[i] {
				t.Errorf("%v Test %v - FileSort failed in slice pos: %v", NotPassed, j+1, i)
				TotalNotPassed += 1
			}
		}

		fmt.Printf("%v Test %v \n  GOT:    %v\n  WANTED: %v\n", Passed, j+1, order, ordExpectSli)
		TotalPassed += 1
	}

}

// TestSwap tests function Test from package blit_cli
func TestSwap(t *testing.T) {
	var tests = []struct{
		sizeSli [][]int
		size_pos int
	}{
		{ 	[][]int{{1229,0},{1048576,1},{3113851290,2},{2202009,3},{1073741824,4},{2097152,5}}, 
			0,
		},
		{ 	[][]int{{1048576,1},{1229,0},{3113851290,2},{2202009,3},{1073741824,4},{2097152,5}}, 
			0,
		},
		{ 	[][]int{{2097152,5},{1229,0},{1048576,1},{1073741824,4},{3113851290,2},{2202009,3}}, 
			0,
		},
	}

	var swapped = []struct {
		expectedSli [][]int
		size_pos int	
	}{
		{ 	[][]int{{1229,0},{3113851290,2},{1048576,1},{2202009,3},{1073741824,4},{2097152,5}}, 
			0,
		},
		{ 	[][]int{{1048576,1},{3113851290,2},{1229,0},{2202009,3},{1073741824,4},{2097152,5}}, 
			0,
		},
		{ 	[][]int{{2097152,5},{1048576,1},{1229,0},{1073741824,4},{3113851290,2},{2202009,3}}, 
			0,
		},
	}

	swappingPos := 1

	for i, test := range tests {
		Swap(test.sizeSli, swappingPos)

		sizeSliSwap_1 		:= test.sizeSli[swappingPos][0]
		expectedSizeIn_1 	:= swapped[i].expectedSli[swappingPos][0]
		
		if sizeSliSwap_1 != expectedSizeIn_1 {
			t.Errorf("%v Test %v - Swap failed at echanging values at row: %v", NotPassed, i+1, swappingPos)
			TotalNotPassed += 1
		}

		fmt.Printf("%v Test %v - \n  GOT:    %v\n  WANTED: %v\n", Passed, i+1, sizeSliSwap_1, expectedSizeIn_1)	
		TotalPassed += 1
	}
	
	fmt.Printf("Total tests %v      %v\n", Passed, TotalPassed)
	if (TotalNotPassed > 0) {
		fmt.Printf("Total tests %v  %v\n", NotPassed, TotalNotPassed)
	} else {
		fmt.Printf("Total tests %v  %v\n", WhiteNotPassed, TotalNotPassed)
	}
	
}
