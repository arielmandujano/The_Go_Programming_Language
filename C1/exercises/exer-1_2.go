package exercises

import (
	"fmt"
	"os"
)

// Modify the echo program to print the index and value of each of its arguments, one per line.

func Echo1_2() {
	for i, arg := range os.Args[1:] {
		fmt.Println("Index:", i+1, ". Value:", arg)
	}
}
