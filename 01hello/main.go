package main

import "fmt"

func main() {
	fmt.Println("Hello world")

	result := add(1, 9)
	fmt.Println(result)
}
func add(a int, b int) int {
	return a + b
}
