package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
)

type cartItem struct {
	name     string
	price    float64
	quantity int
}

// calculateTotal() returns the total value of the shopping cart.
func (c cartItem) calculateItemTotal() float64 {
	return c.price * float64(c.quantity)
}

// calculateTotal() returns the total value of the shopping cart.
func calculateTotal(cart []cartItem) float64 {
	var sum float64 = 0.0
	for _, item := range cart {
		sum = sum + item.calculateItemTotal()
	}
	return sum
}

// calculate() returns the result of the requested operation.
func calculate(input1 string, input2 string, operation string) float64 {
	input1Float := convertInputToValue(input1)
	input2Float := convertInputToValue(input2)
	switch operation {
	case "+":
		return addValues(input1Float, input2Float)
	case "-":
		return subtractValues(input1Float, input2Float)
	case "*":
		return multiplyValues(input1Float, input2Float)
	case "/":
		return divideValues(input1Float, input2Float)
	}
	return 0
}

func convertInputToValue(input string) float64 {
	inputFloat, _ := strconv.ParseFloat(input, 64)
	return inputFloat
}

func addValues(value1, value2 float64) float64 {
	return value1 + value2
}

func subtractValues(value1, value2 float64) float64 {
	return value1 - value2
}

func multiplyValues(value1, value2 float64) float64 {
	return value1 * value2
}

func divideValues(value1, value2 float64) float64 {
	return value1 / value2
}

const url = "http://services.explorecalifornia.org/json/tours.php"

func makeHttpCall() {
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	bytes, err := ioutil.ReadAll(resp.Body)
	fmt.Println(string(bytes))
}

func main() {
	makeHttpCall()

	var cart []cartItem
	var apples = cartItem{"apple", 2.99, 3}
	var oranges = cartItem{"orange", 1.50, 8}
	var bananas = cartItem{"banana", .49, 12}
	cart = append(cart, apples, oranges, bananas)
	fmt.Println("Cart Total ", calculateTotal(cart))

	value1 := "10"
	value2 := "5.5"
	operation := "/"
	fmt.Println("Result ", calculate(value1, value2, operation))

}
