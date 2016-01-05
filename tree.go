package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

func printListing(entry string, depth int) {
	indent := strings.Repeat("|   ", depth)
	fmt.Printf("%s|-- %s\n", indent, entry)
}

func printDirectory(path string, depth int) {
	entries, err := ioutil.ReadDir(path)
	if err != nil {
		fmt.Printf("error reading %s: %s\n", path, err.Error())
		return
	}

	printListing(path, depth)
	for _, entry := range entries {
		if (entry.Mode() & os.ModeSymlink) == os.ModeSymlink {
			full_path, err := os.Readlink(filepath.Join(path, entry.Name()))
			if err != nil {
				fmt.Printf("error reading link: %s\n", err.Error())
			} else {
				printListing(entry.Name()+" -> "+full_path, depth+1)
			}
		} else if entry.IsDir() {
			printDirectory(filepath.Join(path, entry.Name()), depth+1)
		} else {
			printListing(entry.Name(), depth+1)
		}
	}
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("usage: tree <path>")
		os.Exit(1)
	}

	printDirectory(os.Args[1], 0)
}
