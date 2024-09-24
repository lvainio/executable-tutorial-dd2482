package main

import (
	"fmt"
	"os"
)

func main() {
	var studentFile string

	// E.g., a student's file (passed as user input)
	fmt.Println("Enter the student file name (e.g., student123.txt): ")
	fmt.Scan(&studentFile)

	// Construct the full filepath (Files are stored in "./students/")
	baseDir := "./students/"
	fullPath := baseDir + studentFile

	// Open the student file
	file, err := os.Open(fullPath) // Vulnerability: No sanitization or validation on the input
	if err != nil {
		fmt.Println("Error: Could not open the file", err)
		return
	}
	defer file.Close()

	// Print if file is opened successfully
	fmt.Printf("Successfully opened the file: %s\n", fullPath)

	// Read and print the file line by line
	scanner := bufio.NewScanner(file)
	fmt.Println("Contents of the file:")

	lineCount := 0
	for scanner.Scan() {
		fmt.Println(scanner.Text())
		lineCount++
		if lineCount >= 3 { // Limit to first 3 lines
			break
		}
	}

	// Check for errors during scanning
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading the file:", err)
		return
	}
}

