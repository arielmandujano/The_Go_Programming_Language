package exercises

import (
	"fmt"
	"os"
	"strings"
)

// Modify de echo program to also print os.Args[0], the name of the command that invoked it.

func Echo1_1() {
	fmt.Println(strings.Join(os.Args, " "))
}
