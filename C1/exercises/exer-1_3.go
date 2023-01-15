package exercises

import (
	"The_Go_Programming_Lenguage/C1/examples"
	"fmt"
	"time"
)

// Experiment to measure the difference in running time between our potentially inefficient versions and the one that uses strings.Join.

func TimingEcho() {
	startEcho1 := time.Now()

	examples.Echo1()

	secondsEcho1 := time.Since(startEcho1).Seconds()

	startEcho2 := time.Now()

	examples.Echo2()

	secondsEcho2 := time.Since(startEcho2).Seconds()

	startEcho3 := time.Now()

	examples.Echo3()

	secondsEcho3 := time.Since(startEcho3).Seconds()

	fmt.Printf("---- Echo1. Seconds: %.7fs\n", secondsEcho1)
	fmt.Printf("---- Echo2. Seconds: %.7fs\n", secondsEcho2)
	fmt.Printf("---- Echo3. Seconds: %.7fs\n", secondsEcho3)
}
