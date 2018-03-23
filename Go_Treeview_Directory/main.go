package main

import (
	"os"
	"fmt"
	"path/filepath"
	"strings"
)
/*
func main() {
	dir := "/Users/nixon/Documents/Go/src/directory"
	var directory = []string{}

	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Printf("prevent panic by handling failure accessing a path %q: %v\n", dir, err)
			return err
		}
		directory = append(directory,path)

		return nil
	})

	if err != nil {
		fmt.Println("error walking path %q: %v", dir, err)
	}
	printTree(directory)
}


func printTree(dir []string) {
	for _,path := range dir {
		var folder []string
		folder = strings.Split(path, "/")
		result := ""
		printLine(6, (len(folder)-1), folder, result)
	}
}

func printLine(parentDepth int, last int, folder []string, result string){
	if parentDepth < last{
		result += "  "
		parentDepth++
		printLine(parentDepth, last, folder, result)
	} else {
		result += "|->" + folder[last]
		fmt.Println(result)
	}
}


func CheckFolder(folderPath string) (error) {
	dir, err := os.Open(folderPath)
	if err != nil {
		return err
	}
	defer dir.Close()

	fi, err := dir.Stat()
	if err != nil {
		return err
	}

	if fi.IsDir() {
		fis, err := dir.Readdir(-1) // -1 means return all the FileInfos
		if err != nil {
			return err
		}

		for _, fileinfo := range fis {
			if !fileinfo.IsDir() {
				fmt.Println(folderPath + "/" + fileinfo.Name())
			} else {
				err = CheckFolder(folderPath + "/" + fileinfo.Name())
				if err != nil {
					return err
				}
			}
		}
	}
	return nil
}



/*