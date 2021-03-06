package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	fmt.Println(4 ^ 4)
	if len(os.Args) != 2 {
		fmt.Println("Please specify a path.")
		return
	}

	root, err := filepath.Abs(os.Args[1])
	if err != nil {
		fmt.Println("Cannot get absolute path:", err)
		return
	}

	fmt.Println("Listing files in", root)
	var c struct {
		files int
		dirs  int
	}
	filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			c.dirs++
		} else {
			c.files++
		}
		fmt.Println("-", path)
		return nil
	})
	fmt.Printf("Total %d files in %d directories", c.files, c.dirs)
}
