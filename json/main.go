package main

import (
	"encoding/json"
	"fmt"
	"strings"
)

type cartItem struct {
	Name     string
	Price    float64
	Quantity int
}

// getCartFromJson() returns a slice containing cartItem objects.
func getCartFromJson(jsonString string) []cartItem {
	var cart []cartItem
	// Your code goes here.
	decoder := json.NewDecoder(strings.NewReader(jsonString))
	_, err := decoder.Token()
	if err != nil {
		panic(err)
	}
	var item cartItem
	for decoder.More() {
		err := decoder.Decode(&item)
		if err != nil {
			panic(err)
		}
		cart = append(cart, item)
	}

	return cart
}

func main() {

	// This is how your code will be called.
	// Your answer should be a slice containing cartItem objects.
	// You can edit this code to try different testing cases.
	jsonString :=
		`[{"name":"apple","price":2.99,"quantity":3},
  {"name":"orange","price":1.50,"quantity":8},
  {"name":"banana","price":0.49,"quantity":12}]`

	fmt.Println(getCartFromJson(jsonString))

}
