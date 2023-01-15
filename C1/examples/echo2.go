package examples

// Echo2 prints all command-line arguments on a single line

import (
	"fmt"
	"os"
)

func Echo2() {
	s, sep := "", ""
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
}
