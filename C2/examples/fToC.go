package examples

import "fmt"

// FToC prints two Fahrenheit-To-Celsius convertions
func FToC() {
	freezingF := 32.0
	fmt.Printf("%gºF = %gºC\n", freezingF, fToC(freezingF))
	fmt.Printf("%gºF = %gºC\n", boilingF, fToC(boilingF))
}

func fToC(f float64) float64 {
	return (f - 32) * 5 / 9
}
