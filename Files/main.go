package main

import (
	"fmt"
	"os"

	"focusthen.com/files/fileutils"
)

func main() {
	rootPath, _ := os.Getwd()

	filepath := rootPath + "/data"

	c, err := fileutils.ReadTextFile(fmt.Sprintf("%v/text.txt", filepath))
	if err != nil {
		fmt.Printf("Error: %v", err)
	} else {
		fileutils.WriteToFile(fmt.Sprintf("%v/output.txt", filepath), "Hello World")
		fmt.Println(c)
	}

}
