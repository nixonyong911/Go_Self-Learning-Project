package main

import (
	"os"
	"fmt"
	"path/filepath"
	"strings"
)

func main() {
	dir := "/Users/nixon/Documents/Go/src/"
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
	slice := len(strings.Split(dir[0], "/"))
	for _,path := range dir {
		var folder []string
		folder = strings.Split(path, "/")
		printLine(len(folder) - slice, folder[len(folder)-1])
	}
}

func printLine(length int, name string){
	var result string
	for i := 0; i<length; i++ {
		result += "|  "
	}
	result += "|->" + name
	fmt.Println(result)
}