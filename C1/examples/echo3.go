package examples

import (
	"fmt"
	"os"
	"strings"
)

// Echo3 prints all command-line arguments on a single line

func Echo3() {
	fmt.Println(strings.Join(os.Args[1:], " "))
}
