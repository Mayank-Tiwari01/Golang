package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	read := bufio.NewReader(os.Stdin)
	fmt.Println("Enter a value: ")

	// Using comma ok syntax
	input, err := read.ReadString('\n') // read until it encounters a new line

	// If you do not care about the error, you can use the underscore (_) to ignore it.
	// input, _ := read.ReadString('\n')

	// Check if there was an error
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}

	fmt.Println("The value entered was: ", input)
	fmt.Printf("The type of the value entered is %T", input)
}
