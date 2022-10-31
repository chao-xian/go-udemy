package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	filePath := os.Args[1]
	// content, err := ioutil.ReadFile(filePath)
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Err")
	}

	io.Copy(os.Stdout, file)
}
