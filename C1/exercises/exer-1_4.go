package exercises

import (
	"bufio"
	"fmt"
	"os"
)

// Modify dup2 to print the names of all files in which duplicated line occurs

func Exec1_4() {
	counts := make(map[string]int)
	files := os.Args[1:]
	filesDupLines := make(map[string]bool)
	if len(files) == 0 {
		countLinesCLI(os.Stdin, counts)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts, filesDupLines)
			f.Close()
		}
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
	for filename := range filesDupLines {
		fmt.Println("File name:", filename)
	}
}

func countLinesCLI(f *os.File, counts map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
	}
	// NOTE: ignoring potential errors from input.Err()
}

func countLines(f *os.File, counts map[string]int, filesDupLines map[string]bool) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
		if counts[input.Text()] > 1 {
			filesDupLines[f.Name()] = true
		}
	}
	// NOTE: ignoring potential errors from input.Err()
}
