package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter Name: ")
	inputName, _ := reader.ReadString('\n')
	fmt.Println("Hello ", inputName)
}
