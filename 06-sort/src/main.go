package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/george.vince/coding-challenges/06-sort/src/lib"
)

func main() {
	var unique bool
	var sortType int
	flag.BoolVar(&unique, "u", false, "Sort unique words")
	flag.IntVar(&sortType, "sort", 0, "Algorithm to use for sorting (0=merge, 1=quick, 2=radix, 3=heap)")
	flag.Parse()
	if len(os.Args) < 2 {
		fmt.Println("Please provide file path")
		return
	}
	args := flag.Args()
	filePath := args[0]
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Error opening file: ", file)
		return

	}
	defer file.Close()
	contents, err := parse(file, unique)
	if err == nil {
		fmt.Println(contents)
	} else {
		fmt.Println("Error parsing file: ", err)
	}

	sorted := sort(contents, sortType)
	fmt.Println(sorted)
}

func parse(file *os.File, unique bool) ([]string, error) {
	var words []string
	scanner := bufio.NewScanner(file)
	seen := make(map[string]bool)
	isFirstLine := true
	for scanner.Scan() {
		newWord := scanner.Text()
		if isFirstLine {
			newWord = strings.TrimPrefix(newWord, "\uFEFF") // Remove BOM from the first line
			isFirstLine = false
		}
		if !unique || (unique && !seen[newWord]) {
			words = append(words, newWord)
			seen[newWord] = true
		}

	}

	return words, scanner.Err()
}

func sort(words []string, sortType int) []string {
	switch sortType {
	case 1:
		fmt.Println("Quick sort")
	case 2:
		fmt.Println("Radix sort")
	case 3:
		fmt.Println("Heap sort")
	default:
		fmt.Println("Hello")
		lib.MergeSort(words)
	}
	return words
}
