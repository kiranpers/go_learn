// Write your answer here, and then test your code.
// Your job is to implement the findLargest() method.

package main

import (
	"fmt"
	"strconv"
)

// Change these boolean values to control whether you see
// the expected answer and/or hints.
const showExpectedResult = false
const showHints = false

// calculate() returns the sum of the two parameters
func calculate(value1 string, value2 string) float64 {
	// Your code goes here.

	// Convert the first string to a float64

	// Convert the second string to a float64

	// Calculate and return the result
	value1Float, _ := strconv.ParseFloat(value1, 64)
	value2Float, _ := strconv.ParseFloat(value2, 64)
	return value1Float + value2Float
}

func main() {
	value1 := "10"
	value2 := "5.5"
	result := calculate(value1, value2)
	fmt.Println(result)
}
