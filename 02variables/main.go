package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func main() {
	var userName string
	fmt.Println(userName)
	fmt.Printf("the type is: %T \n", userName)

	var numbers [3]int = [3]int{1, 2, 3}         // Array
	var scores []int = []int{95, 85, 75}         // Slice
	var john = Person{Name: "Mayank", Age: 6700} // Struct

	fmt.Println(numbers, scores, john)
}
