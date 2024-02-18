package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter Name: ")
	// _ is added becuase we are adding a number
	inputName, _ := reader.ReadString('\n')
	fmt.Println("Hello ", inputName)

	// to demo  error handling , and type conversion
	fmt.Print("Enter Number: ")
	numberInputString, _ := reader.ReadString('\n')
	numberInput, err := strconv.ParseInt(strings.TrimSpace(numberInputString), 0, 64)
	if err == nil {
		fmt.Println("You Entered Number ", numberInput)
	} else {
		fmt.Println(err)
	}

}
