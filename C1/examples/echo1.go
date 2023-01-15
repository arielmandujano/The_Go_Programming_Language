package examples

// Echo1 prints all command-line arguments on a single line

import (
	"fmt"
	"os"
)

func Echo1() {
	var s, sep string
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
}
