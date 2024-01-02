package main

import (
	"fmt"
	"os"

	"focusthen.com/files/fileutils"
)

func main() {
	rootPath, _ := os.Getwd()

	c, err := fileutils.ReadTextFile(rootPath + "/data/text.txt")
	if err != nil {
		fmt.Printf("Error: %v", err)
	} else {
		fmt.Println(c)
	}

}
