package main

import (
	"fmt"
	"sort"
)

type Student struct {
	FirstName string
	LastName  string
}

func convertToMap(items []string) map[string]float64 {

	// Create a map object
	result := make(map[string]float64)
	percentage := 100.00 / float64(len(items))
	for _, k := range items {
		// perform an operation
		result[k] = percentage
	}
	// Your code goes here
	return result
}

func main() {
	m := make(map[string]int)
	m["Kiran"] = 7
	fmt.Println(m)

	anInt := 42
	p := &anInt
	fmt.Println("Access value using pointer :", *p)

	*p = 43
	fmt.Println("Changing value using pointer", *p)

	// arrays
	var colors = [3]string{"Red", "Green", "Blue"}
	fmt.Println(colors)

	//slices
	var colorSlice = []string{"Red", "Green", "Blue"}
	colorSlice = append(colorSlice, "Orange")
	fmt.Println(colorSlice)

	//slices - skip first element
	fmt.Println(colorSlice[1:len(colorSlice)])
	sort.Strings(colorSlice)
	fmt.Println("Sorted Strings", colorSlice)

	var numbers = []int{5, 7, 9, 1}
	sort.Ints(numbers)
	fmt.Println("Sorted Numbers", numbers)

	states := make(map[string]string)
	states["TX"] = "Texas"
	states["WA"] = "Washington"
	states["FL"] = "Florida"

	fmt.Println(states["FL"])

	delete(states, "FL")
	fmt.Println("Map after deleting a Key :", states)

	fmt.Println("Print Map")

	for k, v := range states {
		fmt.Printf("%v : %v\n", k, v)
	}

	s1 := Student{"John", "Doe"}
	fmt.Printf("%+v", s1)

	slice := []string{"apple", "banana", "orange"}
	// converts a slice to percentages for a 3 item slice, each
	// values is 33.33
	result := convertToMap(slice)
	fmt.Println(result)

}
