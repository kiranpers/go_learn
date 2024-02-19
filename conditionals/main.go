package main

import (
	"fmt"
	"math/rand"
	"time"
)

type cartItem struct {
	name     string
	price    float64
	quantity int
}

// calculateTotal() returns the total value of the shopping cart.
func calculateTotal(cart []cartItem) float64 {
	var sum float64 = 0.0
	for _, item := range cart {
		sum = sum + (item.price * float64(item.quantity))
	}
	return sum
}

func main() {

	var x int = -42
	if x > 0 {
		fmt.Printf("%d is positive\n", x)
	} else if x == 0 {
		fmt.Printf("%d is zero\n", x)
	} else {
		fmt.Printf("%d is negative\n", x)
	}
	rand.Seed(time.Now().Unix())
	dow := rand.Intn(7) + 1
	var result string
	switch dow {
	case 1:
		result = "Mon"
		// fallthrough
	case 2:
		result = "Tue"
		// fallthrough
	default:
		result = "It is someother day"
	}
	fmt.Println(result)

	//
	colors := []string{"Blue", "Green", "Red"}
	for i := 0; i < len(colors); i++ {
		fmt.Println(i, " ", colors[i])
	}

	var cart []cartItem
	var apples = cartItem{"apple", 2.99, 3}
	var oranges = cartItem{"orange", 1.50, 8}
	var bananas = cartItem{"banana", .49, 12}
	cart = append(cart, apples, oranges, bananas)
	fmt.Println("Cart Total ", calculateTotal(cart))

	sum := 1

	for sum < 1000 {
		sum += sum
		if sum > 200 {
			goto theEnd
		}
	}
theEnd:
	fmt.Println("End of program")

}
