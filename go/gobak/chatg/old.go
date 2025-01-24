package main

import (
	"fmt"
	"os"
	"time"
)

// Function to get the time of n weeks ago
func getTimeThreshold(weeks int) time.Time {
	// Subtract n weeks from the current time
	return time.Now().AddDate(0, 0, -weeks*7)
}

// Function to print files older than n weeks
func printOldFiles(directory string, weeks int) error {
	// Get the time threshold (n weeks ago)
	timeThreshold := getTimeThreshold(weeks)

	// Open the directory
	dir, err := os.Open(directory)
	if err != nil {
		return fmt.Errorf("error opening directory: %v", err)
	}
	defer dir.Close()

	// Read the files in the directory
	files, err := dir.Readdir(-1)
	if err != nil {
		return fmt.Errorf("error reading directory: %v", err)
	}

	// Loop through the files and check their modification time
	for _, file := range files {
		// Get the file's modification time
		fileModTime := file.ModTime()

		// Compare with the threshold time (files older than n weeks)
		if fileModTime.Before(timeThreshold) {
			// Print file path and modification time
			fmt.Printf("File: %s, Last Modified: %s\n", file.Name(), fileModTime)
		}
	}
	return nil
}

func main() {
	// Directory to scan and the number of weeks
	directory := "./" // Change this to the directory you want to scan
	weeks := 4        // Change this to the number of weeks

	// Print files older than the specified number of weeks
	err := printOldFiles(directory, weeks)
	if err != nil {
		fmt.Println("Error:", err)
	}
}
