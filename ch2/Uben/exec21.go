// exec 2.1, can be divided into package format
package main

import (
    "fmt"
)
type Celsius float64
type Kelvin  float64

const (
    // C
	AbsoluteZeroC Celsius = -273.15
	FreezingC     Celsius = 0
	BoilingC      Celsius = 100
    // Kelvin
	AbsoluteZeroK Kelvin = 0
	FreezingK     Kelvin = 273.15
	BoilingK      Kelvin = 373.15
)

// CToK converts a Celsius temperature to Kelvin.
func CToK(c Celsius) Kelvin { return Kelvin(c + 273.15) }

// KToC converts a Kelvin temperature to Celsius.
func KToC(f Kelvin) Celsius { return Celsius(f - 273.15) }

func (c Celsius)    String() string { return fmt.Sprintf("%g°C", c) }
func (k Kelvin)     String() string { return fmt.Sprintf("%g°F", k) }

func main() {
    fmt.Printf("Now to tran Kelvin AbsoluteZero to C\n\n")
    fmt.Printf("Kelvin AbsoluteZero: %.2fK\n", AbsoluteZeroK)
    absoluteZeroC := KToC(AbsoluteZeroK)
    fmt.Printf("%.2fº, the C is %.2fº\n\n", absoluteZeroC, AbsoluteZeroC)
    fmt.Printf("Kelvin Freezing: %.2fK\n", FreezingK)
    freezingC := KToC(FreezingK)
    fmt.Printf("%.2fº, the C is %.2fº\n\n", freezingC, FreezingC)
    fmt.Printf("Kelvin Boiling: %.2fK\n", BoilingK)
    boilingC := KToC(BoilingK)
    fmt.Printf("%.2fº, the C is %.2fº\n\n", boilingC, BoilingC)
}
