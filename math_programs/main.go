package main

import (
	"fmt"
	"math"
	"time"
)

func main() {

	var x int = 45
	var y float64 = 64.78

	//type conversion from float to int
	fmt.Println(x + int(y))

	var aFloat = 3.145
	fmt.Printf("%v\n", math.Round(aFloat))

	//multiple variable assignments in a single line
	i1, i2, i3 := 5, 7, 9
	sum := i1 + i2 + i3
	fmt.Println("Sum ", sum)

	// math area
	var radius = 5.67
	var circumference = 2 * math.Pi * radius
	fmt.Printf("circumference %.2f\n", circumference)

	fmt.Println(time.Now().Format(time.ANSIC))

}
