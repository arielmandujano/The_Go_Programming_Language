package examples

// Tempconv performs Celcius and Fahrenheit temperature computations.

type Celcius float64
type Fahrenheit float64

const (
	AbsoluteZeroC Celcius = -273.15
	FreezingC     Celcius = 0
	BoilingC      Celcius = 100
)

func CToF(c Celcius) Fahrenheit  { return Fahrenheit(c*9/5 + 32) }
func FToC2(f Fahrenheit) Celcius { return Celcius((f - 32) * 5 / 9) }
