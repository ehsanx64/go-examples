package main

import (
	"fmt"
	"math"
)

// Numeric constants
const num1 int = 255
const num2 = 1900

// Explicitly define an identifier named 's' as a string constant
const str1 string = "A string"
const str2 = "Another string"

// Character constants
const char1 = 'A'
const char2 = "\n"

// Bolean constants
const always_check = true
const is_running bool = false

const (
	// left shift one 100 times
	Big   = 1 << 100
	Small = Big >> 99
)

func main() {
	// Print the characters
	fmt.Println("char1:", char1)
	fmt.Println("char2:", char2)

	// Constant expressions perform arithmetic with arbitrary precision.
	const n = 500000000
	const d = 3e20 / n
	fmt.Println("d :", d)

	// A numeric constant has no type until it is given one
	fmt.Println("int64(d) :", int64(d))

	// A numeric constant can also be given a type when it is used in a context
	// that requires that type. For example n would have the int64 type because
	// math.Sin() requires its parameter to be an int64.
	fmt.Print("math.Sin(")
	fmt.Print(n)
	fmt.Println(") :", math.Sin(n))

	demoNumericConstants()
}

func needInt(x int) int {
	return x*10 + 1
}

func needFloat(x float64) float64 {
	return x * 0.1
}

func demoNumericConstants() {
	fmt.Println()
	fmt.Println("* Numeric constants")
	fmt.Println("Small:", Small)
	fmt.Println("int:", needInt(Small))
	fmt.Println("float(Small):", needFloat(Small))
	fmt.Println("float(Big):", needFloat(Big))
}
