package main

import (
	"flag"
	"fmt"

	"github.com/nidhey27/coffee-assignment/internal/file"
)

func main() {
	// Read File - Sample Input
	// Define a command-line flag for the file path
	filePath := flag.String("file", "", "Path to the input file")
	flag.Parse()

	// Check if the file path is provided
	if *filePath == "" {
		fmt.Println("Please provide the file path using the -file flag.")
		return
	}

	data, err := file.ReadFile(*filePath)
	if err != nil {
		fmt.Println("ERROR: ", err.Error())
		panic(err)
	}

	file.ProcessInput(data)
}
