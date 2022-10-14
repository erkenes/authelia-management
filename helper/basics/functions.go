package basics

import (
	"fmt"
	"os"
)

const ColorReset = "\033[0m"
const ColorRed = "\033[31m"
const ColorGreen = "\033[32m"
const ColorYellow = "\033[33m"
const ColorBlue = "\033[34m"
const ColorPurple = "\033[35m"
const ColorCyan = "\033[36m"
const ColorWhite = "\033[37m"

// WriteFile write something into a file
func WriteFile(filepath string, filename string, content []byte, filePerm os.FileMode) {
	fullFilepath := filename

	if filepath != "." && filepath != "" {
		fullFilepath = filepath + "/" + filename

		if _, err := os.Stat(fullFilepath); os.IsNotExist(err) {
			os.MkdirAll(filepath, 0700)
		}
	}

	err := os.WriteFile(fullFilepath, content, filePerm)
	if err != nil {
		fmt.Printf("Unable to write file: %v", err)
	}
}

// ReadFile read the content of a file
func ReadFile(filepath string, filename string) []byte {
	fullFilepath := filename

	if filepath != "." && filepath != "" {
		fullFilepath = filepath + "/" + filename

		if _, err := os.Stat(fullFilepath); os.IsNotExist(err) {
			return []byte{}
		}
	}

	content, err := os.ReadFile(fullFilepath)

	if err != nil {
		return []byte{}
	}

	return content
}

// RemoveFromStringSlice removes a string from a slice
func RemoveFromStringSlice(s []string, r string) []string {
	for i, v := range s {
		if v == r {
			return append(s[:i], s[i+1:]...)
		}
	}
	return s
}

// StringSliceContains checks if a string is present in a slice
func StringSliceContains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}

	return false
}

// UniqueStrings remove duplicates from slice
func UniqueStrings(stringSlices ...[]string) []string {
	uniqueMap := map[string]bool{}

	for _, stringSlice := range stringSlices {
		for _, content := range stringSlice {
			uniqueMap[content] = true
		}
	}

	// Create a slice with the capacity of unique items
	// This capacity make appending flow much more efficient
	result := make([]string, 0, len(uniqueMap))

	for key := range uniqueMap {
		result = append(result, key)
	}

	return result
}
